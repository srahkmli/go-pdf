package pdf

// AddTextInColumns adds text to the PDF with multi-column layout.
func (pdf *PDF) AddTextInColumns(text string, columns int, columnWidth float64, columnHeight float64) {
	columnSpacing := 5.0
	xPos := pdf.marginLeft
	yPos := pdf.pageHeight - pdf.marginTop

	// Calculate text per column
	columnTextLength := len(text) / columns
	for col := 0; col < columns; col++ {
		startIdx := col * columnTextLength
		endIdx := (col + 1) * columnTextLength
		if endIdx > len(text) {
			endIdx = len(text)
		}

		columnText := text[startIdx:endIdx]
		// Draw the text in the current column
		pdf.DrawText(xPos, yPos, columnText)

		// Adjust position for next column
		xPos += columnWidth + columnSpacing
	}
}

// AddGridLayout creates a grid of text or images on the page.
func (pdf *PDF) AddGridLayout(rows, cols int, cellWidth, cellHeight float64, content []string) {
	xPos := pdf.marginLeft
	yPos := pdf.pageHeight - pdf.marginTop

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			index := row*cols + col
			if index < len(content) {
				// Draw content (text or image) in the current grid cell
				pdf.DrawText(xPos, yPos, content[index])
			}
			// Adjust position for next cell
			xPos += cellWidth
		}
		// Move to the next row
		xPos = pdf.marginLeft
		yPos -= cellHeight
	}
}
