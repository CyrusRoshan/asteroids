package sprites

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
)

func LoadPicture(path string) pixel.Picture {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return pixel.PictureDataFromImage(img)
}

func LoadSprite(path string) *pixel.Sprite {
	pic := LoadPicture(path)

	return pixel.NewSprite(pic, pic.Bounds())
}
