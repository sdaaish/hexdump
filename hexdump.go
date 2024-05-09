package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const chunkSize = 16

func main() {

	fmt.Println("Reading...\n")

	if len(os.Args) < 2 {
		fmt.Println("error: Enter a valid filename.")
		os.Exit(1)
	}

	filename := os.Args[1]

	infile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}
	defer infile.Close()

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
								l, err := io.ReadFull(infile, buf)

								bytes := []byte(buf[:l])

								if err == io.EOF {
												os.Exit(0)
								}

								//		if err == io.ErrUnexpectedEOF {
								// fmt.Printf("L is: %d, ", l)
								// fmt.Printf("the buf is: %d\n", len(buf))
								//}

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
