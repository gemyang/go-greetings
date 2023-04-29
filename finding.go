package greetings

import "errors"

func FindUniq(arr []float32) (float32, error) {
	if arr == nil || len(arr) < 3 {
		return -1, errors.New("bad input")
	}
	var a1, a2, a3 = arr[0], arr[1], arr[2]
	if a1 != a2 && a2 == a3 {
		return a1, nil
	}
	for _, v := range arr[1:] {
		if v != a1 {
			return v, nil
		}
	}
	return -1, errors.New("bad input")
}
