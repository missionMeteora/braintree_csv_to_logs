package main

import (
	"flag"
	"io"
	"os"
	"strconv"
)

func lineParseErr(pfx string, err error, i int) {
	stderr(pfx+" on line #"+strconv.Itoa(i)+": ", err)
}

func stderr(pfx string, err error) {
	os.Stderr.WriteString(pfx + err.Error() + "\n")
}

func getFlagLocs() (i, o, t string) {
	flag.StringVar(&i, "i", "", "Location of input file to process (Stdin by default)")
	flag.StringVar(&o, "o", "", "Location of output file (Stdout by default)")
	flag.Parse()
	return i, o, i + ".tmp"
}

func getInput(i string) (input io.ReadCloser, err error) {
	if i == "" {
		return os.Stdin, nil
	}

	if input, err = os.Open(i); err != nil {
		return nil, err
	}

	return input, nil
}

func getOutput(o, tmp string) (output io.WriteCloser, err error) {
	if o == "" {
		return os.Stdout, nil
	}

	if output, err = os.Create(tmp); err != nil {
		return nil, err
	}

	return output, nil
}

func handleOutput(err error, o, tmp string) {
	if o == "" {
		return
	}

	if err == nil {
		os.Rename(tmp, o)
	} else {
		os.Remove(tmp)
	}
}

func reportErrors(err error) {
	if err != nil {
		stderr("", err)
	}
}
