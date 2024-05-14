package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	fName  string
	fUsage = "Enter the filename to process"
	nWidth int
	wUsage = "[Optional:] Characters to display"
	offset = 0
	s      string
)

func init() {
	flag.StringVar(&fName, "file", "", fUsage)
	flag.StringVar(&fName, "f", "", fUsage+", short form")

	flag.IntVar(&nWidth, "width", 16, wUsage)
	flag.IntVar(&nWidth, "w", 16, wUsage+", short form.")
}

func printHeader(width int) {

	// Print a header
	fmt.Printf("%-7s", "Offset")
	for i := 0; i < width; i++ {
		fmt.Printf("%04X ", i)
	}
	fmt.Println()

	fmt.Printf("%s ", "------")
	for i := 0; i < width; i++ {
		fmt.Printf("%4s ", "----")
	}
	fmt.Println()
}

func main() {

	flag.Parse()

	if fName == "" {
		flag.Usage()
		fmt.Println("\nerror: Enter a valid filename.")
		os.Exit(1)
	}

	if nWidth < 1 || nWidth > 64 {
		flag.Usage()
		fmt.Println("\nerror: Enter a width between 1 and 64")
		os.Exit(1)
	}

	infile, err := os.Open(fName)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}
	defer infile.Close()

	printHeader(nWidth)

	buf := make([]byte, nWidth)
	sText := make([]string, nWidth)

	for {
		c, err := io.ReadFull(infile, buf)

		bytes := []byte(buf[:c])

		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil && err != io.ErrUnexpectedEOF {
			log.Fatal(err)
		}

		xText := make([]string, 0, nWidth)
		var x string
		for i, b := range bytes {
			x = fmt.Sprintf("0x%02X", b)
			xText = append(xText, x)

			if b < 32 || b > 126 {
				s = "."
			} else {
				s = string(b)
			}
			sText[i] = s
		}

		// Adjust buffer length to width boundary
		if len(xText) < nWidth {
			n := nWidth - len(xText)
			for i := 0; i < n; i++ {
				xText = append(xText, "    ")
			}
		}

		hexT := strings.Trim(fmt.Sprintf("%s", xText), "[]")
		stringT := strings.Trim(fmt.Sprintf("%s", sText), "[]")
		fmt.Printf("%06x %s |%s|", offset, hexT, stringT)
		fmt.Println()
		offset += 1 * nWidth
	}
}
