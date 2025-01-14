package runner

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"gopkg.in/yaml.v3"

	"github.com/cresplanex/bloader/internal/encrypt"
	"github.com/cresplanex/bloader/internal/logger"
)

// BaseExecutor represents the base executor
type BaseExecutor struct {
	Env                   string
	EncryptCtr            encrypt.Container
	Logger                logger.Logger
	SlaveConnectContainer *ConnectionContainer
	TmplFactor            TmplFactor
	Store                 Store
	AuthFactor            AuthenticatorFactor
	OutputFactor          OutputFactor
	TargetFactor          TargetFactor
}

// Execute executes the base executor
func (e BaseExecutor) Execute(
	ctx context.Context,
	filename string,
	str *sync.Map,
	threadOnlyStr *sync.Map,
	outputRoot string,
	index int,
	callCount int,
	slaveValues map[string]any,
	eventCaster EventCaster,
) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Printf(
		"BaseExecutor.Execute: filename=%s, outputRoot=%s, index=%d, callCount=%d\n",
		filename,
		outputRoot,
		index,
		callCount,
	)
	defer fmt.Printf(
		"Terminate BaseExecutor.Execute: filename=%s, outputRoot=%s, index=%d, callCount=%d\n",
		filename,
		outputRoot,
		index,
		callCount,
	)

	if err := eventCaster.CastEvent(ctx, RunnerEventStart); err != nil {
		return fmt.Errorf("failed to cast event: %w", err)
	}

	tmplStr, err := e.TmplFactor.TmplFactorize(ctx, filename)
	if err != nil {
		return fmt.Errorf("failed to factorize template: %w", err)
	}

	tmpl, err := template.New("yaml").Funcs(sprig.TxtFuncMap()).Parse(tmplStr)
	if err != nil {
		return fmt.Errorf("failed to parse yaml: %w", err)
	}
	replacedValuesData := make(map[string]any)
	replaceThreadValuesData := make(map[string]any)

	str.Range(func(key, value any) bool {
		if keyStr, ok := key.(string); ok {
			replacedValuesData[keyStr] = value
		}
		return true
	})

	threadOnlyStr.Range(func(key, value any) bool {
		if keyStr, ok := key.(string); ok {
			replaceThreadValuesData[keyStr] = value
		}
		return true
	})

	data := map[string]any{
		"SlaveValues":  slaveValues,
		"Values":       replacedValuesData,
		"ThreadValues": replaceThreadValuesData,
		"Dynamic": map[string]any{
			"OutputRoot": outputRoot,
			"LoopCount":  index,
			"CallCount":  callCount,
		},
	}

	yamlBuf := &bytes.Buffer{}
	if err := tmpl.Execute(yamlBuf, data); err != nil {
		return fmt.Errorf("failed to execute yaml: %w", err)
	}

	var rawData bytes.Buffer
	reader := io.TeeReader(yamlBuf, &rawData)

	var runner Runner
	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(&runner); err != nil {
		return fmt.Errorf("failed to decode yaml: %w", err)
	}

	validRunner, err := runner.Validate()
	if err != nil {
		return fmt.Errorf("failed to validate runner: %w", err)
	}

	if validRunner.StoreImport.Enabled {
		if err := eventCaster.CastEvent(ctx, RunnerEventStoreImporting); err != nil {
			return fmt.Errorf("failed to cast event: %w", err)
		}

		if err := e.Store.Import(
			ctx,
			validRunner.StoreImport.Data,
			func(_ context.Context, data ValidStoreImportData, val any, _ []byte) error {
				if data.ThreadOnly {
					threadOnlyStr.Store(data.Key, val)
					replaceThreadValuesData[data.Key] = val
				} else {
					str.Store(data.Key, val)
					replacedValuesData[data.Key] = val
				}
				return nil
			},
		); err != nil {
			return fmt.Errorf("failed to import data: %w", err)
		}

		data = map[string]any{
			"SlaveValues":  slaveValues,
			"Values":       replacedValuesData,
			"ThreadValues": replaceThreadValuesData,
			"Dynamic": map[string]any{
				"OutputRoot": outputRoot,
				"LoopCount":  index,
				"CallCount":  callCount,
			},
		}

		yamlBuf := &bytes.Buffer{}
		if err := tmpl.Execute(yamlBuf, data); err != nil {
			return fmt.Errorf("failed to execute yaml: %w", err)
		}
		rawData.Reset()
		reader := io.TeeReader(yamlBuf, &rawData)
		decoder := yaml.NewDecoder(reader)
		if err := decoder.Decode(&runner); err != nil {
			return fmt.Errorf("failed to decode yaml: %w", err)
		}

		validRunner, err = runner.Validate()
		if err != nil {
			return fmt.Errorf("failed to validate runner: %w", err)
		}

		if err := eventCaster.CastEvent(ctx, RunnerEventStoreImported); err != nil {
			return fmt.Errorf("failed to cast event: %w", err)
		}
	}

	if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterInit, filename); err != nil {
		return fmt.Errorf("failed to wait: %w", err)
	}

	switch validRunner.Kind {
	case RunnerKindStoreValue:
		var storeValue StoreValue
		decoder := yaml.NewDecoder(&rawData)
		if err := decoder.Decode(&storeValue); err != nil {
			return fmt.Errorf("failed to decode yaml: %w", err)
		}
		var validStoreValue ValidStoreValue
		if err := validate(ctx, eventCaster, func() error {
			if validStoreValue, err = storeValue.Validate(); err != nil {
				return fmt.Errorf("failed to validate store value: %w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := validStoreValue.Run(ctx, e.Store); err != nil {
			if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterFailedExec, filename); err != nil {
				return fmt.Errorf("failed to wait: %w", err)
			}
			return fmt.Errorf("failed to execute store value: %w", err)
		}
		e.Logger.Info(ctx, "executed store value")
	case RunnerKindMemoryValue:
		var memoryStoreValue MemoryValue
		decoder := yaml.NewDecoder(&rawData)
		if err := decoder.Decode(&memoryStoreValue); err != nil {
			return fmt.Errorf("failed to decode yaml: %w", err)
		}
		var validMemoryValue ValidMemoryValue
		if err := validate(ctx, eventCaster, func() error {
			if validMemoryValue, err = memoryStoreValue.Validate(); err != nil {
				return fmt.Errorf("failed to validate memory store value: %w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := validMemoryValue.Run(ctx, str); err != nil {
			if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterFailedExec, filename); err != nil {
				return fmt.Errorf("failed to wait: %w", err)
			}
			return fmt.Errorf("failed to execute memory store value: %w", err)
		}
		e.Logger.Info(ctx, "executed memory store value")
	case RunnerKindStoreImport:
		var storeImport StoreImport
		decoder := yaml.NewDecoder(&rawData)
		if err := decoder.Decode(&storeImport); err != nil {
			return fmt.Errorf("failed to decode yaml: %w", err)
		}
		var validStoreImport ValidStoreImport
		if err := validate(ctx, eventCaster, func() error {
			if validStoreImport, err = storeImport.Validate(); err != nil {
				return fmt.Errorf("failed to validate store import: %w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := validStoreImport.Run(ctx, e.Store, str); err != nil {
			if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterFailedExec, filename); err != nil {
				return fmt.Errorf("failed to wait: %w", err)
			}
			return fmt.Errorf("failed to execute store import: %w", err)
		}
		e.Logger.Info(ctx, "executed store import")
	case RunnerKindOneExecute:
		var oneExec OneExec
		decoder := yaml.NewDecoder(&rawData)
		if err := decoder.Decode(&oneExec); err != nil {
			return fmt.Errorf("failed to decode yaml: %w", err)
		}
		var validOneExec ValidOneExec
		if err := validate(ctx, eventCaster, func() error {
			if validOneExec, err = oneExec.Validate(ctx, e.AuthFactor, e.OutputFactor, e.TargetFactor); err != nil {
				return fmt.Errorf("failed to validate one exec: %w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := validOneExec.Run(ctx, outputRoot, str, e.Logger, e.Store); err != nil {
			if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterFailedExec, filename); err != nil {
				return fmt.Errorf("failed to wait: %w", err)
			}
			return fmt.Errorf("failed to execute one exec: %w", err)
		}
		e.Logger.Info(ctx, "executed one exec")
	case RunnerKindMassExecute:
		var massExec MassExec
		decoder := yaml.NewDecoder(&rawData)
		if err := decoder.Decode(&massExec); err != nil {
			return fmt.Errorf("failed to decode yaml: %w", err)
		}
		var validMassExec ValidMassExec
		if err := validate(ctx, eventCaster, func() error {
			if validMassExec, err = massExec.Validate(
				ctx,
				e.Logger,
				e.AuthFactor,
				e.OutputFactor,
				e.TargetFactor,
				tmplStr,
				data,
			); err != nil {
				return fmt.Errorf("failed to validate mass exec: %w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := validMassExec.Run(
			ctx,
			e.Logger,
			outputRoot,
			e.AuthFactor,
			e.OutputFactor,
			e.TargetFactor,
		); err != nil {
			if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterFailedExec, filename); err != nil {
				return fmt.Errorf("failed to wait: %w", err)
			}
			return fmt.Errorf("failed to execute mass exec: %w", err)
		}
		e.Logger.Info(ctx, "executed mass exec")
	case RunnerKindSlaveConnect:
		var slaveConnect SlaveConnect
		decoder := yaml.NewDecoder(&rawData)
		if err := decoder.Decode(&slaveConnect); err != nil {
			return fmt.Errorf("failed to decode yaml: %w", err)
		}
		var validSlaveConnect ValidSlaveConnect
		if err := validate(ctx, eventCaster, func() error {
			if validSlaveConnect, err = slaveConnect.Validate(); err != nil {
				return fmt.Errorf("failed to validate slave connect: %w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := e.SlaveConnectContainer.Connect(
			ctx,
			e.Logger,
			e.Env,
			e.EncryptCtr,
			validSlaveConnect,
			eventCaster,
		); err != nil {
			if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterFailedExec, filename); err != nil {
				return fmt.Errorf("failed to wait: %w", err)
			}
			return fmt.Errorf("failed to connect to slave: %w", err)
		}
		var atomicErr atomic.Pointer[syncError]
		var wg sync.WaitGroup
		for _, slave := range validSlaveConnect.Slaves {
			wg.Add(1)
			mapData, ok := e.SlaveConnectContainer.Find(slave.ID)
			if !ok {
				return fmt.Errorf("failed to find slave: %s", slave.ID)
			}
			slHandler := NewSlaveRequestHandler(mapData.ReqChan, mapData.Cli, mapData.ReceiveTermChan)
			go func(slaveHandler *SlaveRequestHandler) {
				defer wg.Done()
				if err := slaveHandler.HandleResponse(
					ctx,
					e.Logger,
					e.TmplFactor,
					e.AuthFactor,
					e.TargetFactor,
					e.Store,
				); err != nil {
					atomicErr.Store(&syncError{Err: err})
					e.Logger.Error(ctx, "failed to handle response: %v",
						logger.Value("error", err))
					cancel()
					return
				}
			}(slHandler)
		}
		wg.Wait()
		if syncErr := atomicErr.Load(); syncErr != nil {
			e.Logger.Error(ctx, "failed to find error",
				logger.Value("error", syncErr.Err))
			return syncErr.Err
		}
		e.Logger.Info(ctx, "connected to slave node")
	case RunnerKindFlow:
		var flow Flow
		decoder := yaml.NewDecoder(&rawData)
		if err := decoder.Decode(&flow); err != nil {
			return fmt.Errorf("failed to decode yaml: %w", err)
		}
		var validFlow ValidFlow
		if err := validate(ctx, eventCaster, func() error {
			if validFlow, err = flow.Validate(); err != nil {
				return fmt.Errorf("failed to validate flow: %w", err)
			}
			return nil
		}); err != nil {
			return err
		}
		if err := validFlow.Run(
			ctx,
			e.Env,
			e.Logger,
			e.SlaveConnectContainer,
			e.EncryptCtr,
			e.TmplFactor,
			e.Store,
			e.AuthFactor,
			e.OutputFactor,
			e.TargetFactor,
			str,
			outputRoot,
			callCount,
			slaveValues,
		); err != nil {
			if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterFailedExec, filename); err != nil {
				return fmt.Errorf("failed to wait: %w", err)
			}
			return fmt.Errorf("failed to execute flow: %w", err)
		}
		e.Logger.Info(ctx, "executed flow")
	default:
		return fmt.Errorf("invalid runner kind: %s", validRunner.Kind)
	}

	if err := wait(ctx, e.Logger, validRunner, RunnerSleepValueAfterExec, filename); err != nil {
		return fmt.Errorf("failed to wait: %w", err)
	}

	return nil
}
