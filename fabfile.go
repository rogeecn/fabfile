package fabfile

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func Find(dst string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	osLen := 1
	if runtime.GOOS == "windows" {
		osLen = 3
	}

	for {
		if len(path) == osLen {
			return "", errors.New("path not exists")
		}

		fp := filepath.Join(path, dst)
		if _, err := os.Stat(fp); err == nil {
			return fp, nil
		}
		path = filepath.Dir(path)
	}
}

func MustFind(dst string) string {
	if path, err := Find(dst); err != nil {
		panic(err)
	} else {
		return path
	}
}

func Read(file string) ([]byte, error) {
	fp, err := Find(file)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(fp)
}

func MustRead(file string) []byte {
	b, err := Read(file)
	if err != nil {
		panic(err)
	}

	return b
}
