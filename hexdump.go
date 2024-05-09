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
)

func init() {
	flag.StringVar(&fName, "file", "", fUsage)
	flag.StringVar(&fName, "f", "", fUsage+", short form")

	flag.IntVar(&nWidth, "width", 16, wUsage)
	flag.IntVar(&nWidth, "w", 16, wUsage+", short form.")
}

func main() {

	flag.Parse()

	if fName == "" {
		flag.Usage()
		fmt.Println("\nerror: Enter a valid filename.")
		os.Exit(1)
	}

	infile, err := os.Open(fName)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}
	defer infile.Close()

	chunkSize := nWidth

	// Print a header
	fmt.Printf("%-5s", "Row")
	for i := 0; i < chunkSize; i++ {
		fmt.Printf("%04X ", i)
	}
	fmt.Println()

	for i := 0; i < chunkSize+1; i++ {
		fmt.Printf("%4s ", "----")
	}
	fmt.Println()

	buf := make([]byte, chunkSize)

	row := 0

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

		for _, b := range bytes {
			fmt.Printf("0x%02X ", b)
		}
		fmt.Println()
	}
}
