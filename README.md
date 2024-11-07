# qr-tools

A command-line QR code generator and reader tool written in Go. This tool allows you to generate QR codes, read QR codes from images, and display QR codes directly in your terminal.

## Features

- Generate QR code images from text/URL content
- Read QR codes from image files (supports PNG, JPEG)
- Display QR codes directly in terminal
- Command-line interface with multiple options

## Installation

```sh
# Clone the repository
git clone https://github.com/walle-soft/qr-tools.git

# Navigate to the directory
cd qr-tools

# Build the project
go build
```

## Usage

### Generate QR Code Image
```sh
./qrcode -g -c "content" [-f filename.png]
```

### Display Existing QR Code Image
```sh
./qrcode -d -f image.png
```

### Directly Display QR Code in Terminal
```sh
./qrcode -d -c "content"
```

### Command Options
- `-g`: Generate QR code image
- `-d`: Display QR code
- `-c`: Content for QR code
- `-f`: File path for input/output image

## Examples

1. Generate a QR code image:
```sh
./qrcode -g -c "https://example.com" -f test.png
```

2. Read and display a QR code from an image:
```sh
./qrcode -d -f test.png
```

3. Display content as QR code in terminal:
```sh
./qrcode -d -c "Hello, World!"
```

## Dependencies

- [github.com/makiuchi-d/gozxing](https://github.com/makiuchi-d/gozxing) - QR code reader
- [github.com/mdp/qrterminal/v3](https://github.com/mdp/qrterminal) - Terminal QR code display
- [github.com/skip2/go-qrcode](https://github.com/skip2/go-qrcode) - QR code generator

## License

This project is licensed under the Apache License 2.0 - see the 

LICENSE

 file for details.
```