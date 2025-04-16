package main

import "errors"

func onlyEven(arr []int) ([]int, error) {
	newArr := []int{}
	for _, v := range arr {
		if v%2 == 0 {
			newArr = append(newArr, v)
		} else {
			continue
		}
	}
	if len(newArr) == 0 {
		return nil, errors.New("NO_INT")
	}
	return newArr, nil
}

func main() {
}
