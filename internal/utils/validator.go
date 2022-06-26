package utils

import (
	"errors"
	"fmt"
)

func IsValidIdArgument(argument string) error {
	if argument == "" {
		return errors.New("-id flag has to be specified")
	}
	return nil
}

func IsValidOperationArgument(argument string) error {
	if argument == "" {
		return errors.New("-operation flag has to be specified")
	}
	if isAvailableOperation(argument) != true {
		return fmt.Errorf("Operation %s not allowed!", argument)
	}
	return nil
}

func IsValidFileNameArgument(argument string) error {
	if argument == "" {
		return errors.New("-fileName flag has to be specified")
	}
	return nil
}

func IsValidItemArgument(argument string) error {
	if argument == "" {
		return errors.New("-item flag has to be specified")
	}
	return nil
}

func isAvailableOperation(op string) bool {
	var operations []string = []string{ListOperation, AddOperation, RemoveOperation, FindByIdOperation}
	for _, operation := range operations {
		if op == operation {
			return true
		}
	}
	return false
}
