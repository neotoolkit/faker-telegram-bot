package main

import (
	"strconv"
	"strings"

	"github.com/neotoolkit/faker"
)

func number(f *faker.Faker, text string) string {
	args := strings.Split(text, " ")
	min := 0
	max := 100
	if len(args) >= 3 {
		convMin, err := strconv.Atoi(args[1])
		if nil == err {
			min = convMin
		}
		convMax, err := strconv.Atoi(args[2])
		if nil == err {
			max = convMax
		}
	}

	return strconv.Itoa(f.Number(min, max))
}
