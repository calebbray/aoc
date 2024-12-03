package util

import "os"

func GetFileContent(source string) string {
	b, err := os.ReadFile(source)
	if err != nil {
		panic("an error occurred reading the file")
	}
	return string(b[:len(b)-1])
}
