package main

import "fmt"

func buildPiramid(start, delt, dim int) {
	piramidMatrix := make([][]int, dim)

	for i := range piramidMatrix {
		piramidMatrix[i] = make([]int, dim)
	}

	mid := dim / 2

	//higler half horizontal cicle
	prev := 0
	dlt := start

	for i := 0; i <= mid; i++ {
		border := dim - 1 - prev
		for j := i; j <= border; j++ {
			piramidMatrix[i][j] = dlt
		}
		prev += 1
		dlt += delt
	}

	//higler half vertical cicle
	prev = 0
	dlt = start

	for j := 0; j <= mid; j++ {
		border := dim - 1 - prev
		for i := j; i <= border; i++ {
			piramidMatrix[i][j] = dlt
		}
		prev += 1
		dlt += delt
	}

	//lower half horizontal
	border := 0
	dlt = start

	for i := dim - 1; i > mid; i-- {
		for j := i; j > border; j-- {
			piramidMatrix[i][j] = dlt
		}
		dlt += delt
		border++
	}

	//lower half vertical
	border = 0
	dlt = start

	for j := dim - 1; j > mid; j-- {
		for i := j; i > border; i-- {
			piramidMatrix[i][j] = dlt
		}
		dlt += delt
		border++
	}

	for i := range piramidMatrix {
		fmt.Println(piramidMatrix[i])
	}

}

func main() {
	var (
		start int = 1
		delt  int = 1
		dim   int = 5
	)

	buildPiramid(start, delt, dim)
}
