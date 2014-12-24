package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func AddStrings(str1, str2 string, int1 int) string {
	return str1 + " " + str2 + " " + strconv.Itoa(int1)
}

func AddStringswJoin(str1, str2 string, int1 int) string {
	return strings.Join([]string{str1, str2, strconv.Itoa(int1)}, " ")
}

func AddStringswFormat(str1, str2 string, int1 int) string {
	return fmt.Sprintf("%s %s %v", str1, str2, int1)
}

func AddBytesBuffer(str1, str2 string, int1 int) string {

	var b bytes.Buffer
	b.Reset()
	b.WriteString(str1)
	b.WriteString(" ")
	b.WriteString(str2)
	b.WriteString(" ")
	b.WriteString(strconv.Itoa(int1))
	return b.String()

}
