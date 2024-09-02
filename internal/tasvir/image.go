package tasvir

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
)

func CompressImage(img image.Image) io.ReadWriter {
	buf := &bytes.Buffer{}
	rw := bufio.NewReadWriter(bufio.NewReader(buf), bufio.NewWriter(buf))

	err := jpeg.Encode(rw, img, &jpeg.Options{Quality: 10})
	if err != nil {
		log.Fatalf("failed to write to image: %v", err)
	}

	return rw
}

func DownsizeJPG(src, dst string) {
	fileSrc, err := os.Open(src)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	org, format, err := image.Decode(fileSrc)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	fmt.Println(format)

	out, err := os.OpenFile(dst, os.O_CREATE, 0777)
	if err != nil {
		log.Fatalf("failed to open dst image: %v", err)
	}

	err = jpeg.Encode(out, org, &jpeg.Options{Quality: 20})
	if err != nil {
		log.Fatalf("failed to write to image: %v", err)
	}
}
