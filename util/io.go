package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//解析text文件内容
func ReadFile(path string) (content string, err error) {
	//打开文件的路径
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	content = string(contentByte)
	return
}

func ReadCsvFile(path string) (result [][]string) {
	fs, err := os.Open(path)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	result = [][]string{}
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
		result = append(result, row)
	}
	return result
}
