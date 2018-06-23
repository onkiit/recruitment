package queue

import "fmt"

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type Item struct {
	Head	int
	Tail	int
	Size	int
	Array	[]int
}

func New(size int) Queue {
	i := Item{}
	i.Size = size
	return nil
}

func Push(key interface{}) {
	fmt.Println(key)
}

func Pop() interface{}{
	return nil
}

func Contains(key interface{}) bool {
	return true 
}

func Len() int {
	return 0
}

func Keys() []interface{}{
	return nil
}