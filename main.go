package main

import (
	"fmt"

	commonMod1 "common-mod1"
	pkgGroupCM1 "common-mod1/pkg-group"
	pkg1CM1 "common-mod1/pkg-group/pkg1"
	cm2 "common-mod2"
	pkgGroup "test1/files-structure/pkg-group"
	"test1/files-structure/pkg-group/pkg1"
)

func main() {
	fmt.Println("[main] Hello World")
	pkgGroup.Test()
	pkg1.Test()
	pkg1.Test1()
	commonMod1.Test()
	pkgGroupCM1.Test()
	pkg1CM1.Test()
	pkg1CM1.Test1()
	cm2.Test()
}
