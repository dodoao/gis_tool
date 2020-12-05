package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	ListDir(dir, ".pgw")

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
	f.Close()
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

func ListDir(dir, suffix string) (files []string, err error) {
	files = []string{}

	_dir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToLower(suffix) //匹配后缀

	for _, _file := range _dir {
		if _file.IsDir() {
			continue //忽略目录
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(_file.Name()), suffix) {
			//文件后缀匹配
			//files = append(files, path.Join(dir, _file.Name()))
			name := _file.Name()[0:strings.Index(_file.Name(), suffix)]

			os.Rename(dir+"\\"+_file.Name(), dir+"\\"+name+".txt")

			txt, err := ReadAll(dir + "\\" + name + ".txt")

			if err != nil {
				fmt.Println(err)
				continue
			}

			txt_arr := strings.Split(txt, "\n")

			new_name := txt_arr[4] + "-" + txt_arr[5]

			oldname := dir + "\\" + name + ".txt"
			newname := dir + "\\" + new_name + ".pgw"
			err = os.Rename(oldname, newname)
			if err != nil {
				fmt.Println(err)
			}
			oldname = dir + "\\" + name + ".png"
			newname = dir + "\\" + new_name + ".png"
			os.Rename(oldname, newname)
			oldname = dir + "\\" + name + ".prj"
			newname = dir + "\\" + new_name + ".prj"
			os.Rename(oldname, newname)

		}
	}

	return files, nil
}
