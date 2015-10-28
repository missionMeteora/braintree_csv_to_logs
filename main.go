package main

import (
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
		err       error
		f         *os.File
		output    io.Writer
		i, o, tmp string = getFlagLocs()
	)

	if o == "" {
		output = os.Stdout
	} else {
		if f, err = os.Create(tmp); err != nil {
			stderr("Error when trying to ", err)
			return
		}

		output = f
	}

	if i == "" {
		err = process(os.Stdin, output)
	} else {
		err = processLoc(i, output)
	}

	closeFile(f, err, o, tmp)
	reportErrors(err)
}
