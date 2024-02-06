package pdfGen

type PDFGenInterface interface {
	Create(htmlFile string) (string, error)
}
