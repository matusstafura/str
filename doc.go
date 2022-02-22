/*
# String Helpers for Go Lang

## Usage

import (
	"fmt"
	"github.com/matusstafura/str"
)

func main() {
	append := str.Append("Matus", " Stafura")
	fmt.Println(append) // 'Matus Stafura'

	before := str.Before("Do not let the behavior of others destroy your inner peace. - Dalai Lama"," -")
	fmt.Println(before) // 'Do not let the behavior of others destroy your inner peace.'

	length := str.Length("Acta non verba.")
	fmt.Println(length) // 15

	lower := str.Lower("JANUARY")
	fmt.Println(lower) // january

	limit := str.Limit("This is a string!", 6)
	fmt.Println(limit) // This i

	reverse := str.Reverse("Never give up")
	fmt.Println(reverse) // pu evig reveN
}

## Tests

go test str_test.go -v

## Contributions

Contributions are welcome and will be fully credited at https://github.com/matusstafura/str

## License

This project is open-sourced software licensed under the MIT license.
*/

package str
