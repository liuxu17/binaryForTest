package helper

func Swap(arr []int, i int, j int) {
	if len(arr) <= i || len(arr) <= j {
		return
	}
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}