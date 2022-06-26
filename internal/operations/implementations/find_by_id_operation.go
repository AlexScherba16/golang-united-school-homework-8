package implementations

import (
	"errors"
	"homework_8/internal/utils"
	"io"
)

type FindByIdOperation struct{}

func (add *FindByIdOperation) Process(args map[string]string, writer io.Writer) error {
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

	return processFindById(args, writer)
}

func processFindById(args map[string]string, writer io.Writer) error {
	users, err := utils.GetUsersFromFile(args[utils.FileNameArgument])
	if err != nil {
		return err
	}

	result := []byte("")
	for _, user := range users {
		if user.Id == args[utils.IdArgument] {
			result, err = utils.UserToBytes(user)
			if err != nil {
				return err
			}
			break
		}
	}
	writer.Write(result)
	return nil
}
