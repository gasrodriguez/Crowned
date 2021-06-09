package util

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"unicode/utf16"
	"unicode/utf8"
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
	// ref: https://stackoverflow.com/a/22467409/12742920
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func DecodeUTF16(b []byte) []byte {
	// ref: https://gist.github.com/bradleypeabody/185b1d7ed6c0c2ab6cec
	if len(b)%2 != 0 {
		panic("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.Bytes()
}
