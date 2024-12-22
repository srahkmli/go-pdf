package pdf

import "fmt"

// AddText adds text to the current page.
func (pdf *PDF) AddText(x, y float64, font string, size float64, text string) {
	textStream := fmt.Sprintf("BT /F1 %.2f Tf %.2f %.2f Td (%s) Tj ET", size, x, y, text)
	stream := fmt.Sprintf("<< /Length %d >> stream\n%s\nendstream", len(textStream), textStream)
	pdf.AddObject(stream)
}
