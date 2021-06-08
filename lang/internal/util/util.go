package util

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func LineCounter(filename string) (count uint32, err error) {
	//ref: https://stackoverflow.com/a/24563853/12742920
	r, err := os.Open(filename)
	buf := make([]byte, 32*1024)
	count = 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += uint32(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func SplitLines(s []byte) (lines []string) {
	// ref: https://stackoverflow.com/a/61938973/12742920
	sc := bufio.NewScanner(bytes.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func Exists(name string) bool {
	// https://stackoverflow.com/a/22467409/12742920
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
