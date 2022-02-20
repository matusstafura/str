# String Helpers for Go Lang

## Usage

```go
package main

import (
	"fmt"
	"github.com/matusstafura/str"
)

func main() {
	append := str.Append("Matus", " Stafura")
	fmt.Println(append) // 'Matus Stafura'

	lower := str.Lower("JANUARY")
	fmt.Println(lower) // january

	limit := str.Limit("This is a string!", 6)
	fmt.Println(limit) // This i

	reverse := str.Reverse("Never give up")
	fmt.Println(reverse) // pu evig reveN

	length := str.Length("Acta non verba.")
	fmt.Println(length) // 15
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
