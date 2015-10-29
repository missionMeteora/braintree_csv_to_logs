package main

import (
	"io"
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
		input     io.ReadCloser
		output    io.WriteCloser
		i, o, tmp string = getFlagLocs()
	)

	if input, err = getInput(i); err != nil {
		stderr("", err)
	}

	if output, err = getOutput(o, tmp); err != nil {
		stderr("", err)
	}

	if err == nil {
		err = process(input, output)
	}

	input.Close()
	output.Close()

	handleOutput(err, o, tmp)
	reportErrors(err)
}
