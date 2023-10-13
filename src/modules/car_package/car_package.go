package carpackage

import (
	"time"
)

type Car struct {
	Name     string
	NickName string
	Year     int
}

type fullCarPrivate struct {
	FullName string
	Year     int
	YearOld  int
}

func (c Car) CreateFullCar() fullCarPrivate {
	yearNow := time.Now().Year()
	finalCar := fullCarPrivate{
		FullName: c.Name + " " + c.NickName,
		Year:     c.Year,
		YearOld:  yearNow - c.Year,
	}
	return finalCar
}
