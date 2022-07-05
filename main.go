package main

import (
	"fmt"

	"github.com/sberryman/playground-golang/api/cellular/hologram"
	"github.com/sberryman/playground-golang/api/cellular/verizon"
	"github.com/sberryman/playground-golang/api/cellular/vodafone"
)

func main() {

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
