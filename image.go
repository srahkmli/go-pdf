package pdf

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// EmbedImage embeds an image into the PDF.
func (pdf *PDF) EmbedImage(x, y, width, height float64, imgPath string) error {
	imageObj := fmt.Sprintf("<< /Type /XObject /Subtype /Image /Width %.2f /Height %.2f /ColorSpace /DeviceRGB /BitsPerComponent 8 /Filter /DCTDecode /Length ...>>", width, height)
	// Placeholder for actual image embedding logic.
	pdf.AddObject(imageObj)
	return nil
}

// AddImage embeds an image into the PDF document.
func (pdf *PDF) AddImage(imagePath, imageID string) error {
	// Read image file
	imgData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return fmt.Errorf("failed to read image: %v", err)
	}

	// Encode image data in Base64
	encodedImage := base64.StdEncoding.EncodeToString(imgData)

	// Store encoded image in the map
	pdf.images[imageID] = encodedImage
	return nil
}
