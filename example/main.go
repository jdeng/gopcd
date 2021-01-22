package main

import (
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/jdeng/gopcd"
)

func main() {
	src := os.Args[1]
	r, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}

	dec := gopcd.NewDecoder(r)
	img, rotate, err := dec.Decode()
	if err != nil || img == nil {
		log.Fatal(err)
	}
	switch rotate {
	case 1:
		img = imaging.Rotate90(img)
	case 2:
		img = imaging.Rotate180(img)
	case 3:
		img = imaging.Rotate270(img)
	}

	outf := filepath.Base(src)
	outf = strings.TrimSuffix(outf, filepath.Ext(src)) + ".jpg"

	log.Printf("Saving to %s\n", outf)
	out, _ := os.Create(outf)
	defer out.Close()

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s saved\n", outf)
}
