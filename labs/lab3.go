package lab

import (
	"image/color"
	"math"
)

func Lab3(savePath string, dataset []Dot[int]) {
	img := CreateImage(960, 960, color.RGBA{255,255,255,255})

	center := Dot[float64] {960.0/2.0, 960.0/2.0}
	
	const n int = 3
	const angle float64 = float64(10 * (n-1)) * math.Pi / 180.0

	for _, value := range dataset {

		point := Dot[float64] {
			float64(value.X), float64(value.Y),
		}

		point = subtract(point, center)
		point = rotate(point, angle)
		point = add(point, center)

		img.SetRGBA(int(point.X), int(point.Y), color.RGBA{0,0,255,255} )

	}

	SaveImage(img, savePath)
}

func add(lhs, rhs Dot[float64]) Dot[float64] {
	return Dot[float64] {
		lhs.X + rhs.X,
		lhs.Y + rhs.Y,
	}
}

func subtract(lhs, rhs Dot[float64]) Dot[float64] {
	return Dot[float64] {
		lhs.X - rhs.X,
		lhs.Y - rhs.Y,
	}
}

func rotate(point Dot[float64], angle float64) Dot[float64] {

	return Dot[float64] {
		point.X * math.Cos(angle) - point.Y * math.Sin(angle),
		point.X * math.Sin(angle) + point.Y * math.Cos(angle),
	}

}