package slice

import "fmt"

func slice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumAllSlice(slicedata ...[]int) []int {

	length := len(slicedata)
	//result := make([]int, length)
	// OR below one
	var result []int
	for i := 0; i < length; i++ {
		//result[i] = slice(slicedata[i])
		// OR below one
		result = append(result, slice(slicedata[i]))
	}
	fmt.Println(result)
	return result

}

func sumTail(slices ...[]int) []int {
	var result []int

	for _, v := range slices {
		fmt.Println(v)
		result = append(result, slice(v[1:])) // [2,3,5] => [3,5] // :1 -> [2,3,5]=>[2]
	}
	return result

}
