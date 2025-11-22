package zeros

import "fmt"

func duplicateZeros(arr []int) {
	zeroCount := 0
	lastIndex := len(arr) - 1

	for i := 0; i <= lastIndex-zeroCount; i++ {
		if arr[i] == 0 {
			if i == lastIndex-zeroCount {
				arr[lastIndex] = 0
				lastIndex--
				break
			}

			zeroCount++
		}
	}

	newLastIndex := lastIndex - zeroCount

	for j := newLastIndex; j >= 0; j-- {
		if arr[j] == 0 {
			arr[j+zeroCount] = 0
			zeroCount--
			arr[j+zeroCount] = 0
		} else {
			arr[j+zeroCount] = arr[j]
		}
	}

	fmt.Println(lastIndex, zeroCount)
	fmt.Println(arr)
}

func main() {
	numbers := []int{1, 0, 2, 3, 0, 0, 5, 6}
	duplicateZeros(numbers)
}
