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
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	// get path
	fmt.Printf("Enter path to image folder/directory and press enter: \n>>> ")
	in := bufio.NewReader(os.Stdin)
	imagePath, err := in.ReadString('\n')
	if err != nil {
		log.Fatal("image path not found", err)
	}

	imagePath = strings.TrimSpace(imagePath)

	// create output folder if not exist
	// or check bucket
	store := storage.MakeBucket("output", storage.LOCAL)

	// walk path
	var allImages []string
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
		err = store.Save(img, tasvir.CompressImage(imgFile))
		if err != nil {
			log.Println("err in Save :img ", img)
		}
	}

	var b data.Bucket
	b.Name = "Test"
	b.Images = allImages
	b.VendorID = "vendor1"
	b.ClientID = "client1"
	b.CreatedAt = time.Now()

	url := "http://localhost:8080/update"

	postData, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}
	web.WebRequest(url, string(postData))

}
