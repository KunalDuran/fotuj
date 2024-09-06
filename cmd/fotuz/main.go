package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/KunalDuran/image-reviewer/internal/data"
	"github.com/KunalDuran/image-reviewer/internal/storage"
	"github.com/KunalDuran/image-reviewer/internal/tasvir"
	"github.com/KunalDuran/image-reviewer/internal/web"
	"github.com/google/uuid"
)

type config struct {
	Client    string
	Vendor    string
	ServerURI string
}

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	imagePath := prompt("Enter path to image folder/directory")

	var b data.Bucket
	b.Name = prompt("Enter Project Name")
	b.PhotographerID = prompt("Enter Photographer's Name/ID")
	b.ClientID = prompt("Enter Client's Name")
	b.CreatedAt = time.Now()
	b.Key = GenerateKey()

	// create output folder if not exist
	// or check bucket
	store := storage.MakeBucket(filepath.Join("output", b.Key), storage.LOCAL)

	// walk path
	var allImages []string
	var err error
	err = filepath.Walk(imagePath, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(strings.ToLower(path), "jpg") || strings.HasSuffix(strings.ToLower(path), "jpeg") {
			allImages = append(allImages, path)
		}
		return err
	})
	if err != nil {
		log.Fatal("Incorrect Folder Path: ", err)
	}

	// compress image
	for _, img := range allImages {
		f, err := os.Open(img)
		if err != nil {
			log.Println(img, err)
			continue
		}
		imgFile, _, err := image.Decode(f)
		if err != nil {
			log.Println(img, err)
			continue
		}

		// save output
		err = store.Save(img, store.Path, tasvir.CompressImage(imgFile))
		if err != nil {
			log.Println("err in Save :img ", err)
		}
	}
	for _, i := range allImages {
		var img data.Image
		img.AbsolutePath = i
		img.Path = filepath.Base(i)
		b.Images = append(b.Images, img)
	}

	url := "http://localhost:8080/bucket"

	postData, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}
	_, _, err = web.WebRequest(url, string(postData))
	if err != nil {
		log.Fatal("request failed", err)
	}

	fmt.Println("Shareable Link: http://localhost:8080?key=" + b.Key)
}

func GenerateKey() string {
	id := uuid.New()
	return id.String()
}

func prompt(msg string) string {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, msg+" ")
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			return input
		}
	}
}
