package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		insert := make([]uint8, dx)
		for j := 0; j < len(insert)/2; j++ {
			insert[j] = uint8(len(insert) - j)
		}
		for j := len(insert) - 1; j >= len(insert)/2; j-- {
			insert[j] = uint8(len(insert) - j)
		}
		ret[i] = insert
	}

	return ret
}

func main() {
	pic.Show(Pic)
}
