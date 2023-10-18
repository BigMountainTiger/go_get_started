package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const APPNAME = "launch-chromium"

func Home_path() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/.%s", home_dir, APPNAME)
	return path, nil
}

func Save_file(path string, name string, content []byte) (er error) {

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

func Read_file(path string, name string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("%s/%s", path, name))
}

func Remove_file(path string, name string) error {
	return os.Remove(fmt.Sprintf("%s/%s", path, name))
}

func Remove_directory(path string) error {
	return os.RemoveAll(path)
}

func File_info(path string, name string) (fs.FileInfo, error) {
	return os.Stat(fmt.Sprintf("%s/%s", path, name))
}

// Exposed functions
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func Save_file_to_home(name string, content []byte) error {
	path, err := Home_path()
	if err != nil {
		return err
	}

	return Save_file(path, name, content)
}

func Read_file_from_home(name string) ([]byte, error) {

	path, err := Home_path()
	if err != nil {
		return nil, err
	}

	return Read_file(path, name)
}

func Remove_file_from_home(name string) error {
	path, err := Home_path()
	if err != nil {
		return err
	}

	return Remove_file(path, name)
}

func Remove_directory_from_home(path string) error {
	home_path, err := Home_path()
	if err != nil {
		return err
	}

	return Remove_directory(home_path + "/" + path)
}

func File_info_in_home(name string) (fs.FileInfo, error) {
	path, err := Home_path()
	if err != nil {
		return nil, err
	}

	return File_info(path, name)
}
