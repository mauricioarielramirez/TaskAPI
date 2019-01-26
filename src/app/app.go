package app

import (
	"TaskAPI/src/controller"
	"fmt"
)

func DoSomething () {
	fmt.Println(controller.ListUsers())
}
