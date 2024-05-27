package main

import (
	"fmt"
	"os"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println("Usage: colormap <magma,inferno,plasma,viridis,cividis,twilight,turbo> image_name.png")
		return
	}
	cms := map[string][][3]float64{
		"magma":    magma,
		"inferno":  inferno,
		"plasma":   plasma,
		"viridis":  viridis,
		"cividis":  cividis,
		"twilight": twilight,
		"turbo":    turbo,
	}
	cm, ok := cms[os.Args[1]]
	if !ok {
		fmt.Println("Usage: colormap <magma,inferno,plasma,viridis,cividis,twilight,turbo> image_name.png")
		return
	}
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImage(os.Args[2])
	if err != nil {
		fmt.Println("Error loading image: ", err.Error())
		return
	}
	err = mw.TransformImageColorspace(imagick.COLORSPACE_GRAY)
	if err != nil {
		fmt.Println("Failed converting image to grayscale: ", err.Error())
	}
	err = mw.TransformImageColorspace(imagick.COLORSPACE_RGB)
	if err != nil {
		fmt.Println("Failed converting image to rgb: ", err.Error())
	}
	it := mw.NewPixelIterator()
	defer it.Destroy()
	height := mw.GetImageHeight()
	for i := 0; i < int(height); i++ {
		pws := it.GetNextIteratorRow()
		for _, p := range pws {
			if !p.IsVerified() {
				panic("unverified pixel")
			}
			color := int(p.GetBlue() * float64(len(cm)))
			p.SetRed(cm[color][0])
			p.SetGreen(cm[color][1])
			p.SetBlue(cm[color][2])
		}
		if err := it.SyncIterator(); err != nil {
			fmt.Println("Error syncing iterator", err.Error())
			break
		}
	}
	err = mw.WriteImage(os.Args[2] + "_" + os.Args[1] + "_colormap.jpg")
	if err != nil {
		panic(err)
	}
}
