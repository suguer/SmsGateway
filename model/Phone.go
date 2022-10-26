package model

import (
	"strconv"
	"strings"
)

type Phone struct {
	number  int
	idccode int
}

func (phone *Phone) GetNumber() string {
	return strconv.Itoa(phone.number)
}

func (phone *Phone) GetIdCCode() int {
	return phone.idccode
}

func (phone *Phone) IsChineseCode() bool {
	return phone.idccode == 86
}

func (phone *Phone) GetUniversalNumber() string {
	return "+" + strconv.Itoa(phone.idccode) + strconv.Itoa(phone.number)
}

func NewPhone(content string) *Phone {
	idccode := 0
	number := 0
	array := strings.Split(content, ".")
	if len(array) == 2 {
		str, _ := strconv.Atoi(array[1])
		number = str
		str, _ = strconv.Atoi(strings.Replace(array[0], "+", "", -1))
		idccode = str
	} else {
		str, _ := strconv.Atoi(array[0])
		number = str
		idccode = 86
	}
	return &Phone{
		number:  number,
		idccode: idccode,
	}
}
