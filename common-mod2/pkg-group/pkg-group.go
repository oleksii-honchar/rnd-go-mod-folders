package pkgGroup

import (
	cm1 "common-mod1"
	"fmt"
)

func Test() {
	fmt.Println("[common-mod2] pkg-group(")
	fmt.Println("Calling for cm1:")
	cm1.Test()
	fmt.Println("[common-mod2] pkg-group)")
}
