package lab

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"strings"
)


type Dot[T any] struct {
	X T
	Y T
}


func ReadDataset(path string) []Dot[int] {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	result := make([]Dot[int], 0)

	for scanner.Scan() {

		line := scanner.Text()

		x, y := parseLine(line) 
		
		result = append(result, Dot[int]{x, y})

	}

	return result
}

func CreateImage(w, h int, fill color.RGBA) *image.RGBA {

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.SetRGBA(x, y, fill)	
		}
	}

	return img
}

func SaveImage(img *image.RGBA, path string) {

	file, _ := os.Create(path)
	defer file.Close()

	png.Encode(file, img)
}

func RandomColor(r *rand.Rand) color.RGBA {

	return color.RGBA{
		uint8(r.Int()),
		uint8(r.Int()),
		uint8(r.Int()),
		255,
	}
}

func parseLine(line string) (int, int) {
	arr := strings.Split(line, " ")

	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])

	return x, y
}
