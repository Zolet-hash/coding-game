package main

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func readPdf(filepath string) {
	ctx, err := api.ReadContextFile(filepath, pdfcpu.NewDefaultConfiguration())
	if err != nil {
		fmt.Println("Error reading PDF: ", err)
		return
	}
	for i := 1; i <= ctx.PageCount; i++ {
		page, err := ctx.ExtractPageText(i)
		if err != nil {
			fmt.Printf("Error extracting text from page: %d :%s\n", i, err)
		} else {
			fmt.Printf("Page: %d\n%s\n", i, page)
		}
	}

}

func main() {
	readPdf("/home/zolet-sec/Documents/tax-system.pdf")
}
