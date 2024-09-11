package main

import (
	"bufio"
	"fmt"
	"image"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/KunalDuran/fotuj/internal/data"
	"github.com/KunalDuran/fotuj/internal/storage"
	"github.com/KunalDuran/fotuj/internal/tasvir"
	"github.com/google/uuid"
)

func ProcessImages() {
	db := data.NewSQLiteDB()
	var b data.Project
	b.Path = prompt("Enter path to image folder/directory: ")
	b.Name = prompt("Enter project name: ")
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	b.Key = GenerateKey()
	b.Link = "http://localhost:8080?key=" + b.Key

	store := storage.Boot(filepath.Join("output", b.Key), storage.LOCAL)

	// walk path
	var err error
	err = filepath.Walk(b.Path, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(strings.ToLower(path), "jpg") || strings.HasSuffix(strings.ToLower(path), "jpeg") {
			var img data.Image
			img.Path = path
			img.Name = filepath.Base(path)
			b.Images = append(b.Images, img)
		}
		return err
	})
	if err != nil {
		log.Fatal("Incorrect Folder Path: ", err)
	}

	// compress image
	for _, img := range b.Images {
		f, err := os.Open(img.Path)
		if err != nil {
			log.Println(img, err)
			continue
		}
		imgFile, _, err := image.Decode(f)
		if err != nil {
			log.Println(img.Path, err)
			continue
		}

		err = store.Save(img.Name, store.Path, tasvir.CompressImage(imgFile))
		if err != nil {
			log.Println("err in Save :img ", err)
		}
	}

	err = db.AddProject(b)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Shareable Link: http://localhost:8080?key=" + b.Key)
}

func ShowProjects() {
	db := data.NewSQLiteDB()
	projects, _ := db.GetProjects()
	for idx, project := range projects {
		fmt.Printf("%d. %s\n\n", idx+1, project.Name)
		fmt.Println("\tKey: ", project.Key)
		fmt.Println("\tLink: ", project.Link)
	}

	for {
		fmt.Println("Enter 0 to exit")
		projIdx := prompt("View Project number: ")

		if projIdx == "0" {
			break
		}

		idx, err := strconv.Atoi(projIdx)
		if err != nil {
			fmt.Println("Incorrect input, please try again.")
			continue
		}

		if idx > len(projects) {
			fmt.Println("Incorrect project number, please try again.")
			continue
		}

		selectedProj := projects[idx-1]
		images, err := db.GetImagesByKey(selectedProj.Key)
		if err != nil {
			log.Fatal(err)
		}

		err = storage.CreatePath(filepath.Join("selections", selectedProj.Name))
		if err != nil {
			log.Fatal(err)
		}

		for _, img := range images {
			if img.Status != 1 {
				continue
			}
			err := os.Link(img.Path, filepath.Join("selections", selectedProj.Name, img.Path))
			if err != nil {
				log.Println("Error linking image: ", img.Path, err)
			}
		}
	}
}

func GenerateKey() string {
	id := uuid.New()
	return id.String()
}

func prompt(msg string) string {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, msg)
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			return input
		}
	}
}
