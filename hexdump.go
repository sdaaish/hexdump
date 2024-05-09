package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	fName  string
	fUsage = "Enter the filename to process"
	nWidth int
	wUsage = "[Optional:] Characters to display"
	row    = 0
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
	fmt.Printf("%-5s", "Row")
	for i := 0; i < width; i++ {
		fmt.Printf("%04X ", i)
	}
	fmt.Println()

	for i := 0; i < width+1; i++ {
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

		fmt.Printf("%04d ", row)
		row += 1

		for i, b := range bytes {
			fmt.Printf("0x%02X ", b)

			if b < 33 || b > 126 {
				s = "."
			} else {
				s = string(b)
			}
			sText[i] = s

		}
		fmt.Printf("%s", sText)
		fmt.Println()
	}
}
