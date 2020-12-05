package pix4d_2_cc

type Images struct {
	Id              string
	Image_name      string
	Path            string
	Image_full_path string
}

type Control_point struct {
	Control_point_name string
	X                  string
	Y                  string
	Z                  string
}

type Pix4d struct {
	Control_point_name string
	Image_x            string
	Image_y            string
	Image_full_path    string
}
