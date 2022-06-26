package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const (
	IdArgument        = "id"
	OperationArgument = "operation"
	ItemArgument      = "item"
	FileNameArgument  = "fileName"
)

const (
	ListOperation     = "list"
	AddOperation      = "add"
	RemoveOperation   = "remove"
	FindByIdOperation = "findById"
)

const FilePermission = 0644

func IsFileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func GetUsersFromFile(fileName string) ([]User, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(content, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserFromItem(item string) (User, error) {
	var user User
	err := json.Unmarshal([]byte(item), &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func UserToBytes(item User) ([]byte, error) {
	result, err := json.Marshal(&item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SkipUserById(users []User, id string) []User {
	var result []User
	for _, user := range users {
		if user.Id == id {
			continue
		}
		result = append(result, user)
	}
	return result
}
