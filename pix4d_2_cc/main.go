package main

import (
	"os"
)

var images_data_arr map[string]Images
var data_2000_arr map[string]Control_point
var pix4d_data_arr map[string]Pix4d
var dir string

func main() {
	images_data_arr = map[string]Images{}
	data_2000_arr = map[string]Control_point{}
	pix4d_data_arr = map[string]Pix4d{}
	//获取当前路径
	dir, _ = os.Getwd()
	Read_cc_export_xml(dir+"\\cc_export.txt")
	Read_images_file(dir + "\\images.txt")
	Read_control_point(dir + "\\control_point.txt")
	Read_pix4d(dir + "\\pix4d.txt")
	Output_cc_control_point(dir + "/cc.xml")
}
