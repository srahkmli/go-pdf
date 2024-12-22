package pdf

import (
	"fmt"
)

// DrawRectangle draws a rectangle on the page.
func (pdf *PDF) DrawRectangle(x, y, width, height float64, color string, fill bool) {
	// Assuming color is in RGB format "r g b"
	rect := fmt.Sprintf("%f %f %f %f re", x, y, width, height)
	if fill {
		rect = fmt.Sprintf("%s f", rect)
	} else {
		rect = fmt.Sprintf("%s S", rect)
	}
	pdf.AddObject(rect)
}

// DrawText writes text at a specific position on the page.
func (pdf *PDF) DrawText(x, y float64, text string) {
	// Placeholder for drawing text logic
	pdf.AddObject(fmt.Sprintf("BT /F1 12 Tf %f %f Td (%s) Tj ET", x, y, text))
}
