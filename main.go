package main

import (
	"fmt"

	"test1/files-structure/group"
	"test1/files-structure/group/pkg4"
	"test1/files-structure/pkg1"
	"test1/files-structure/pkg2"
)

func main() {
	fmt.Println("[main] Hello World")
	pkg1.Test()
	pkg2.Test()
	group.Test()
	pkg4.Test()
	pkg4.Test1()
}
