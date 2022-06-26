package main

import (
	"flag"
	"homework_8/internal/operations"
	impl "homework_8/internal/operations/implementations"
	u "homework_8/internal/utils"
	"io"
	"os"
)

type Arguments map[string]string

func parseArgs() (args Arguments) {
	id := flag.String(u.IdArgument, "", "item id")
	operation := flag.String(u.OperationArgument, "", "operation for processing")
	item := flag.String(u.ItemArgument, "", "item for processing")
	fileName := flag.String(u.FileNameArgument, "", "file name for processing")

	flag.Parse()

	return Arguments{
		u.IdArgument:        *id,
		u.OperationArgument: *operation,
		u.ItemArgument:      *item,
		u.FileNameArgument:  *fileName,
	}
}

func Perform(args Arguments, writer io.Writer) error {
	processors := map[string]operations.OperationAlgo{
		u.ListOperation:     &impl.ListOperation{},
		u.AddOperation:      &impl.AddOperation{},
		u.FindByIdOperation: &impl.FindByIdOperation{},
		u.RemoveOperation:   &impl.RemoveOperation{},
	}

	operation := args[u.OperationArgument]
	if err := u.IsValidOperationArgument(operation); err != nil {
		return err
	}

	if operation == u.ListOperation {
		if err := processors[u.ListOperation].Process(args, writer); err != nil {
			return err
		}
	}

	if operation == u.AddOperation {
		if err := processors[u.AddOperation].Process(args, writer); err != nil {
			return err
		}
	}

	if operation == u.FindByIdOperation {
		if err := processors[u.FindByIdOperation].Process(args, writer); err != nil {
			return err
		}
	}

	if operation == u.RemoveOperation {
		if err := processors[u.RemoveOperation].Process(args, writer); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
