package Helpers

import (
	"fmt"
	"os"
)

func Dd(args ...interface{})  {

	for _,arg := range args{
		fmt.Println(arg)
	}

	os.Exit(-1)
}
