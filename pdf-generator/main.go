package main

import (
	"fmt"
	"pdfGenerator/htmlParser"
	"pdfGenerator/pdfGen"

	"os"
)

type Data struct {
	Name string
	Date string
}

func main() {

	h := htmlParser.New("tmp")
	wk := pdfGen.NewWkHtmlToPDF("tmp")

	dataHTML := Data{
		Name: "Pedro",
		Date: "06/02/2024",
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer os.Remove(htmlGenerated)
	fmt.Println("HTML gerado", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("PDF gerado", filePDFName)

}
