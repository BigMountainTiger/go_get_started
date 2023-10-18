package utils

import (
	"fmt"
	"log"
	"os"
)

const APPNAME = "oauth2_example"

func home_path() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/.%s", home_dir, APPNAME)
	return path, nil
}

func save_file(path string, name string, content []byte) (er error) {

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/%s", path, name), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
			er = err
		}
	}()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func read_file(path string, name string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", path, name))
}

func remove_file(path string, name string) error {
	return os.Remove(fmt.Sprintf("%s/%s", path, name))
}

// Exposed functions
func Save_file_to_home(name string, content []byte) error {
	path, err := home_path()
	if err != nil {
		return err
	}

	return save_file(path, name, content)
}

func Read_file_from_home(name string) ([]byte, error) {

	path, err := home_path()
	if err != nil {
		return nil, err
	}

	return read_file(path, name)
}

func Remove_file_from_home(name string) error {
	path, err := home_path()
	if err != nil {
		return err
	}

	return remove_file(path, name)
}
