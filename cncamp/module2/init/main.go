package main

import (
	"fmt"

	_ "github.com/cncamp/golang/examples/module2/init/a"
	_ "github.com/cncamp/golang/examples/module2/init/b"
)

func init() {
	fmt.Println("main init")
}
func main() {

}
