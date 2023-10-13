package main

import (
	"fmt"
	carpackage "start/modules/car_package"
	"start/modules/mathpackage"
)

func main() {
	intSum := mathpackage.SumInt(50, 21)
	fmt.Println("Result of SumInt: ", intSum)
	floatSum := mathpackage.SumFloat(3.14, 17.8)
	fmt.Println("Result of SumFloat: ", floatSum)

	fmt.Println("****")
	fmt.Println("Starting explore Car Module")
	car := carpackage.Car{
		Name:     "Corolla",
		NickName: "Giant car",
		Year:     2020,
	}
	fullCar := car.CreateFullCar()
	fmt.Println("Full Car name: ", fullCar.FullName)
	fmt.Println("Full Car Years Old: ", fullCar.YearOld)
}
