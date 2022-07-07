package main

import (
	"fmt"
	"math"
	"math/rand"
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
	testing_provider_connection()
	testing_learning()

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
