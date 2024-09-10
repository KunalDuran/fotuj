package storage

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const AMZ = "S3"
const LOCAL = "local"

type Storage struct {
	Path string
	Type string
}

func MakeBucket(path, bType string) Storage {
	if bType == LOCAL {
		if ok, _ := exists(path); !ok {
			// log.Fatal("image path does not exist")
			if err := os.MkdirAll(path, 0777); err != nil {
				log.Fatal("Error creating dir: ", path)
			}
		}
	}

	return Storage{Path: path, Type: bType}
}

func (b Storage) Save(img, out string, w io.ReadWriter) error {
	outputPath := filepath.Join(out, filepath.Base(img))
	f, err := os.OpenFile(outputPath, os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	result, err := io.ReadAll(w)
	if err != nil {
		return err
	}

	f.Write(result)

	return nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(path string) error {
	if ok, _ := exists(path); !ok {
		if err := os.MkdirAll(path, 0777); err != nil {
			return fmt.Errorf("Error creating dir: %s, error: %w", path, err)
		}
	}
	return nil
}
