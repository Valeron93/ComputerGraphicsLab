package lab

import (
	"image"
	"image/color"
	"math/rand"
)

func Lab4(path string, dataset []Dot[int]) {

	groups := FindGroups(dataset)

	r := rand.New(rand.NewSource(674))

	points := AveragePointsGroups(groups)
	colors := make(map[Dot[int]]color.RGBA)

	for _, element := range points {
		colors[element] = RandomColor(r)
	}

	img := BuildVoronoiDiagram(colors, 960, 540)

	for _, point := range dataset {
		DrawPointAlpha(point, img)
	}

	SaveImage(img, path)
}

func BuildVoronoiDiagram(colors map[Dot[int]]color.RGBA, w, h int) *image.RGBA {

	img := CreateImage(960, 540, color.RGBA{0,0,0,0})

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			color := closestColor(colors, Dot[int]{i, j})
			img.SetRGBA(i, j, color)
		}
	}

	for point := range colors {
		DrawCircle(point, 5, color.RGBA{0,0,0,255}, img)
	}

	return img
}


func DrawPointAlpha(point Dot[int], img *image.RGBA) {

	x := point.X
	y := point.Y

	color := img.RGBAAt(x, y)

	color.R = uint8(float32(color.R) * 0.9)
	color.G = uint8(float32(color.G) * 0.9)
	color.B = uint8(float32(color.B) * 0.9)

	img.SetRGBA(x, y, color)
}


func DrawCircle(center Dot[int], r int, fill color.RGBA, img *image.RGBA) {

	x0, y0 := center.X, center.Y

	x, y := r-1, 0

	dx, dy := 1, 1

    err := dx - (r * 2)

    for x >= y {
        img.SetRGBA(x0+x, y0+y, fill)
		img.SetRGBA(x0+y, y0+x, fill)
        img.SetRGBA(x0-x, y0+y, fill)
        img.SetRGBA(x0-y, y0+x, fill)
        img.SetRGBA(x0+x, y0-y, fill)
        img.SetRGBA(x0+y, y0-x, fill)
        img.SetRGBA(x0-y, y0-x, fill)
        img.SetRGBA(x0-x, y0-y, fill)

        if err <= 0 {
            y++
            err += dy
            dy += 2
        }
        if err >= 0 {
            x--
            dx += 2
            err += dx - (r * 2)
        }
    }
}

func closestColor(colors map[Dot[int]]color.RGBA, point Dot[int]) color.RGBA {

	const MaxInt = int((^uint(0)) >> 1)

	var closest Dot[int]

	distance := MaxInt

	for key := range colors {
		if sqrDistance(point, key) < distance {
			closest = key
			distance = sqrDistance(point, key)
		}
	} 

	return colors[closest]
}

func sqrDistance(a, b Dot[int]) int {
	x := a.X - b.X
	y := a.Y - b.Y

	return x * x + y * y

}