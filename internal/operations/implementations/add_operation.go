package implementations

import (
	"fmt"
	"homework_8/internal/utils"
	"io"
	"os"
)

type AddOperation struct{}

func (add *AddOperation) Process(args map[string]string, writer io.Writer) error {
	item := args[utils.ItemArgument]
	if err := utils.IsValidItemArgument(item); err != nil {
		return err
	}

	fileName := args[utils.FileNameArgument]
	if err := utils.IsValidFileNameArgument(fileName); err != nil {
		return err
	}

	if utils.IsFileExists(fileName) == true {
		return processExistingFile(args, writer)
	}
	return createAndWriteFile(args, writer)
}

func processExistingFile(args map[string]string, writer io.Writer) error {
	itemUser, err := utils.GetUserFromItem(args[utils.ItemArgument])
	if err != nil {
		return err
	}

	users, err := utils.GetUsersFromFile(args[utils.FileNameArgument])
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Id == itemUser.Id {
			err = fmt.Errorf("Item with id %s already exists", string(itemUser.Id))
			writer.Write([]byte(err.Error()))
			break
		}
	}
	return nil
}

func createAndWriteFile(args map[string]string, writer io.Writer) error {
	_ = writer

	file, err := os.OpenFile(args[utils.FileNameArgument], os.O_WRONLY|os.O_CREATE|os.O_APPEND, utils.FilePermission)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write([]byte("[" + args[utils.ItemArgument] + "]"))
	return nil
}
