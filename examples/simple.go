package main

import (
	"github.com/srahkmli/go-pdf"
	"log"
)

func simpleExample() {
	pdfDoc := pdf.NewPDF(595.28, 841.89, "pt", 72, 72, 72, 72) // A4 size
	pdfDoc.AddPage()

	// Add text
	pdfDoc.AddText(100, 700, "Helvetica", 14, "Hello, World!")

	// Draw rectangle
	pdfDoc.DrawRectangle(50, 50, 200, 100, "0 0 0", true)

	// Add header and footer
	pdfDoc.AddHeader("Example Header")
	pdfDoc.AddFooter("Page 1")

	// Set metadata
	pdfDoc.SetMetadata(pdf.Metadata{
		Title:    "Example PDF",
		Author:   "Your Name",
		Subject:  "Demonstration",
		Keywords: "example, pdf, golang",
		Creator:  "Your Project",
	})

	// Save PDF
	err := pdfDoc.Save("example.pdf")
	if err != nil {
		log.Fatalf("Failed to save PDF: %v", err)
	}

	log.Println("PDF created successfully!")
}
