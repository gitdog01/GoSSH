package utils;

import (
    "fmt"
)

struct DrawItem {
	frame uint8
	page uint8
}

func DrawInit(frame) {
	fmt.Println("draw init")
	instance := DrawItem{frame,0}
	return *instance
}

func (a DrawItem)run() {
	
}