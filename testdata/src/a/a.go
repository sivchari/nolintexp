//nolint:file // reason // want "nolint directive \"file\" expired on 2023-10-01"
//nolintexp:2023-10-01
package a

//nolint:block // want "nolint directive \"block\" expired on 2024-01-01"
//nolintexp:2024-01-01
var (
	foo int
	bar int
	baz int
)

func f() {
	//nolint:variable // want "nolint directive \"variable\" expired on 2024-01-01"
	//nolintexp:2024-01-01
	var variable int
	_ = variable

	//nolint:variable
	//nolintexp:2025-01-01 // ok due to not exceeding expiration date
	var variable2 int
	_ = variable2

	//nolint:variable
	var variable3 int
	_ = variable3
}

//nolint:type // want "nolint directive \"type\" expired on 2024-01-01"
//nolintexp:2024-01-01
type Int int
