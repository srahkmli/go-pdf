package pdf

// SetMetadata sets metadata for the PDF document.
func (pdf *PDF) SetMetadata(meta Metadata) {
	pdf.metadata = meta
}
