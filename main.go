package main;

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
	"utils/draw"
	"utils/ssh"
)

func main() {
	fmt.Println("");
	draw.init();
	ssh.init();
}