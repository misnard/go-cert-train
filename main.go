package main

import (
	"fmt"
	"go-cert-train/cert"
	"go-cert-train/pdf"
	"os"
)

func main() {
	c, err := cert.New("Golang programming", "Bob Dylan", "2021-06-21")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err = pdf.New("output")
	if err != nil {
		fmt.Printf("Error during pdf generator creation: %v", err)
		os.Exit(1)
	}

	err = saver.Save(*c)
	if err != nil {
		fmt.Printf("Error during pdf saving: %v", err)
		os.Exit(1)
	}
}
