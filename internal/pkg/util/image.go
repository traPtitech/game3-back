package util

import (
	"bytes"
	"golang.org/x/image/draw"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"math"
)

// EncodeImageToPNG は画像データをPNG形式にエンコードします。
func EncodeImageToPNG(img io.Reader) ([]byte, error) {
	srcImg, _, err := image.Decode(img)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, srcImg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ResizeImageMaintainingAspectRatio は画像を最大幅と最大高さを保ちつつアスペクト比を維持してリサイズします。
// 画像が最大サイズを超えていない場合はそのまま返します。
func ResizeImageMaintainingAspectRatio(imageBytes []byte, maxWidth, maxHeight int) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	srcBounds := img.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()

	// 画像が最大サイズを超えていない場合はそのまま返す
	if srcWidth <= maxWidth && srcHeight <= maxHeight {
		return imageBytes, nil
	}

	widthRatio := float64(maxWidth) / float64(srcWidth)
	heightRatio := float64(maxHeight) / float64(srcHeight)
	ratio := math.Min(widthRatio, heightRatio)

	newWidth := int(float64(srcWidth) * ratio)
	newHeight := int(float64(srcHeight) * ratio)

	return resizeImage(img, newWidth, newHeight)
}

func resizeImage(img image.Image, width, height int) ([]byte, error) {
	dstImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(dstImg, dstImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	var buf bytes.Buffer
	if err := png.Encode(&buf, dstImg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
