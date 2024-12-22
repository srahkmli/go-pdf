# `go-newPDF` - PDF Generator in Go

`go-newPDF` is a Go library for creating PDF documents programmatically. It supports basic PDF functionalities such as adding text, images, tables, and custom layouts with adjustable margins. It also includes advanced features like text justification and custom table layouts.

## Table of Contents
1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
   - [Creating a New PDF](#creating-a-new-newPDF)
   - [Adding Text](#adding-text)
   - [Justified Text](#justified-text)
   - [Adding Tables](#adding-tables)
   - [Headers and Footers](#headers-and-footers)
   - [Saving the PDF](#saving-the-newPDF)
4. [Code Explanation](#code-explanation)
   - [NewPDF](#newpdf)
   - [AddText](#addtext)
   - [AddTextJustified](#addtextjustified)
   - [AddTable](#addtable)
   - [AddTextJustifiedLine](#addtextjustifiedline)
   - [Save](#save)
5. [License](#license)

## Features
- Create a new PDF document with custom dimensions and margins.
- Add pages to the document.
- Add simple and justified text with custom fonts and sizes.
- Add tables with adjustable columns, rows, and custom alignments.
- Manage headers, footers, and fonts.
- Save the PDF to a file.

## Installation

To install the `go-newPDF` package, use the following Go command:

```bash
go get github.com/srahkmli/go-newPDF
```

## Usage

### Creating a New PDF

To create a new PDF document with custom page size, unit of measurement, and margins:

```go
package multicolumnText

import (
	"fmt"
	"github.com/srahkmli/go-newPDF"
)

func multicolumnText() {
	// Create a new PDF document with custom page size and margins
	newPDF := pdfgen.NewPDF(600, 800, "pt", 50, 50, 50, 50)

	// Add a page to the document
	newPDF.AddPage()

	// Save the PDF to a file
	err := newPDF.Save("example.newPDF")
	if err != nil {
		fmt.Println("Error saving PDF:", err)
	}
}
```

### Adding Text

To add text to the PDF:

```go
newPDF.AddText(x, y, font, size, text)
```

- `x` and `y`: Position of the text in the PDF.
- `font`: Font name (e.g., `"F1"`).
- `size`: Font size (e.g., `12`).
- `text`: The text string to be added.

### Justified Text

To add justified text:

```go
newPDF.AddTextJustified(x, y, width, font, size, text)
```

- `width`: Maximum width for the text to fit within.
- `font` and `size`: Font style and size.
- `text`: The text string to be justified.

The function splits the text into lines, adjusts the spacing between words, and ensures the text fills the specified width.

### Adding Tables

To add a table with adjustable rows and columns:

```go
newPDF.AddTable(x, y, tableData, columnWidths, alignments)
```

- `tableData`: A 2D slice of strings containing the table data (rows and columns).
- `columnWidths`: A slice of floats specifying the width of each column.
- `alignments`: A slice of strings for column text alignment (e.g., `["left", "center", "right"]`).

Example of adding a table:

```go
tableData := [][]string{
	{"ID", "Name", "Age"},
	{"1", "John Doe", "30"},
	{"2", "Jane Smith", "25"},
}

columnWidths := []float64{100, 200, 100}
alignments := []string{"left", "center", "right"}

newPDF.AddTable(50, 700, tableData, columnWidths, alignments)
```

### Headers and Footers

To set headers and footers:

```go
newPDF.SetHeader(header)
newPDF.SetFooter(footer)
```

- `header`: The string for the page header.
- `footer`: The string for the page footer.

### Saving the PDF

Once your document is complete, save it to a file:

```go
err := newPDF.Save("output.newPDF")
if err != nil {
	fmt.Println("Error saving PDF:", err)
}
```

## Code Explanation

### `NewPDF`

The `NewPDF` function initializes a new PDF document. It accepts parameters for page width, height, unit of measurement, and margins.

### `AddText`

The `AddText` function adds text to the current page. You can specify the font, size, and position of the text.

### `AddTextJustified`

This function adds text to the page and justifies it across the specified width. It adjusts the spacing between words to ensure the text is evenly distributed.

### `AddTable`

The `AddTable` function generates a table with adjustable columns, rows, and custom alignment for each column. You can customize the widths of the columns and how the text is aligned (left, center, or right).

### `AddTextJustifiedLine`

Helper function used by `AddTextJustified` to handle the justification of each line by calculating the space between words.

### `Save`

The `Save` function writes the generated PDF to a file on disk.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
