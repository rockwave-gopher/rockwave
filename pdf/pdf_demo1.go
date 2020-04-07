package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)
func main() {



	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Chinese","","pdf/MicrosoftYahei.ttf")
	pdf.SetFont("Chinese", "", 12)
	pdf.Cell(40, 10, "阳光保险集团")
	err := pdf.OutputFileAndClose("pdf/pdf.pdf")
	if err !=nil{
		fmt.Println("error in", err)
	}

}
