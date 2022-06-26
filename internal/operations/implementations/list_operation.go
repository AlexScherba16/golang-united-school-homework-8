package implementations

import (
	"homework_8/internal/utils"
	"io"
	"io/ioutil"
)

type ListOperation struct{}

func (lo *ListOperation) Process(args map[string]string, writer io.Writer) error {
	fileName := args[utils.FileNameArgument]
	if err := utils.IsValidFileNameArgument(fileName); err != nil {
		return err
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	writer.Write(file)
	return nil
}
