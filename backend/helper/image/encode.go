package image

import (
	"bytes"
	"image"
	"image/png"

	"github.com/bsthun/gut"
	"github.com/nfnt/resize"
)

func EncodePhotoToPng(data []byte) ([]byte, *gut.ErrorInstance) {
	// * decode image to get format
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, gut.Err(false, "failed to decode image", err)
	}

	// * get original dimensions
	originalBounds := img.Bounds()
	originalWidth := originalBounds.Dx()
	originalHeight := originalBounds.Dy()

	// * check if resize is needed
	if originalWidth > 1024 {
		// * calculate new height maintaining aspect ratio
		ratio := float64(1024) / float64(originalWidth)
		newHeight := uint(float64(originalHeight) * ratio)

		// * resize image using high quality resampling
		resizedImg := resize.Resize(1024, newHeight, img, resize.Lanczos3)

		// * encode to PNG
		var buf bytes.Buffer
		err = png.Encode(&buf, resizedImg)
		if err != nil {
			return nil, gut.Err(false, "failed to encode resized image to PNG", err)
		}

		return buf.Bytes(), nil
	} else {
		var buf bytes.Buffer
		err = png.Encode(&buf, img)
		if err != nil {
			return nil, gut.Err(false, "failed to encode image to PNG", err)
		}

		return buf.Bytes(), nil
	}
}
