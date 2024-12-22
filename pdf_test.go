package pdf

import (
	"os"
	"testing"
)

func TestPDFCreation(t *testing.T) {
	pdfDoc := NewPDF(595.28, 841.89, "pt", 72, 72, 72, 72) // A4 size
	pdfDoc.AddPage()
	pdfDoc.AddText(100, 700, "F1", 14, "Test PDF Creation")
	pdfDoc.DrawRectangle(50, 50, 200, 100, "0 0 0", true)
	pdfDoc.SetMetadata(Metadata{
		Title:   "Test PDF",
		Author:  "Test Author",
		Subject: "Test Subject",
	})

	err := pdfDoc.Save("test_output.pdf")
	if err != nil {
		t.Fatalf("Failed to save PDF: %v", err)
	}

	// Check if file exists
	_, err = os.Stat("test_output.pdf")
	if os.IsNotExist(err) {
		t.Fatalf("PDF file was not created")
	}

	// Clean up
	os.Remove("test_output.pdf")
}

func TestHeaderFooter(t *testing.T) {
	pdfDoc := NewPDF(595.28, 841.89, "pt")
	pdfDoc.AddHeader("Test Header")
	pdfDoc.AddFooter("Test Footer")

	if pdfDoc.headers != "Test Header" {
		t.Errorf("Expected header 'Test Header', got %s", pdfDoc.headers)
	}

	if pdfDoc.footers != "Test Footer" {
		t.Errorf("Expected footer 'Test Footer', got %s", pdfDoc.footers)
	}
}
