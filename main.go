package main

import (
	"flag"
	"io"
	"os"
)

//	Checks for input and output flags:
//		- If output is set:
//			- Create a file as output location with ".tmp" suffix added
//			- When process is finished:
//				- If there are no errors, move temporary file to the output location
//				- Else, remove temporary file
//		- Else, use stdout
//
//		- If input is set, use processLoc func
//		- Else, pass process func Stdin
//
//		- After processing - if err exists, write to Stderr
func main() {
	var (
		inputLoc  string
		outputLoc string
		tmpLoc    string

		f      *os.File
		output io.Writer

		err error
	)

	flag.StringVar(&inputLoc, "i", "", "Location of input file to process (Stdin by default)")
	flag.StringVar(&outputLoc, "o", "", "Location of output file (Stdout by default)")
	flag.Parse()

	if len(outputLoc) == 0 {
		output = os.Stdout
	} else {
		tmpLoc = outputLoc + ".tmp"
		if f, err = os.Create(tmpLoc); err != nil {
			stderr("Error when trying to ", err)
			return
		}
		defer f.Close()
		output = f
	}

	switch len(inputLoc) {
	case 0:
		err = process(os.Stdin, output)
	default:
		err = processLoc(inputLoc, output)
	}

	if f != nil {
		f.Close()

		if err == nil {
			os.Rename(tmpLoc, outputLoc)
		} else {
			os.Remove(tmpLoc)
		}
	}

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
