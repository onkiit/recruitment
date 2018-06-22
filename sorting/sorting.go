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

func createArr() (arr []int){
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

func findMax(arr []int) (max int) {
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

//TODO: fixing sorting
func insertionSort(arr []int) {
	for i := 0;i < len(arr);i++ {
		temp := arr[i]
		a := i-1
		for a >= 0 && arr[a] > temp {
			arr[a+1] = arr[a]
			a = a-1
		}
		arr[a+1] = temp
		fmt.Println()
		showBar(arr)
		fmt.Println()
	}
}

func main(){
	initialRun()
}