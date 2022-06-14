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

func (a DrawItem)next() {
	
}

func (a DrawItem)prev() {
	
}

func (a DrawItem)move(pageNumber) {
	
}

func (a DrawItem)parser(pageNumber) {
	
}