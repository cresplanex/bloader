package slave

import (
	"context"

	pb "buf.build/gen/go/cresplanex/bloader/protocolbuffers/go/cresplanex/bloader/v1"

	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/output"
	"github.com/cresplanex/bloader/internal/runner"
)

// Output represents the slave output service
type Output struct {
	// OutputID represents the output ID
	OutputID string
	// outputChan represents the output channel
	outputChan chan<- *pb.CallExecResponse
}

// NewSlaveOutput creates a new SlaveOutput
func NewSlaveOutput(outputID string, outputChan chan<- *pb.CallExecResponse) Output {
	return Output{
		OutputID:   outputID,
		outputChan: outputChan,
	}
}

// HTTPDataWriteFactory returns the HTTPDataWrite function
func (o Output) HTTPDataWriteFactory(
	ctx context.Context,
	_ logger.Logger,
	enabled bool,
	uniqueName string,
	header []string,
) (output.HTTPDataWrite, output.Close, error) {
	select {
	case <-ctx.Done():
		return nil, nil, nil
	case o.outputChan <- &pb.CallExecResponse{
		OutputId:   o.OutputID,
		OutputType: pb.CallExecOutputType_CALL_EXEC_OUTPUT_TYPE_HTTP,
		OutputRoot: uniqueName,
		Output: &pb.CallExecResponse_OutputHttp{
			OutputHttp: &pb.CallExecOutputHTTP{
				Data: header,
			},
		},
	}: // do nothing
	}

	return func(
			ctx context.Context,
			_ logger.Logger,
			data []string,
		) error {
			if !enabled {
				return nil
			}
			select {
			case <-ctx.Done():
				return nil
			case o.outputChan <- &pb.CallExecResponse{
				OutputId:   o.OutputID,
				OutputType: pb.CallExecOutputType_CALL_EXEC_OUTPUT_TYPE_HTTP,
				OutputRoot: uniqueName,
				Output: &pb.CallExecResponse_OutputHttp{
					OutputHttp: &pb.CallExecOutputHTTP{
						Data: data,
					},
				},
			}: // do nothing
			}
			return nil
		}, func() error {
			return nil
		}, nil
}

var _ output.Output = Output{}

// OutputFactor represents the factory
type OutputFactor struct {
	outputChan chan<- *pb.CallExecResponse
}

// Factorize returns the factorized output
func (f *OutputFactor) Factorize(_ context.Context, outputID string) (output.Output, error) {
	o := NewSlaveOutput(outputID, f.outputChan)
	return o, nil
}

var _ runner.OutputFactor = &OutputFactor{}
