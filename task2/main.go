package main

import (
	"fmt"
	conv "strconv"
)

func change(item int, index int) string {
	return conv.Itoa(item % 2)
}

func MapTo(list []int, change func(int, int) string) (result []string) {
	for index, value := range list {
		//value, _ := conv.Atoi(change(value, index))
		result = append(result, change(value, index))
	}
	return
}

// func ConvertItem(item string) string {
// 	switch item {
// 	case "0":
// 		return "zero"
// 	case "1":
// 		return "one"
// 	case "2":
// 		return "two"
// 	case "3":
// 		return "three"
// 	case "4":
// 		return "four"
// 	case "5":
// 		return "five"
// 	case "6":
// 		return "six"
// 	case "7":
// 		return "seven"
// 	case "8":
// 		return "eight"
// 	case "9":
// 		return "nine"
// 	default:
// 		return "unknown"
// 	}
// }

func ConvertItem(item int) string {
	switch item {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return "unknown"
	}
}

func Convert(arr []int) (result []string) {
	for _, value := range arr {
		result = append(result, ConvertItem(value))
	}
	return
}

func ConvertFromStringToInt(list []string) (result []int) {
	for _, value := range list {
		value, _ := conv.Atoi(value)
		result = append(result, value)
	}
	return
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Convert(ConvertFromStringToInt(MapTo(input, change))))
}
