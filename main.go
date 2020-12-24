package main

import (
	"bytes"
	"html/template"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	//html template path
	templatePath := "sample.html"

	//path for download pdf
	outputPath := "example.pdf"

	//html template data
	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}
	var buffer = bytes.Buffer{}

	err = t.Execute(&buffer, templateData)
	if err != nil {
		panic(err)
	}

	pdfg := wkhtmltopdf.NewPDFPreparer()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(buffer.Bytes())))
	pdfg.Dpi.Set(800)

	// The contents of htmlsimple.html are saved as base64 string in the JSON file
	jb, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	// Server code
	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jb))
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	pdfgFromJSON.WriteFile(outputPath)

}
