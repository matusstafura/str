# String Helpers for Go Lang

## Usage

```go
package main

import (
	"fmt"
	"github.com/matusstafura/string-helpers/str"
)

func main() {
	lower := str.Lower("KATE")
	fmt.Println(lower) // kate

	limit := str.Limit("This is a string!", 6)
	fmt.Println(limit) // This i
}
```

## Tests

```bash
go test str/str_test.go -v
```

## Contributions

Contributions are welcome and will be fully credited.

## License

This project is open-sourced software licensed under the MIT license.
