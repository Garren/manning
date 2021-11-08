package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sync"
)

type DB struct {
	mutex *sync.Mutex
	store map[string][3]float64
}

func (db *DB) nearest(target [3]float64) string {
	var filename string
	db.mutex.Lock()
	smallest := 1000000.0
	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(db.store, filename)
	db.mutex.Unlock()
	return filename
}

// resize the image using a new width. return an in memory image
func resize(in image.Image, newWidth int) image.NRGBA {
	bds := in.Bounds()
	width := bds.Dx()
	ratio := width / newWidth
	out := image.NewNRGBA(image.Rect(
		bds.Min.X/ratio, bds.Min.X/ratio,
		bds.Max.X/ratio, bds.Max.Y/ratio))

	for y, j := bds.Min.Y, bds.Min.Y; y < bds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bds.Min.X, bds.Min.X; x < bds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j,
				color.NRGBA{
					uint8(r >> 8),
					uint8(g >> 8),
					uint8(b >> 8),
					uint8(a >> 8)})
		}
	}

	return *out
}

// add up each of the rgb channels, then divide by the total numb(db *DB) er of
// pixels to get the average color of the image
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

// TILESDB get rid of warning
var TILESDB map[string][3]float64

func cloneTilesDB() DB {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	tiles := DB{
		store: db,
		mutex: &sync.Mutex{},
	}
	return tiles
}

func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")
	for _, f := range files {
		name := filepath.Join("tiles", f.Name())
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating the TILESDB", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	return db
}

func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(
		sq(p2[0]-p1[0]) +
			sq(p2[1]-p1[1]) +
			sq(p2[2]-p1[2]))
}

func sq(n float64) float64 {
	return n * n
}
