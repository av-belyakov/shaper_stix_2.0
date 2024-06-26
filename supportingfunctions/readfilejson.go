package supportingfunctions

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func ReadFileJson(fpath, fname string) ([]byte, error) {
	var newResult []byte

	rootPath, err := GetRootPath("placeholder_misp")
	if err != nil {
		return newResult, err
	}

	tmp := strings.Split(rootPath, "/")

	fmt.Println("func 'readFileJson', path = ", path.Join(path.Join(tmp[:6]...), fpath, fname))

	f, err := os.OpenFile("/"+path.Join(path.Join(tmp[:6]...), fpath, fname), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return newResult, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		newResult = append(newResult, sc.Bytes()...)
	}

	return newResult, nil
}
