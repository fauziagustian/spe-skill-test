package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// nomor 1 narsistik
	fmt.Println("Naricissistic")
	fmt.Println(narcissistic(153))
	fmt.Println(narcissistic(111))

	fmt.Println("====================")

	//nomor 2 parity outlier
	fmt.Println("Parity Outlier")
	ex1 := []int{2, 4, 0, 100, 11, 2602, 36}
	ex2 := []int{160, 3, 1719, 19, 11, 13, -21}
	ex3 := []int{11, 13, 15, 19, 9, 13, -21}
	fmt.Println(parityOutlier(ex1))
	fmt.Println(parityOutlier(ex2))
	fmt.Println(parityOutlier(ex3))

	fmt.Println("====================")
	//nomor 3 NEEDLE IN THE HAYSTACK
	fmt.Println("NEEDLE IN THE HAYSTACK")
	v1 := []string{"red", "blue", "yellow", "black", "grey"}
	fmt.Printf("%v", findNeedle(v1, "blue"))

	fmt.Println("====================")
	//nomor 4 NEEDLE IN THE HAYSTACK
	fmt.Println("blue ocean ")
	b1 := []int{1, 2, 3, 4, 6, 10}
	fmt.Printf("%v", blueOceanReverse(b1, 1))
}

func narcissistic(x int) bool {
	arr := splitToGigits(x)
	total := 0

	lenA := len(arr)

	for _, v := range arr {
		total += int(math.Pow(float64(v), float64(lenA)))
	}

	if x == total {
		return true
	} else {
		return false
	}
}

func splitToGigits(n int) []int {
	var ret []int

	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}

	reverseInt(ret)

	return ret
}

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func parityOutlier(x []int) string {

	countOdd := 0
	countEven := 0
	keyodd := 0
	keyeven := 0

	for i := 0; i < len(x); i++ {
		if x[i]%2 == 0 {
			countEven++
			keyeven = i
		} else {
			countOdd++
			keyodd = i
		}
	}

	if countEven == 1 {
		return strconv.Itoa(x[keyeven])
	} else if countOdd == 1 {
		return strconv.Itoa(x[keyodd])
	} else {
		return "false"
	}
}

func findNeedle(x []string, find string) []string {

	var result []string

	for _, v := range x {
		if v == find {
			result = append(result, v)
		}
	}

	return result
}

func blueOceanReverse(x []int, y int) []int {

	var result []int

	for _, v := range x {
		if v != y {
			result = append(result, v)
		}

	}

	return result
}
