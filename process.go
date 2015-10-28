package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strconv"
)

func processLoc(loc string, w io.Writer) (err error) {
	var f *os.File
	if f, err = os.Open(loc); err != nil {
		stderr("Error when trying to ", err)
		return
	}
	defer f.Close()

	return process(f, w)
}

func process(r io.Reader, w io.Writer) (err error) {
	var (
		i int
		d []byte

		bd  breakdown     = breakdown{}
		buf *bufio.Reader = bufio.NewReader(r)
	)

	// Move past line of column names
	buf.ReadLine()
	d, _, err = buf.ReadLine()

	for err == nil {
		var (
			l    *LogLine
			lerr error

			b []byte
		)

		if l, lerr = newLogLine(d); lerr != nil {
			lineParseErr("Error processing line", lerr, i)
			continue
		}

		if b, lerr = json.Marshal(l); lerr != nil {
			lineParseErr("Error marshaling line", lerr, i)
			continue
		}

		str := strconv.FormatInt(l.CreatedTimestamp.Unix(), 10) + "@" + string(b) + "\n"
		bd.Set("all", l.CreatedTimestamp, str)

		if l.TransactionStatus == "settled" {
			bd.Set("success", l.CreatedTimestamp, str)
		} else {
			bd.Set("fail", l.CreatedTimestamp, str)
		}

		d, _, err = buf.ReadLine()
		i++
	}

	if err != io.EOF {
		stderr("Error processing input: ", err)
		return err
	}

	return bd.exportTar(w)
}
