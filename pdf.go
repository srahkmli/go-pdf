package go-pdf

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// PDF represents a PDF document.
type PDF struct {
	buffer          *bytes.Buffer
	objects         []string
	objectOffsets   []int
	currentPage     int
	pageWidth       float64
	pageHeight      float64
	measurementUnit string
	marginTop       float64
	marginBottom    float64
	marginLeft      float64
	marginRight     float64
	headers         string
	footers         string
	fonts           map[string]string
	images          map[string]string
}

// NewPDF initializes a new PDF instance.
func NewPDF(pageWidth, pageHeight float64, unit string, margins ...float64) *PDF {
	var marginTop, marginBottom, marginLeft, marginRight float64
	if len(margins) == 4 {
		marginTop, marginBottom, marginLeft, marginRight = margins[0], margins[1], margins[2], margins[3]
	}
	return &PDF{
		buffer:          new(bytes.Buffer),
		objects:         []string{},
		pageWidth:       pageWidth,
		pageHeight:      pageHeight,
		measurementUnit: unit,
		marginTop:       marginTop,
		marginBottom:    marginBottom,
		marginLeft:      marginLeft,
		marginRight:     marginRight,
		fonts:           make(map[string]string),
		images:          make(map[string]string),
	}
}

// AddObject adds a new object to the PDF.
func (pdf *PDF) AddObject(obj string) int {
	objectID := len(pdf.objects) + 1
	pdf.objects = append(pdf.objects, obj)
	return objectID
}

// AddPage adds a new page to the PDF.
func (pdf *PDF) AddPage() {
	page := fmt.Sprintf("<< /Type /Page /Parent 1 0 R >>")
	pdf.AddObject(page)
	pdf.currentPage++
}

// SetHeader sets the page header.
func (pdf *PDF) SetHeader(header string) {
	pdf.headers = header
}

// SetFooter sets the page footer.
func (pdf *PDF) SetFooter(footer string) {
	pdf.footers = footer
}

// AddText adds text to the current page.
func (pdf *PDF) AddText(x, y float64, font string, size float64, text string) {
	textStream := fmt.Sprintf("BT /F1 %.2f Tf %.2f %.2f Td (%s) Tj ET", size, x, y, text)
	stream := fmt.Sprintf("<< /Length %d >> stream\n%s\nendstream", len(textStream), textStream)
	pdf.AddObject(stream)
}

// AddGradient adds a gradient to the PDF.
func (pdf *PDF) AddGradient(x, y, width, height float64, colorStart, colorEnd string) {
	gradientStream := fmt.Sprintf("q %.2f 0 0 %.2f %.2f %.2f cm /%s Sh Q", width, height, x, y, colorStart+":"+colorEnd)
	stream := fmt.Sprintf("<< /Length %d >> stream\n%s\nendstream", len(gradientStream), gradientStream)
	pdf.AddObject(stream)
}

// AddTextJustified adds justified text to the PDF.
func (pdf *PDF) AddTextJustified(x, y, width float64, font string, size float64, text string) {
	words := splitTextIntoWords(text)
	line := ""
	lineWidth := 0.0

	// Iterate over the words and add them to lines
	for _, word := range words {
		// Estimate the width of the line if this word is added
		newLineWidth := textWidth(line+" "+word, font, size)

		// If the line fits within the width, add the word
		if line == "" || newLineWidth <= width {
			if line != "" {
				line += " "
			}
			line += word
			lineWidth = newLineWidth
		} else {
			// Justify the current line and move to the next
			pdf.AddTextJustifiedLine(x, y, line, lineWidth, width, font, size)
			y -= size * 1.2 // Move to the next line
			line = word
			lineWidth = textWidth(line, font, size)
		}
	}

	// Add the last line (if any)
	if line != "" {
		pdf.AddTextJustifiedLine(x, y, line, lineWidth, width, font, size)
	}
}

// AddTextJustifiedLine adds a single justified line of text to the PDF.
func (pdf *PDF) AddTextJustifiedLine(x, y float64, line string, lineWidth, lineMaxWidth float64, font string, size float64) {
	// Split the line into words
	words := splitTextIntoWords(line)

	// If there's only one word, don't justify
	if len(words) == 1 {
		pdf.AddText(x, y, font, size, line)
		return
	}

	// Calculate total space to distribute
	spaceWidth := textWidth(" ", font, size)
	totalSpaces := len(words) - 1
	remainingWidth := lineMaxWidth - lineWidth
	spaceBetweenWords := remainingWidth / float64(totalSpaces)

	// Calculate the X position for the first word
	lineX := x
	for i, word := range words {
		// Add the word to the line
		pdf.AddText(lineX, y, font, size, word)

		// Move to the next position after the word
		lineX += textWidth(word, font, size)

		// Add space between words (extra space for justification)
		if i < len(words)-1 {
			lineX += spaceBetweenWords + spaceWidth
		}
	}
}

// Save writes the PDF to the specified file.
func (pdf *PDF) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	pdf.buffer.WriteString("%PDF-1.7\n")

	for _, obj := range pdf.objects {
		offset := pdf.buffer.Len()
		pdf.objectOffsets = append(pdf.objectOffsets, offset)
		pdf.buffer.WriteString(fmt.Sprintf("%d 0 obj\n%s\nendobj\n", len(pdf.objectOffsets), obj))
	}

	xrefStart := pdf.buffer.Len()
	pdf.buffer.WriteString("xref\n0 1\n0000000000 65535 f \n")
	for _, offset := range pdf.objectOffsets {
		pdf.buffer.WriteString(fmt.Sprintf("%010d 00000 n \n", offset))
	}

	pdf.buffer.WriteString(fmt.Sprintf("trailer\n<< /Size %d /Root 1 0 R >>\nstartxref\n%d\n%%EOF", len(pdf.objectOffsets), xrefStart))
	_, err = io.Copy(file, pdf.buffer)
	return err
}

// Utility: Split text into words.
func splitTextIntoWords(text string) []string {
	return []string{} // Implement splitting logic
}

// Utility: Calculate text width based on font and size.
func textWidth(text, font string, size float64) float64 {
	return float64(len(text)) * size * 0.5 // Simplified width calculation
}

// Utility: Compress data using zlib.
func compressData(data string) string {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, _ = w.Write([]byte(data))
	w.Close()
	return b.String()
}

// Utility: Base64 encode data.
func base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// AddTable adds a table to the PDF with adjustable rows and columns.
func (pdf *PDF) AddTable(x, y float64, tableData [][]string, columnWidths []float64, alignments []string) {
	// Default font size for table cells
	fontSize := 12.0

	// Loop through the rows
	for _, row := range tableData {
		lineX := x // Starting position of each row

		// Loop through the columns in each row
		for j, cell := range row {
			width := columnWidths[j]
			alignment := "left"
			if len(alignments) > j {
				alignment = alignments[j]
			}

			// Calculate text width for alignment purposes
			textWidth := textWidth(cell, "F1", fontSize)

			// Adjust x position based on alignment
			switch alignment {
			case "left":
				// Left alignment: No change to lineX
				pdf.AddText(lineX, y, "F1", fontSize, fmt.Sprintf("(%s)", cell))
			case "center":
				// Center alignment: Move x to center the text within the column width
				lineX += (width - textWidth) / 2
				pdf.AddText(lineX, y, "F1", fontSize, fmt.Sprintf("(%s)", cell))
			case "right":
				// Right alignment: Move x to the right edge of the column
				lineX += width - textWidth
				pdf.AddText(lineX, y, "F1", fontSize, fmt.Sprintf("(%s)", cell))
			}

			// Move the lineX position for the next cell
			lineX += width
		}

		// Move to the next row
		y -= 14 // Adjust vertical space for next row (based on font size)
	}
}
