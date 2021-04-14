package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func BytesToBinString(bb []byte) string {
	str := ""
	for _, b := range bb {
		str += fmt.Sprintf("%08b", b)
	}

	return str
}

func GetPublicIp() string {
	res, err := http.Get("http://ip-api.com/line")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(body), "\n")
	return split[len(split)-2]
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
