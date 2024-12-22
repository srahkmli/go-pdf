package main

import (
	"fmt"
	"github.com/srahkmli/go-pdf"
)

func multiColumnText() {
	// Create a new PDF with specified size and margins
	newPDF := pdf.NewPDF(595, 842, "pt", 40, 40, 40, 40)

	// Set metadata
	newPDF.SetMetadata(pdf.Metadata{
		Title:   "Sample PDF",
		Author:  "John Doe",
		Subject: "Advanced Layout Example",
	})

	// Add an image
	//err := newPDF.AddImage("example_image.jpg", "image1")
	//if err != nil {
	//	fmt.Println("Error adding image:", err)
	//	return
	//}

	// Add a page
	newPDF.AddPage()

	// Add a header and footer
	newPDF.AddHeader("Sample Header")
	newPDF.AddFooter("Sample Footer")

	// Add multi-column text
	text := "This is a sample text that will span across multiple columns."
	newPDF.AddTextInColumns(text, 2, 250, 600)

	// Add a grid layout
	gridContent := []string{"Cell 1", "Cell 2", "Cell 3", "Cell 4"}
	newPDF.AddGridLayout(2, 2, 200, 100, gridContent)

	// Save the PDF to a file
	err := newPDF.Save("output.pdf")
	if err != nil {
		fmt.Println("Error saving PDF:", err)
	}
}
