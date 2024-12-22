package pdf

import (
	"bytes"
	"fmt"
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
	fonts           map[string]string
	images          map[string]string
	metadata        Metadata
	headers         string
	footers         string
}

// Metadata holds document metadata.
type Metadata struct {
	Title    string
	Author   string
	Subject  string
	Keywords string
	Creator  string
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

// AddPage starts a new page.
func (pdf *PDF) AddPage() {
	page := fmt.Sprintf("<< /Type /Page /Parent 1 0 R /MediaBox [0 0 %f %f] >>", pdf.pageWidth, pdf.pageHeight)
	pdf.AddObject(page)
	pdf.currentPage++
}

// AddObject adds a new object to the PDF.
func (pdf *PDF) AddObject(obj string) int {
	objectID := len(pdf.objects) + 1
	pdf.objects = append(pdf.objects, obj)
	return objectID
}

// AddHeader sets the header text for the PDF.
func (pdf *PDF) AddHeader(header string) {
	pdf.headers = header
}

// AddFooter sets the footer text for the PDF.
func (pdf *PDF) AddFooter(footer string) {
	pdf.footers = footer
}

// Save writes the PDF to the specified file.
func (pdf *PDF) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	pdf.writeHeader()
	pdf.writeObjects()
	pdf.writeXref()
	pdf.writeTrailer()
	pdf.writeEOF()

	_, err = file.Write(pdf.buffer.Bytes())
	return err
}

func (pdf *PDF) writeObjects() {
	for i, obj := range pdf.objects {
		offset := pdf.buffer.Len()
		pdf.objectOffsets = append(pdf.objectOffsets, offset)
		pdf.buffer.WriteString(fmt.Sprintf("%d 0 obj\n%s\nendobj\n", i+1, obj))
	}
}

func (pdf *PDF) writeXref() {
	pdf.buffer.WriteString("xref\n")
	pdf.buffer.WriteString(fmt.Sprintf("0 %d\n", len(pdf.objects)+1)) // Including the root object
	pdf.buffer.WriteString("0000000000 65535 f \n")
	for _, offset := range pdf.objectOffsets {
		pdf.buffer.WriteString(fmt.Sprintf("%010d 00000 n \n", offset))
	}
}

func (pdf *PDF) writeTrailer() {
	pdf.buffer.WriteString(fmt.Sprintf("trailer\n<< /Size %d /Root 1 0 R >>\n", len(pdf.objects)+1))
}

func (pdf *PDF) writeEOF() {
	pdf.buffer.WriteString(fmt.Sprintf("startxref\n%d\n%%EOF\n", pdf.buffer.Len()))
}

func (pdf *PDF) writeHeader() {
	pdf.buffer.WriteString("%PDF-1.7\n")
}
