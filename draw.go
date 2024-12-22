package pdf

import (
	"fmt"
)

// DrawRectangle draws a rectangle on the PDF.
func (pdf *PDF) DrawRectangle(x, y, width, height float64, color string, fill bool) {
	colorCmd := fmt.Sprintf("%s RG %s rg", color, color)
	shapeCmd := ""
	if fill {
		shapeCmd = fmt.Sprintf("%.2f %.2f %.2f %.2f re f", x, y, width, height)
	} else {
		shapeCmd = fmt.Sprintf("%.2f %.2f %.2f %.2f re S", x, y, width, height)
	}
	cmd := fmt.Sprintf("q\n%s\n%s\nQ", colorCmd, shapeCmd)
	stream := fmt.Sprintf("<< /Length %d >> stream\n%s\nendstream", len(cmd), cmd)
	pdf.AddObject(stream)
}

// DrawText writes text at a specific position on the page.
func (pdf *PDF) DrawText(x, y float64, text string) {
	// Placeholder for drawing text logic
	pdf.AddObject(fmt.Sprintf("BT /F1 12 Tf %f %f Td (%s) Tj ET", x, y, text))
}
