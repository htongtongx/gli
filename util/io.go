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
func ReadFile(path string) (str string, err error) {
	//打开文件的路径
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败")
		fmt.Println(err)
	}
	defer f.Close()
	//读取文件的内容
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("读取文件失败")
		fmt.Println(err)
		return "", err
	}
	str = string(fd)
	return str, nil
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
