package pdfGen

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wk struct {
	rootPath string
}

func NewWkHtmlToPDF(rootPath string) PDFGenInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile string) (string, error) {
	f, err := os.Open(htmlFile)
	if err != nil {
		return "", err
	}
	wkhtmltopdf.SetPath("C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe")

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	if err := pdfg.WriteFile(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
