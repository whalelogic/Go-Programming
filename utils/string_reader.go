package utils 

import (
	"strings"
	"io"
)

func StringReader(s string) io.Reader {
	return strings.NewReader(s)
}
