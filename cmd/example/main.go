package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/Snare295/architecture-lab2"
)

var (
	fileOutp  = flag.String("o", "", "Write outpun in file")
	fileInp   = flag.String("f", "", "Get input from file")
	stringInp = flag.String("e", "", "Compute this string")
)

func main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	if *stringInp != "" {
		input = strings.NewReader(*stringInp)
	}

	if *fileInp != "" {
		f, err := os.Open(*fileInp)
		if err != nil {
			os.Stderr.WriteString("Error : Opening file!")
			return
		}

		input = f
		defer f.Close()
	}

	if *fileOutp != "" {
		o, err := os.Create(*fileOutp)
		if err != nil {
			os.Stderr.WriteString("Error : Saving!")
			return
		}

		output = o
		defer o.Close()
	}

	if input == nil {
		os.Stderr.WriteString("Error : Empty input!")
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
