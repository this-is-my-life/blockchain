package utils

import "fmt"

func BytesToBinString(bb []byte) string {
	str := ""
	for _, b := range bb {
		str += fmt.Sprintf("%08b", b)
	}

	return str
}

// from https://stackoverflow.com/a/39281081
func PadOrTrim(bb []byte, size int) []byte {
	l := len(bb)
	if l == size {
		return bb
	}
	if l > size {
		return bb[l-size:]
	}
	tmp := make([]byte, size)
	copy(tmp[size-l:], bb)
	return tmp
}
