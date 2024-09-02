package bucket

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

const AMZ = "S3"
const LOCAL = "local"

type Bucket struct {
	Path string
	Type string
}

func MakeBucket(path, bType string) Bucket {
	if bType == LOCAL {
		if ok, _ := exists(path); !ok {
			// log.Fatal("image path does not exist")
			if err := os.Mkdir("output", 0777); err != nil {
				log.Fatal("Error creating output dir")
			}
		}
	}

	return Bucket{Path: path, Type: bType}
}

func (b Bucket) Save(img string, w io.ReadWriter) error {
	outputPath := filepath.Join("output", "_"+filepath.Base(img))
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
