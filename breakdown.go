package main

import (
	"archive/tar"
	"bytes"
	"io"
	"sort"
	"time"
)

const (
	mnFmt = `01`
	yrFmt = `2006`
)

type breakdown struct {
	All     breakdownMap
	Success breakdownMap
	Fail    breakdownMap
}

func (b *breakdown) exportTar(w io.Writer) (err error) {
	t := tar.NewWriter(w)
	defer t.Close()

	if b.All != nil {
		b.All.process(t, "all")
		if err = b.All.process(t, "all"); err != nil {
			return err
		}
	}

	if b.Success != nil {
		if err = b.Success.process(t, "success"); err != nil {
			return err
		}
	}

	if b.Fail != nil {
		if err = b.Fail.process(t, "fail"); err != nil {
			return err
		}
	}

	return nil
}

func (b *breakdown) Set(k string, ts time.Time, str string) {
	var bdm breakdownMap

	yv := ts.Format(yrFmt)
	mv := ts.Format(mnFmt)

	switch k {
	case "all":
		if bdm = b.All; bdm == nil {
			b.All = make(breakdownMap)
			bdm = b.All
		}
	case "success":
		if bdm = b.Success; bdm == nil {
			b.Success = make(breakdownMap)
			bdm = b.Success
		}
	case "fail":
		if bdm = b.Fail; bdm == nil {
			b.Fail = make(breakdownMap)
			bdm = b.Fail
		}
	default:
		return
	}

	if _, ok := bdm[yv]; !ok {
		bdm[yv] = make(map[string][]string)
	}

	if _, ok := bdm[yv][mv]; !ok {
		bdm[yv][mv] = make([]string, 0)
	}

	bdm[yv][mv] = append(bdm[yv][mv], str)
}

//				   Year		  Month	   Item
type breakdownMap map[string]map[string][]string

func (b breakdownMap) process(t *tar.Writer, ttl string) error {
	now := time.Now()

	for yk, y := range b {
		for mk, mn := range y {
			buf := bytes.NewBuffer(nil)
			sort.Strings(mn)

			for _, item := range mn {
				buf.WriteString(item)
			}

			if err := t.WriteHeader(&tar.Header{
				Name:    ttl + "/" + yk + "/" + mk + ".log",
				Mode:    0600,
				Size:    int64(buf.Len()),
				ModTime: now,
			}); err != nil {
				return err
			}

			io.Copy(t, buf)
		}
	}

	return nil
}
