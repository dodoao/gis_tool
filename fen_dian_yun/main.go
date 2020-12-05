package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, err := ReadAll("xyz.txt");
	dir, _ := os.Getwd()
	if err == nil {
		var arr_tmp = strings.Split(string(s), "\n")
		line := 15000
		str_tmp := ""
		fmt.Println(len(arr_tmp))

		for i := 0; i < len(arr_tmp); i++ {

			if (strings.Index(arr_tmp[i], "-99999") == -1 && arr_tmp[i]!="" && strings.Index(arr_tmp[i]," ")==-1) {
				str_tmp += strconv.Itoa(i+1) + ",," + arr_tmp[i] + "\n"
			}

			if (i+1)%line == 0 {
				WriteFile(dir+"\\xyz_"+strconv.Itoa(i)+".dat", str_tmp)
				str_tmp = ""
			}
		}

		WriteFile(dir+"\\xyz_last.dat", str_tmp)
		str_tmp = ""

	}
}

func ReadAll(filePth string) (string, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return "", err
	}
	b, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		return "", err2
	}
	str_b := string(b)
	str_b = strings.ReplaceAll(str_b, "\r", "")
	str_b = strings.ReplaceAll(str_b, "\\", "/")
	return str_b, nil
}

func WriteFile(filename string, data string) {
	os.Remove(filename)
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(data))
		//fmt.Println(err)
	}
}
