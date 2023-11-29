package props

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Convert2Unicode 字符串转成 unicode
func Convert2Unicode(str string) error {
	fmt.Printf("========== \tConvert2Unicode\t ==========\n")
	split := strings.Split(str, "")
	for _, item := range split {
		inString, _ := utf8.DecodeRuneInString(item)
		ascii := strconv.QuoteRuneToASCII(inString)
		fmt.Printf(strings.ReplaceAll(ascii, "'", ""))
	}
	fmt.Println()
	return nil
}

// Convert2Str unicode 转成 字符串
func Convert2Str(str string) error {
	fmt.Printf("========== \tConvert2Str\t ==========\n")
	unquote, err := strconv.Unquote(`"` + str + `"`)
	if err != nil {
		return err
	}
	fmt.Println(unquote)
	return nil
}
