package models

import (
	"io/ioutil"
	"strings"
)

func SeedData() {

	file, err := ioutil.ReadAll("./test-setup.sql")

	if err != nil {

	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		result, err := GetDB().Table("locations").Insert(request)
	}
}
