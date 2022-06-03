package util

import (
	"fmt"
	"io"
)

func CheckError(err error, w io.Writer) bool {
	if err != nil {
		writeWithWriter(w, "ERROR: "+err.Error()+"\n")
		return true
	}
	return false
}

func writeWithWriter(w io.Writer, message string) {
	if w == nil {
		fmt.Print(message)
	}
	fmt.Fprint(w, message)
}
