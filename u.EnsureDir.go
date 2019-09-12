package utils

import "os"

// EnsureDir checks to see if directory exists, if not create it.
func EnsureDir(dir string) (os.FileInfo, error) {
	_path := PathTo(dir)
	ok, err := PathExists(_path)

	if err != nil {
		return nil, err
	}

	if ok == true {
		return os.Stat(PathTo(dir))
	}

	if err := os.MkdirAll(_path, 0700); err != nil {
		return nil, err
	}

	return os.Stat(PathTo(dir))
}
