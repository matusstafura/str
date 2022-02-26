# String Helpers for Go Lang

## Usage

```go
package main

import (
	"fmt"
	"github.com/matusstafura/str"
)

func main() {
	// str.After() returns everything after first occurence
	after := str.After("Do not let the behavior of others destroy your inner peace. - Dalai Lama","- ")
	fmt.Println(after) // 'Dalai Lama'

	// str.Append(...) concatenates strings
	append := str.Append("Matus", " Stafura")
	fmt.Println(append) // 'Matus Stafura'

	// str.Before() returns everything before first occurence
	before := str.Before("Do not let the behavior of others destroy your inner peace. - Dalai Lama"," -")
	fmt.Println(before) // 'Do not let the behavior of others destroy your inner peace.'

	// str.EndsWith() checks if string ends with the given target string
	endsWith := str.EndsWith("abcdef","ef")
	fmt.Println(endsWith) // true

	// str.Length() returns lenght of the given string
	length := str.Length("Acta non verba.")
	fmt.Println(length) // 15

	// str.Limit() truncates the string to the specified length
	limit := str.Limit("This is a string!", 6)
	fmt.Println(limit) // This i

	// str.Lower() converts the string to lowercase
	lower := str.Lower("JANUARY")
	fmt.Println(lower) // january

	// str.Mask() converts the string to lowercase
	mask := str.Mask("4242 4242 4242 4242 4242", 4, "#")
	fmt.Println(mask) // 4242####################

	// str.Reverse() reverses the string
	reverse := str.Reverse("Never give up")
	fmt.Println(reverse) // pu evig reveN

	// str.StartsWith() checks if string starts with the given target string
	startsWith := str.startsWith("abcdef","ab")
	fmt.Println(startsWith) // true

	// str.Upper() converts the string to uppercase
	upper := str.Upper("Be the friend you wish you had.")
	fmt.Println(upper) // BE THE FRIEND YOU WISH YOU HAD.
}
```

## Tests

```bash
go test ./... -v
```

## Contributions

Contributions are welcome and will be fully credited.

## License

This project is open-sourced software licensed under the MIT license.
