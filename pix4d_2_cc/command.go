package pix4d_2_cc

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

func Read_images_file(file_name string) {
	s, err := ReadAll(file_name)
	if err == nil {
		var arr_tmp = strings.Split(string(s), "\n")
		for i := 0; i < len(arr_tmp); i++ {
			var arr_timp2 = strings.Split(arr_tmp[i], "\t")
			nm := strings.LastIndex(arr_timp2[1], "/")
			images_data_arr[strconv.Itoa(i)] = Images{
				Id:              arr_timp2[0],
				Image_name:      arr_timp2[1][nm:],
				Path:            arr_timp2[1][:nm],
				Image_full_path: arr_timp2[1],
			}
		}
	}
}

func Read_control_point(file_name string) {
	s, err := ReadAll(file_name)
	if err == nil {
		s = strings.ReplaceAll(s, "\r", "")
		var arr_tmp = strings.Split(string(s), "\n")
		for i := 0; i < len(arr_tmp); i++ {
			var arr_timp2 = strings.Split(arr_tmp[i], "\t")
			data_2000_arr[strconv.Itoa(i)] = Control_point{
				Control_point_name: arr_timp2[0],
				X:                  arr_timp2[1],
				Y:                  arr_timp2[2],
				Z:                  arr_timp2[3],
			}
		}
	}
}

func Read_pix4d(file_name string) {
	s, err := ReadAll(file_name)
	if err == nil {
		var arr_tmp = strings.Split(string(s), "\n")
		for i := 0; i < len(arr_tmp); i++ {
			var arr_timp2 = strings.Split(arr_tmp[i], ",")
			if (len(arr_timp2) != 5) {
				continue
			}
			pix4d_data_arr[strconv.Itoa(i)] = Pix4d{
				Control_point_name: arr_timp2[1],
				Image_x:            arr_timp2[2],
				Image_y:            arr_timp2[3],
				Image_full_path:    arr_timp2[4] + "/" + arr_timp2[0],
			}
		}
	}
}

func Get_images_data_arr_from_image_name(name string) []Images {
	var retu_arr []Images
	for i := 0; i < len(images_data_arr); i++ {
		str_i := strconv.Itoa(i)
		if strings.ToUpper(images_data_arr[str_i].Image_full_path) == strings.ToUpper(name) {
			retu_arr = append(retu_arr, images_data_arr[str_i])
		}
	}
	return retu_arr
}

func Get_pix4d_data_arr_from_name(name string) []Pix4d {
	var retu_arr []Pix4d
	for i := 0; i < len(pix4d_data_arr); i++ {
		str_i := strconv.Itoa(i)
		if strings.ToUpper(pix4d_data_arr[str_i].Control_point_name) == strings.ToUpper(name) {
			retu_arr = append(retu_arr, pix4d_data_arr[str_i])
		}
	}
	return retu_arr
}

func Output_cc_control_point(filename string) {
	block_surveys, err := ReadAll(dir + "\\Block_1 - Surveys.txt")
	find_str := "<CheckPoint>false</CheckPoint>"
	find_str_len := len(find_str)
	var new_data = ""
	if (err == nil) {

		last_nm := 0

		for {
			nm := strings.Index(block_surveys[last_nm:], find_str)
			if nm == -1 {
				break
			}
			nm += find_str_len
			new_data += block_surveys[last_nm : last_nm+nm]
			last_nm += nm

			tmp_name_start := strings.LastIndex(new_data, "<Name>")
			tmp_name_end := strings.LastIndex(new_data, "</Name>")
			control_point_name := new_data[tmp_name_start+6 : tmp_name_end]

			var arr_tmp = Get_pix4d_data_arr_from_name(control_point_name)
			if len(arr_tmp) != 0 {
				for j := 0; j < len(arr_tmp); j++ {
					tmp_image := Get_images_data_arr_from_image_name(arr_tmp[j].Image_full_path)

					if len(tmp_image) > 0 {
						new_data += "\t\t\t<Measurement>\r\n" +
							"\t\t\t\t<PhotoId>" + tmp_image[0].Id + "</PhotoId>\r\n" +
							"\t\t\t\t<ImagePath>" + arr_tmp[j].Image_full_path + "</ImagePath>\r\n" +
							"\t\t\t\t<x>" + arr_tmp[j].Image_x + "</x>\r\n" +
							"\t\t\t\t<y>" + arr_tmp[j].Image_y + "</y>\r\n" +
							"\t\t\t</Measurement>\r\n"
					}
				}
			}

		}
	}

	new_data += "</ControlPoint>\n\t</ControlPoints>\n" + "</SurveysData>"
	WriteFile(filename, new_data)
}
func WriteFile(filename string, data string) {
	os.Remove(filename)
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(data))
		fmt.Println(err)
	}
}

func Read_cc_export_xml(filename string) {
	data := ""
	tmp_data, err := ReadAll(filename)
	if (err == nil) {
		image_start := 0
		image_end := 0

		for image_start != -1 {
			end := image_end
			if (strings.Index(tmp_data[end:], "<ImagePath>") == -1) {
				break
			}
			image_start = len(tmp_data[:end]) + strings.Index(tmp_data[end:], "<ImagePath>")
			image_end = image_start + strings.Index(tmp_data[image_start:], "</ImagePath>")

			image_path := tmp_data[image_start+11 : image_end]

			tmp_data2 := tmp_data[:image_start]

			id_start := strings.LastIndex(tmp_data2, "<Id>")
			id_end := strings.LastIndex(tmp_data2, "</Id>")
			id := tmp_data2[id_start+4 : id_end]

			data += id + "\t" + image_path + "\n"

		}
		if (data != "") {
			WriteFile(dir+"\\images.txt", data[:len(data)-1])
		}

	}

}
