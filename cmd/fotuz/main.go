package main

import (
	"bufio"
	"fmt"
	"image"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/KunalDuran/image-reviewer/internal/bucket"
	"github.com/KunalDuran/image-reviewer/internal/data"
	"github.com/KunalDuran/image-reviewer/internal/tasvir"
)

func main() {

	data.InitDB("")

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

	fmt.Println(imagePath)

	// create output folder if not exist
	// or check bucket
	bucket := bucket.MakeBucket("output", bucket.LOCAL)

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
		err = bucket.Save(img, tasvir.CompressImage(imgFile))
		if err != nil {
			log.Println("err in Save :img ", img)
		}

		// save to DB
		data.InsertOne(data.COLLECTION_BUCKET, allImages)

	}

}
