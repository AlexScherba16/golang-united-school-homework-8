package operations

import "io"

type OperationAlgo interface {
	Process(args map[string]string, writer io.Writer) error
}
