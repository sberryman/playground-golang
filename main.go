package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"

	// "Testing_learning"
	"github.com/sberryman/playground-golang/api/cellular/hologram"
	"github.com/sberryman/playground-golang/api/cellular/verizon"
	"github.com/sberryman/playground-golang/api/cellular/vodafone"
)

var c, python, java bool
var i int
var j, k int = 1, 2

func main() {
	// testing_provider_connection()
	// testing_learning()
	// loop()
	// loop2()
	// fmt.Println(sqrt(2), sqrt(-4))
	// fmt.Println(pow(3, 3, 1))
	// fmt.Println(pow2(2, 5, 10))
	// fmt.Println(Sqrt2(25))
	// switchtest()
	date()
}

func testing_learning() {
	fmt.Println("My favorite number is", rand.Intn(1000))
	fmt.Println(time.Now())
	fmt.Printf("Now you have %v problems.\n", math.Sqrt(7))
	// fmt.Println(v)
	fmt.Println(math.Pi)
	fmt.Println(add(2, 1))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	// var i int
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	v := 42
	fmt.Printf("v is of type %T\n", v)
}
func add(x int, y int) int {
	return x + y
}
func swap(x string, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func testing_provider_connection() {
	// hologram
	fmt.Println("Hologram: fetching iccid: 8935711001077014154")
	holoDevice, err := hologram.ListDeviceByIccid("8935711001077014154")
	if err != nil {
		fmt.Println("Hologram: error:", err)
	} else {
		fmt.Println("Hologram: device:", holoDevice)
	}
	fmt.Println("")
	fmt.Println("")

	// vodafone
	fmt.Println("Vodafone: fetching iccid: 89314404000489735796")
	vodaDevice, err := vodafone.ListDeviceByIccid("89314404000489735796")
	if err != nil {
		fmt.Println("Vodafone: error:", err)
	} else {
		fmt.Println("Vodafone: device:", vodaDevice)
	}
	fmt.Println("")
	fmt.Println("")

	// verizon
	fmt.Println("Verizon: fetching iccid: 89148000004573400960")
	verizonDevice, err := verizon.RetrieveDeviceInformation("89148000004573400960")
	if err != nil {
		fmt.Println("Verizon: error:", err)
	} else {
		fmt.Println("Verizon: device:", verizonDevice)
	}
}
func loop() {
	sum := 0
	for i := 0; i < 4; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func loop2() {
	sum := 1
	loop := 1
	for sum < 1000000000 { //there is normally a semicolon before and afer the sum<1000 because you need to define the init and post but these are optional
		sum = sum + sum
		fmt.Println("loop nº: ", loop, " value is :", sum)
		loop = loop + 1
	}
	// fmt.Println("total: ", sum, " in ", loop-1, " loops")
	fmt.Printf("loop n: %d value is: %d\n", loop, sum)
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%f >= %g \n", v, lim)
	}
	return lim
}
func Sqrt2(x float64) float64 {
	z := float64(1)
	if z -= ((z*z - x) / (2 * z)); z < x {
		// Println(z)
		return z
	} else {
		fmt.Printf("nope: ")
	}
	return z
}
func switchtest() {
	fmt.Println("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
func date() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	default:
		fmt.Println("Not today")
	}
}
