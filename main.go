package main

import (
	"fmt"
)

func main() {
	var test product
	test.name = "Lorazepam 2 mg/mL Injection 1 mL Carpuject"
	test.medID = 0
	inputNDC(&test, "00409-1985-30")
	test.isOutterPackage = true
	test.packageSize = 10
	NDCstring := test.getNDC()
	fmt.Println(NDCstring)
}
