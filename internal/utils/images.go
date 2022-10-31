package utils

import (
	"bytes"
	"github.com/issue9/identicon/v2"
	"image/color"
	"image/jpeg"
)

var avatarMaker *identicon.Identicon

func init() {
	avatarMaker = identicon.New(identicon.Style2, 128, color.RGBA{R: 251, G: 146, B: 158, A: 100}, color.RGBA{R: 255, G: 246, B: 246, A: 100})
}
func GenAvatar(username string) *bytes.Buffer {
	img := avatarMaker.Make([]byte(username))
	buffer := bytes.NewBuffer([]byte{})
	jpeg.Encode(buffer, img, &jpeg.Options{})
	return buffer
}
