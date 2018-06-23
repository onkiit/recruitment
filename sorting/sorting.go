package main

import (
	"fmt"
)

func initialRun(){
	var input int
	fmt.Println("Please select below: ")
	fmt.Println("1. Problem I -> Show Vertical Bar")
	fmt.Println("2. Problem II -> Insertion Sort with Visualization")
	fmt.Print("Select Number: ")
	fmt.Scan(&input)

	switch input {
		case 1:
			arr := createArr()
			fmt.Println()
			showBar(arr)
		case 2:
			arr:= createArr()
			fmt.Println()
			insertionSort(arr)
		default:
			fmt.Println("Not Available")
	}	
}

func createArr() (arr []int){ //array builder
	var size, input int
	fmt.Print("Array size: ")
	fmt.Scan(&size)
	for i := 1;i <= size;i++ {
		fmt.Printf("Array %d : ", i)
		fmt.Scan(&input)
		arr = append(arr, input)
	}
	return
}

func findMax(arr []int) (max int) { //find maximum value of an array. used for vertical bar height
	max = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return
}

func showBar(arr []int) {
	height := findMax(arr)
	for i := height;i > 0; i-- {
		for j := 0;j < len(arr);j++ {
			if i <= arr[j]{
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	for _, val := range arr{
		fmt.Print(val)
		fmt.Print(" ")
	}
}


func insertionSort(arr []int) {
	showBar(arr)
	fmt.Println()
	for i := len(arr) - 1; i > 0; i-- { //comparing from tail to head
		var temp int
		j := i-1
		for j >= 0 && arr[i] < arr[j] {
			temp = arr[j]
			arr[j] = arr[i]
			arr[i] = temp
			showBar(arr)
			fmt.Println()
			j = j - 1 //loop till j >= 0
			
		}
	}
}

func main(){
	initialRun()
}