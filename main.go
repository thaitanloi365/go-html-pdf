package main

import (
	"bytes"
	"html/template"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type InvoiceBillingData struct {
	ContactName string
	CompanyName string
	Address     string
	Phone       string
	Email       string
}
type InvoiceItem struct {
	Description string
	Quantity    string
	UnitPrice   string
	Total       string
	IsLast      bool
}

type InvoiceData struct {
	InvoiceDate              string
	InvoiceNumber            string
	Billing                  InvoiceBillingData
	Shipping                 InvoiceBillingData
	Items                    []InvoiceItem
	SubTotal                 string
	Discount                 string
	SubTotalLessThanDiscount string
	TaxRate                  string
	TotalTax                 string
	ShippingOrHandlingFee    string
	BalanceDue               string
}

func main() {
	//html template path
	var templatePath = "sample.html"

	//path for download pdf
	var outputPath = "example.pdf"

	//html template data
	var templateData = InvoiceData{
		InvoiceDate:   "1/1/2021",
		InvoiceNumber: "123456789",
		Billing: InvoiceBillingData{
			ContactName: "Loi",
			CompanyName: "Relia",
			Address:     "95 Nguyen Tat Thanh",
			Phone:       "+123456789",
			Email:       "thaitanloi365@gmail.com",
		},
		Shipping: InvoiceBillingData{
			ContactName: "Loi 2",
			CompanyName: "Relia",
			Address:     "607 Nguyen Kiem",
			Phone:       "+123456789",
			Email:       "thaitanloi365@gmail.com",
		},
		Items: []InvoiceItem{
			{
				Description: "Label - NAME OF 3PL - DATE",
				Quantity:    "x1",
				UnitPrice:   "$20",
				Total:       "$20",
			},
			{
				Description: "Label - NAME OF 3PL - DATE",
				Quantity:    "x1",
				UnitPrice:   "$20",
				Total:       "$20",
			},
			{
				Description: "Label - NAME OF 3PL - DATE",
				Quantity:    "x1",
				UnitPrice:   "$20",
				Total:       "$20",
				IsLast:      true,
			},
		},
		SubTotal:                 "$10",
		Discount:                 "$10",
		SubTotalLessThanDiscount: "$10",
		TaxRate:                  "$10",
		TotalTax:                 "$10",
		ShippingOrHandlingFee:    "$10",
		BalanceDue:               "$10",
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
