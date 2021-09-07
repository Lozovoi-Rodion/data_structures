package array

import (
	"errors"
)

type Array []int

func NewArray(args... int) Array{
	var arr []int
	arr = append(arr, args...)
	return arr
}

func (a *Array) Append(el int) {
	*a = append(*a, el)
}

func (a *Array) Delete(idx int) (bool, error) {
	if idx >= len(*a) {
		return false, errors.New("index exceeds array size")
	}

	*a = append((*a)[:idx], (*a)[idx+1:]...)
	return true, nil
}

func (a *Array) Insert(idx, el int) (bool, error) {
	if idx >= len(*a) {
		return false, errors.New("out of boundary")
	}

	*a = append((*a)[:idx+1], (*a)[idx:]...)
	(*a)[idx] = el

	return true, nil
}

