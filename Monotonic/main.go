package main

import "fmt"

func IsMonotonic(array []int) bool {

	if len(array) <=2 {
		return true
	}
	direction := array[1] - array[0]

	for i:=2;i < len(array);i++{

		if direction == 0{
			direction = array[i] - array[i-1]
		}
		if breaksDirection(direction, array[i-1],array[i]){
			return false
		}
	}
	return true
}

func breaksDirection(direction, prev,curr int) bool  {
	difference := curr - prev
	if direction >0 {
		return difference < 0
	}
	return difference > 0
}
//
//func IsMonotonic(array []int) bool {
//	if len(array) <=2 {
//		return true
//	}
//	Incr := false
//	Descr := false
//
//	for i:=1;i<len(array);i++{
//		if array[i] > array[i-1]{
//			Incr = true
//			break
//		}
//		if array[i] < array[i-1]{
//			Descr = true
//			break
//		}
//
//	}
//	if ! Incr && ! Descr {
//		return true
//	}
//
//
//	for i:=1;i < len(array);i++{
//		if Descr && array[i-1] < array[i]{
//			return false
//		}
//		if  Incr && array[i-1] > array[i]{
//			return false
//
//		}
//	}
//	return true
//}

func testmono(array []int)  {
	return all
}




func main() {
	//a := []int{1,-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	a := []int{1,1,1,1,1,1,1,1,1,1,1}
	fmt.Println("Array is Monotonic ", IsMonotonic(a))


}
