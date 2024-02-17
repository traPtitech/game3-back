package util

import (
	"bytes"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/logica0419/resigif"
	"golang.org/x/image/draw"
	_ "golang.org/x/image/webp"
	"image"
	"image/gif"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"math"
	"mime/multipart"
)

// DecodeImage は画像データをデコードします。
func DecodeImage(srcImage multipart.File) (image.Image, string, error) {
	img, format, err := image.Decode(srcImage)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

// DecodeGifImage はGIF画像データをデコードします。
func DecodeGifImage(srcImage multipart.File) (*gif.GIF, error) {
	gifData, err := gif.DecodeAll(srcImage)
	if err != nil {
		return nil, err
	}

	return gifData, nil
}

// EncodeImageToPNG は画像データをPNG形式にエンコードします。
func EncodeImageToPNG(img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// EncodeGifImage はGIF画像データをエンコードします。
func EncodeGifImage(gifData *gif.GIF) ([]byte, error) {
	var buf bytes.Buffer

	if err := gif.EncodeAll(&buf, gifData); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ResizePngImageMaintainingAspectRatio は画像を最大幅と最大高さを保ちつつアスペクト比を維持してリサイズします。
// 画像が最大サイズを超えていない場合はそのまま返します。
func ResizePngImageMaintainingAspectRatio(imageBytes []byte, maxWidth, maxHeight int) ([]byte, error) {
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

	newWidth, newHeight := calculateResizedSize(srcWidth, srcHeight, maxWidth, maxHeight)
	return resizeImage(img, newWidth, newHeight)
}

// ResizeGifImageMaintainingAspectRatio はGIF画像を最大幅と最大高さを保ちつつアスペクト比を維持してリサイズします。
// 画像が最大サイズを超えていない場合はそのまま返します。
func ResizeGifImageMaintainingAspectRatio(c echo.Context, gifData *gif.GIF, maxWidth, maxHeight int) (*gif.GIF, error) {
	srcWidth, srcHeight, err := getGifImageSize(gifData)
	if err != nil {
		return nil, err
	}

	// 画像が最大サイズを超えていない場合はそのまま返す
	if srcWidth <= maxWidth && srcHeight <= maxHeight {
		return gifData, nil
	}

	return resigif.Resize(c.Request().Context(), gifData, maxWidth, maxHeight)
}

func getGifImageSize(gifData *gif.GIF) (width, height int, err error) {
	if len(gifData.Image) == 0 {
		return 0, 0, errors.New("GIF contains no frames")
	}

	firstFrame := gifData.Image[0]

	bounds := firstFrame.Bounds()

	width = bounds.Dx()
	height = bounds.Dy()

	return width, height, nil
}

func calculateResizedSize(srcWidth, srcHeight, maxWidth, maxHeight int) (int, int) {
	widthRatio := float64(maxWidth) / float64(srcWidth)
	heightRatio := float64(maxHeight) / float64(srcHeight)
	ratio := math.Min(widthRatio, heightRatio)

	newWidth := int(float64(srcWidth) * ratio)
	newHeight := int(float64(srcHeight) * ratio)

	return newWidth, newHeight
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
