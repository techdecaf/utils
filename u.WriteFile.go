package utils

import (
	"io/ioutil"
	"os"
)

// WriteFile and return as a string
func WriteFile(file, data string) error {
	_path := PathTo(file)
	// try to create file
	if _, err := os.Create(_path); err != nil {
		return err
	}

	return ioutil.WriteFile(_path, []byte(data), 0700)
}
