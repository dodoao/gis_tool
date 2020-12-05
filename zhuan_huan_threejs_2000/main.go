package main

import (
	"fmt"
)

type Eps_fl struct {
	X float64
	Y float64
	Z float64
}

func fl_2_s_int(epsFl Eps_fl) Eps_fl {

	epsInt := Eps_fl{}
	epsInt.Z = float64(int64(epsFl.Z * Bei))
	epsInt.Y = float64(int64(epsFl.Y * Bei))
	epsInt.X = float64(int64(epsFl.X * Bei))

	return epsInt
}

var Bei float64 = 100000000

var a_2000 Eps_fl
var b_2000 Eps_fl
var a_threejs Eps_fl
var b_threejs Eps_fl

var a_b_2000_x_long float64
var a_b_2000_y_long float64
var a_b_2000_z_long float64

var a_b_threejs_x_long float64
var a_b_threejs_y_long float64
var a_b_threejs_z_long float64

var a_b_2000_threejs_x float64
var a_b_2000_threejs_y float64
var a_b_2000_threejs_z float64
var a_b_threejs_2000_x float64
var a_b_threejs_2000_y float64
var a_b_threejs_2000_z float64

func main() {
	a_2000 = fl_2_s_int(Eps_fl{
		78830.516,
		2697143.420,
		205.422,
	})
	b_2000 = fl_2_s_int(Eps_fl{
		79983.080,
		2697326.686,
		101.86,
	})
	a_threejs = fl_2_s_int(Eps_fl{
		-1436.573468627886,
		570.5277243139094,
		84.03069621363522,
	})
	b_threejs = fl_2_s_int(Eps_fl{
		-285.71485038294196,
		752.6214550612525,
		-20.283219026296685,
	})

	a_b_2000_x_long = a_2000.X - b_2000.X
	a_b_2000_y_long = a_2000.Y - b_2000.Y
	a_b_2000_z_long = a_2000.Z - b_2000.Z
	a_b_threejs_x_long = a_threejs.X - b_threejs.X
	a_b_threejs_y_long = a_threejs.Y - b_threejs.Y
	a_b_threejs_z_long = a_threejs.Z - b_threejs.Z
	a_b_2000_threejs_x = a_b_2000_x_long / a_b_threejs_x_long
	a_b_2000_threejs_y = a_b_2000_y_long / a_b_threejs_y_long
	a_b_2000_threejs_z = a_b_2000_z_long / a_b_threejs_z_long
	a_b_threejs_2000_x = a_b_threejs_x_long / a_b_2000_x_long
	a_b_threejs_2000_y = a_b_threejs_y_long / a_b_2000_y_long
	a_b_threejs_2000_z = a_b_threejs_z_long / a_b_2000_z_long

	////////////////////////////////////////
	threejs_to_2000(b_threejs)
	_2000_to_threejs(b_2000)
}

func threejs_to_2000(c_threejs Eps_fl) {
	a_c_threejs_x_long := c_threejs.X - a_threejs.X
	a_c_threejs_y_long := c_threejs.Y - a_threejs.Y
	a_c_threejs_z_long := c_threejs.Z - a_threejs.Z
	a_c_2000_x_long := a_b_2000_threejs_x * a_c_threejs_x_long
	a_c_2000_y_long := a_b_2000_threejs_y * a_c_threejs_y_long
	a_c_2000_z_long := a_b_2000_threejs_z * a_c_threejs_z_long
	a_c_2000_x := a_2000.X + a_c_2000_x_long
	a_c_2000_y := a_2000.Y + a_c_2000_y_long
	a_c_2000_z := a_2000.Z + a_c_2000_z_long
	x := a_c_2000_x / Bei
	y := a_c_2000_y / Bei
	z := a_c_2000_z / Bei
	fmt.Println(Decimal(x, "4"), Decimal(y, "4"), Decimal(z, "4"))
}

func _2000_to_threejs(c_2000 Eps_fl) {

	a_c_2000_x_long := c_2000.X - a_2000.X
	a_c_2000_y_long := c_2000.Y - a_2000.Y
	a_c_2000_z_long := c_2000.Z - a_2000.Z

	a_c_threejs_x_long := a_b_threejs_2000_x * a_c_2000_x_long
	a_c_threejs_y_long := a_b_threejs_2000_y * a_c_2000_y_long
	a_c_threejs_z_long := a_b_threejs_2000_z * a_c_2000_z_long
	a_c_threejs_x := a_threejs.X + a_c_threejs_x_long
	a_c_threejs_y := a_threejs.Y + a_c_threejs_y_long
	a_c_threejs_z := a_threejs.Z + a_c_threejs_z_long
	x := a_c_threejs_x / Bei
	y := a_c_threejs_y / Bei
	z := a_c_threejs_z / Bei
	fmt.Println(Decimal(x, "9"), Decimal(y, "9"), Decimal(z, "9"))
}

func Decimal(value float64, s string) string {
	return fmt.Sprintf("%."+s+"f", value)
}
