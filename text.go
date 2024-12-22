package pdf

import "fmt"

// AddText writes text at a specific position on the page.
func (pdf *PDF) AddText(x, y float64, font string, size float64, text string) {
	// Using Helvetica as default font
	textObject := fmt.Sprintf("BT /%s %f Tf %f %f Td (%s) Tj ET", font, size, x, y, text)
	pdf.AddObject(textObject)
}
