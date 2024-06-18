package main

import "fmt"

func main() {

	sliced_array := []string{}

	fmt.Println(len(sliced_array))
	fmt.Println(cap(sliced_array))
	fmt.Println(sliced_array)
	
	slicedArray := []string{"foo", "bar", "baz", "quux"}
	
	
	fmt.Println(len(slicedArray))
	fmt.Println(cap(slicedArray))
	fmt.Println(slicedArray)


	//slice
	var myarray = [...]string{"foo", "bar", "baz", "quux"}

	mySlice := myarray[0:2]

	fmt.Println(mySlice)

}
