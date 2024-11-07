package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/mdp/qrterminal/v3"
	goqrcode "github.com/skip2/go-qrcode"
)

// ReadQRCode reads QR code content from an image file
func ReadQRCode(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %v", err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", fmt.Errorf("failed to convert to bitmap: %v", err)
	}

	reader := qrcode.NewQRCodeReader()
	result, err := reader.Decode(bmp, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decode QR code: %v", err)
	}

	return result.String(), nil
}

// GenerateQRCode generates QR code image file
func GenerateQRCode(content string, filename string) error {
	qr, err := goqrcode.New(content, goqrcode.Medium)
	if err != nil {
		return fmt.Errorf("failed to generate QR code: %v", err)
	}

	err = qr.WriteFile(256, filename)
	if err != nil {
		return fmt.Errorf("failed to save QR code image: %v", err)
	}

	return nil
}

// DisplayQRInTerminal displays QR code content in terminal
func DisplayQRInTerminal(content string) {
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(content, config)
}

// DisplayQRFromImage reads and displays QR code from image file in terminal
func DisplayQRFromImage(filename string) error {
	// First read QR code content
	content, err := ReadQRCode(filename)
	if err != nil {
		return err
	}

	// Display QR code content
	fmt.Printf("QR Code Content: %s\n", content)

	// Display QR code in terminal
	DisplayQRInTerminal(content)
	return nil
}

func main() {
	// Define command line parameters
	var (
		content  = flag.String("c", "", "content to generate QR code")
		filename = flag.String("f", "", "QR code image file path (for reading or saving)")
		generate = flag.Bool("g", false, "generate QR code image")
		display  = flag.Bool("d", false, "display QR code")
	)

	flag.Parse()

	// Execute corresponding operation based on parameters
	if *generate && *content != "" {
		// Generate QR code image
		outputFile := "qr.png" // default filename
		if *filename != "" {
			outputFile = *filename
		}

		err := GenerateQRCode(*content, outputFile)
		if err != nil {
			fmt.Printf("Failed to generate QR code: %v\n", err)
			return
		}
		fmt.Printf("Successfully generated QR code image: %s\n", outputFile)

		// Display if needed
		if *display {
			DisplayQRInTerminal(*content)
		}
	} else if *display {
		if *filename != "" {
			// Read and display QR code from file
			err := DisplayQRFromImage(*filename)
			if err != nil {
				fmt.Printf("Failed to display QR code: %v\n", err)
				return
			}
		} else if *content != "" {
			// Directly display content as QR code
			DisplayQRInTerminal(*content)
		} else {
			fmt.Println("Error: Displaying QR code requires content (-c) or file path (-f)")
		}
	} else {
		// Show usage instructions
		fmt.Println("Usage:")
		fmt.Println("  Generate QR code image:")
		fmt.Println("    -g -c \"content\" [-f filename.png]")
		fmt.Println("  Display existing QR code image:")
		fmt.Println("    -d -f filename.png")
		fmt.Println("  Directly display content as QR code:")
		fmt.Println("    -d -c \"content\"")
		fmt.Println("\nExamples:")
		fmt.Println("  Generate QR code image:")
		fmt.Println("    ./qrcode -g -c \"https://example.com\" -f test.png")
		fmt.Println("  Read and display QR code:")
		fmt.Println("    ./qrcode -d -f test.png")
		fmt.Println("  Directly display QR code:")
		fmt.Println("    ./qrcode -d -c \"Hello, World!\"")
	}
}
