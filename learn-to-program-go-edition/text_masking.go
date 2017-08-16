// reference: ttps://stackoverflow.com/questions/33139020/can-golang-multiply-strings-like-python-can

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	x := "my new text is this long"
	y := strings.Repeat("#", utf8.RuneCountInString(x))
	fmt.Println(y)
}
