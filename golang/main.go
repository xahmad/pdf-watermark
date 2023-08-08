// Code generated by chatGPT

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: pdfwatermark input.pdf watermark [output.pdf]")
		return
	}

	inputPDF := os.Args[1]
	watermarkText := os.Args[2]
	outputPDF := "wb-" + filepath.Base(inputPDF)

	if len(os.Args) > 3 {
		outputPDF = os.Args[3]
	}

	// Check if the input PDF file is valid
	if !isValidPDF(inputPDF) {
		fmt.Println("Error: Invalid PDF file.")
		return
	}

	// Add watermark to the PDF
	if err := addWatermark(inputPDF, watermarkText, outputPDF); err != nil {
		fmt.Printf("Error adding watermark: %v\n", err)
		return
	}

	fmt.Printf("Watermarked PDF saved as %s\n", outputPDF)
}

func isValidPDF(filePath string) bool {
	config := pdfcpu.NewDefaultConfiguration()
	config.Validate.Command = true

	_, err := api.ReadFile(filePath, config)
	if err != nil {
		return false
	}
	return true
}

func addWatermark(inputPDF, watermarkText, outputPDF string) error {
	config := pdfcpu.NewDefaultConfiguration()

	// Add watermark text to all pages
	onTop := true
	overlayText := pdfcpu.TextStamp{
		Text:      watermarkText,
		FontSize:  48,
		Pos:       pdfcpu.TopLeft,
		OnTop:     onTop,
		Opacity:   0.5,
		FontColor: pdfcpu.NewPdfColorDeviceRGB(255, 0, 0),
	}

	err := api.AddTextStamp([]string{inputPDF}, outputPDF, overlayText, config)
	if err != nil {
		return err
	}
	return nil
}
