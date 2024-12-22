package main

import (
	"github.com/srahkmli/go-pdf"
	"log"
)

func main() {
	pdfDoc := pdf.NewPDF(595.28, 841.89, "pt", 72, 72, 72, 72) // A4 size
	pdfDoc.AddPage()
	pdfDoc.AddText(100, 700, "F1", 14, "Hello, World!")
	pdfDoc.DrawRectangle(50, 50, 200, 100, "0 0 0", true)
	pdfDoc.AddHeader("Example Header")
	pdfDoc.AddFooter("Page 1")
	pdfDoc.SetMetadata(pdf.Metadata{
		Title:    "Example PDF",
		Author:   "Your Name",
		Subject:  "Demonstration",
		Keywords: "example, pdf, golang",
		Creator:  "Your Project",
	})

	err := pdfDoc.Save("example.pdf")
	if err != nil {
		log.Fatalf("Failed to save PDF: %v", err)
	}

	log.Println("PDF created successfully!")
}
