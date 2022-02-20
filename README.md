# String Helpers for Go Lang

## Usage

```go
package main

import (
	"fmt"
	"github.com/matusstafura/str"
)

func main() {
	lower := str.Lower("KATE")
	fmt.Println(lower) // kate

	limit := str.Limit("This is a string!", 6)
	fmt.Println(limit) // This i

	reverse := str.Reverse("Never give up")
	fmt.Println(reverse) // pu evig reveN

	reverse := str.Length("Acta non verba.")
	fmt.Println(reverse) // 15
}
```

## Tests

```bash
go test str_test.go -v
```

## Contributions

Contributions are welcome and will be fully credited.

## License

This project is open-sourced software licensed under the MIT license.
