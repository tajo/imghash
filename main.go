package main

import (
	"crypto/sha256"
	"encoding/hex"
	"image/png"
	"os"
)

func QuantizeColor(r, g, b uint32) (uint8, uint8, uint8) {
	const levels = 2
	const scale = 256 / levels
	return uint8((r >> 8) / scale * scale),
		uint8((g >> 8) / scale * scale),
		uint8((b >> 8) / scale * scale)
}

// HashPNG loads a PNG and computes a hash quantizing colors
func HashPNG(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return "", err
	}
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	hasher := sha256.New()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x+bounds.Min.X, y+bounds.Min.Y).RGBA()
			r8, g8, b8 := QuantizeColor(r, g, b)
			hasher.Write([]byte{r8, g8, b8})
		}
	}

	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}
	numbers := []string{"1", "2"}

	for _, letter := range letters {
		for _, number := range numbers {
			filename := "fixtures/" + letter + number + ".png"
			hash, err := HashPNG(filename)
			if err != nil {
				panic(err)
			}
			println("Hash for "+filename+":", hash)
		}
	}
}
