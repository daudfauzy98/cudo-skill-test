package main

import "fmt"

func main() {
	var size, x, y int
	fmt.Print("Input size, x, and y (space separated): ")
	fmt.Scan(&size, &x, &y)
	// fmt.Printf("size: %d  x: %d  y: %d\n", size, x, y)

	fmt.Println("Required number:", getRequiredNumber(size, x, y))
}

func getRequiredNumber(size int, x int, y int) int {
	matrix := generateMatrixReversed(size)
	// printMatrix(matrix)
	element := getElementByCoordinate(matrix, x, y)
	return element
}

func getElementByCoordinate(mx [][]int, x, y int) int {
	convY := (len(mx) - 1) - (y - 1)
	convX := x - 1
	return mx[convY][convX]
}

func generateMatrixReversed(size int) [][]int {
	numberedEle := countNumberedElements(size)
	formatedNum := getFormatedNumber(numberedEle)

	numIdx := getFormatedNumber(size)
	myArr := make([][]int, size)
	for i := 0; i < size; i++ {
		// fmt.Print("i:", i)
		temp := numIdx[i]
		for j := 0; j < size; j++ {
			// fmt.Print(" j:", j)
			// fmt.Print(" temp:", temp)
			if j <= i {
				myArr[i] = append(myArr[i], formatedNum[len(formatedNum)-temp])
			} else {
				myArr[i] = append(myArr[i], 0)
			}
			temp--
		}
	}
	return myArr
}

func countNumberedElements(sz int) int {
	var num int
	// j := sz
	for i := 1; i <= sz; i++ {
		for j := 1; j <= i; j++ {
			num += 1
		}
	}
	return num
}

func getFormatedNumber(num int) []int {
	var arr = []int{1}
	j := 2
	for i := 1; i < num; i++ {
		arr = append(arr, arr[i-1]+j)
		j += 1
	}
	return arr
}

func printMatrix(mx [][]int) {
	for _, val := range mx {
		fmt.Println(val)
	}
}
