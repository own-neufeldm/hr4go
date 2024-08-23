package main

import (
	"fmt"
	"math"
	"strings"
)

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GetHorizontalRule(title string, length int, border string, filler string) string {
	var reversedBorder = ReverseString(border)
	if title == "" {
		fillerLength := max(length-2-2*len(border), 1)
		return fmt.Sprintf("%s %s %s", border, strings.Repeat(filler, fillerLength), reversedBorder)
	}
	fillerLength := max(length-4-len(title)-2*len(border), 2)
	sideFillerLength := math.Ceil(float64(fillerLength) / 2)
	sideFiller := strings.Repeat(filler, int(sideFillerLength))
	return fmt.Sprintf("%s %s %s %s %s", border, sideFiller, title, sideFiller, reversedBorder)
}
