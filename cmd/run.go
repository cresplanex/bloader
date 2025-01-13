/*
Copyright © 2024 hayashi kenta <k.hayashi@cresplanex.com>
*/
package cmd

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/runner"
)

var (
	runnerFile string
	runnerData []string
)

const (
	runnerDataTypesInt         = "i"
	runnerDataTypesString      = "s"
	runnerDataTypesBool        = "b"
	runnerDataTypesFloat       = "f"
	runnerDataTypesUint        = "u"
	runnerDataTypesArrayInt    = "ai"
	runnerDataTypesArrayString = "as"
	runnerDataTypesArrayBool   = "ab"
	runnerDataTypesArrayFloat  = "af"
	runnerDataTypesArrayUint   = "au"

	defaultRunnerDataTypes = runnerDataTypesString
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the load test",
	Long: `This command runs the load test.
It sends requests to the specified server and measures the response time.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ctr.Config.Type == config.ConfigTypeSlave {
			color.Red("This command is not available in slave mode")
			return
		}

		ctx, cancel := context.WithCancel(ctr.Ctx)
		defer cancel()

		ctr.Ctx = ctx

		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-signalChan
			cancel()
		}()

		data := make(map[string]any)
		var err error
		for _, d := range runnerData {
			kv := strings.Split(d, "=")
			if len(kv) != 2 {
				color.Red("Invalid data: %s\n", d)
				return
			}

			vt := defaultRunnerDataTypes
			var vo string
			if strings.Contains(kv[1], ":") {
				v := strings.Split(kv[1], ":")
				vt = v[1]
				vo = v[0]
			}

			switch vt {
			case runnerDataTypesInt:
				data[kv[0]], err = strconv.Atoi(vo)
				if err != nil {
					color.Red("Failed to parse int: %v\n", err)
					return
				}
			case runnerDataTypesString:
				data[kv[0]] = vo
			case runnerDataTypesBool:
				data[kv[0]], err = strconv.ParseBool(vo)
				if err != nil {
					color.Red("Failed to parse bool: %v\n", err)
					return
				}
			case runnerDataTypesFloat:
				data[kv[0]], err = strconv.ParseFloat(vo, 64)
				if err != nil {
					color.Red("Failed to parse float: %v\n", err)
					return
				}
			case runnerDataTypesUint:
				data[kv[0]], err = strconv.ParseUint(vo, 10, 64)
				if err != nil {
					color.Red("Failed to parse uint: %v\n", err)
					return
				}
			case runnerDataTypesArrayInt:
				var arr []int
				for _, v := range strings.Split(vo, ",") {
					i, err := strconv.Atoi(v)
					if err != nil {
						color.Red("Failed to parse int: %v\n", err)
						return
					}
					arr = append(arr, i)
				}
				data[kv[0]] = arr
			case runnerDataTypesArrayString:
				data[kv[0]] = strings.Split(vo, ",")
			case runnerDataTypesArrayBool:
				var arr []bool
				for _, v := range strings.Split(vo, ",") {
					b, err := strconv.ParseBool(v)
					if err != nil {
						color.Red("Failed to parse bool: %v\n", err)
						return
					}
					arr = append(arr, b)
				}
				data[kv[0]] = arr
			case runnerDataTypesArrayFloat:
				var arr []float64
				for _, v := range strings.Split(vo, ",") {
					f, err := strconv.ParseFloat(v, 64)
					if err != nil {
						color.Red("Failed to parse float: %v\n", err)
						return
					}
					arr = append(arr, f)
				}
				data[kv[0]] = arr
			case runnerDataTypesArrayUint:
				var arr []uint64
				for _, v := range strings.Split(vo, ",") {
					u, err := strconv.ParseUint(v, 10, 64)
					if err != nil {
						color.Red("Failed to parse uint: %v\n", err)
						return
					}
					arr = append(arr, u)
				}
				data[kv[0]] = arr
			default:
				color.Red("Invalid data type: %s\n", vt)
				return
			}
		}

		if err := runner.Run(ctr, runnerFile, data); err != nil {
			color.Red("Failed to run the load test: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&runnerFile, "file", "f", "", "The file to run the load test")
	runCmd.Flags().StringArrayVarP(&runnerData, "data", "d", []string{}, "The data to run the load test")
}
