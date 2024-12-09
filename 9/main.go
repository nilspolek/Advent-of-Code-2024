package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", "input.txt", "File Path")
	flag.Parse()

	input, _ := os.ReadFile(filePath)
	inputString := string(input)
	disk := &Deque[int]{}

	current := 0
	for index, c := range inputString {
		num, _ := strconv.Atoi(string(c))
		if index%2 == 0 {
			for range num {
				disk.PushBack(current)
			}
			current++
			continue
		}
		for range num {
			disk.PushBack(-1)
		}
	}
	total := 0
	for i := 0; !disk.IsEmpty(); i++ {
		num, _ := disk.PopFront()
		for num == -1 {
			num, _ = disk.PopBack()
		}
		total += (i * num)
	}
	fmt.Println(total)
}

type Deque[T any] struct {
    items []T
}

func (d *Deque[T]) PushFront(item T) {
    d.items = append([]T{item}, d.items...)
}

func (d *Deque[T]) PushBack(item T) {
    d.items = append(d.items, item)
}

func (d *Deque[T]) PopFront() (T, bool) {
    if len(d.items) == 0 {
        var zeroValue T
        return zeroValue, false
    }
    item := d.items[0]
    d.items = d.items[1:]
    return item, true
}

func (d *Deque[T]) PopBack() (T, bool) {
    if len(d.items) == 0 {
        var zeroValue T
        return zeroValue, false
    }
    item := d.items[len(d.items)-1]
    d.items = d.items[:len(d.items)-1]
    return item, true
}

func (d *Deque[T]) Front() (T, bool) {
    if len(d.items) == 0 {
        var zeroValue T
        return zeroValue, false
    }
    return d.items[0], true
}

func (d *Deque[T]) Back() (T, bool) {
    if len(d.items) == 0 {
        var zeroValue T
        return zeroValue, false
    }
    return d.items[len(d.items)-1], true
}

func (d *Deque[T]) IsEmpty() bool {
    return len(d.items) == 0
}
