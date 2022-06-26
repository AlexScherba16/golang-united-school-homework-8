package implementations

import (
	"errors"
	"fmt"
	"homework_8/internal/utils"
	"io"
	"os"
)

type RemoveOperation struct{}

func (lo *RemoveOperation) Process(args map[string]string, writer io.Writer) error {
	id := args[utils.IdArgument]
	if err := utils.IsValidIdArgument(id); err != nil {
		return err
	}

	fileName := args[utils.FileNameArgument]
	if err := utils.IsValidFileNameArgument(fileName); err != nil {
		return err
	}

	if utils.IsFileExists(fileName) == false {
		return errors.New("file doesn't exist")
	}

	return processRemoveById(args, writer)
}

func processRemoveById(args map[string]string, writer io.Writer) error {
	users, err := utils.GetUsersFromFile(args[utils.FileNameArgument])
	if err != nil {
		return err
	}

	resultUsers := utils.SkipUserById(users, args[utils.IdArgument])
	if len(users) == len(resultUsers) {
		notFoundMessage := []byte(fmt.Sprintf("Item with id %s not found", string(args[utils.IdArgument])))
		writer.Write(notFoundMessage)
		return nil
	}

	return writeDeletedUsersToFile(args[utils.FileNameArgument], resultUsers)
}

func writeDeletedUsersToFile(fileName string, users []utils.User) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, utils.FilePermission)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Seek(0, io.SeekStart)
	file.Truncate(0)

	_, err = file.Write([]byte("["))
	if err != nil {
		return err
	}

	for _, user := range users {
		bytes, err := utils.UserToBytes(user)
		if err != nil {
			return err
		}

		_, err = file.Write(bytes)
		if err != nil {
			return err
		}
	}
	_, err = file.Write([]byte("]"))
	if err != nil {
		return err
	}

	return nil
}
