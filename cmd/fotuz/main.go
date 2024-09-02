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

	imagePath = strings.TrimSpace(imagePath)

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
	}

	var b Bucket
	b.Name = "Test"
	b.Images = allImages
	b.VendorID = "vendor1"
	b.ClientID = "client1"
	b.CreatedAt = time.Now()
	data.InsertOne(data.COLLECTION_BUCKET, b)

}

type Bucket struct {
	Name          string              `json:"name" bson:"name"`
	Images        []string            `json:"images" bson:"images"`
	Selected      map[string][]string `json:"selected" bson:"selected"`
	Rejected      map[string][]string `json:"rejected" bson:"rejected"`
	ShareableLink string              `json:"shareable_link" bson:"shareable_link"`
	StorageType   string              `json:"storage_type" bson:"storage_type"`
	VendorID      string              `json:"vendor_id" bson:"vendor_id"`
	ClientID      string              `json:"client_id" bson:"client_id"`
	Selectors     []string            `json:"selectors" bson:"selectors"`
	CreatedAt     time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt     string              `json:"updated_at" bson:"updated_at"`
}
