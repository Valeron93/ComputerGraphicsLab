package lab

import "image/color"

func Lab2(savePath string, dataset []Dot[int]) {
	img := CreateImage(960, 540, color.RGBA{255,255,255,255})

	for _, value := range dataset {
		img.SetRGBA(value.X, value.Y, color.RGBA{0,0,255,255} )
	}

	SaveImage(img, savePath)
}
