package a

import (
	"fmt"

	_ "github.com/cncamp/golang/examples/module2/init/b"
)

func init() {
	fmt.Println("init from a")
}
