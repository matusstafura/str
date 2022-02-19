package main

import (
	"fmt"

	"github.com/matusstafura/string-helpers/str"
)

func main() {
	lower := str.Lower("KATE")
	fmt.Println(lower)
	
	limit := str.Limit("This is a string!", 6)
	fmt.Println(limit)
	
	reverse := str.Reverse("Never give up")
	fmt.Println(reverse) // pu evig reveN
}

