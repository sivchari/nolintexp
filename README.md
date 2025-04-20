# nolintexp

golangci-lint provides way to ignore diagnostics error by using `nolint` directive. This feature is useful, but content that should have been corrected may be left ignored.

nolintexp supports new unique directive called `nolintexp` that indicates the deadline when we 
can remain `nolint` directive.

## Installation

```go
go install github.com/sivchari/nolintexp/cmd/nolintexp@latest
```

## Format

Currently, `YYYY-MM-DD`is only supported and you can specify the deadline in this format.
For example, if you want to set the expiration date to 2025/01/01, can specify //nolintexp:2025-01-01.

`nolintexp` must be included with comment group that has `nolint` directive.

### Good

```go
package a

//nolint:lint
//nolintexp:2025-01-01
type Int int
```

### Bad

```go
package a

// The line break means //nolint:lint and //nolintexp:2025-01-01 are separately.
//nolint:lint

//nolintexp:2025-01-01
type Int int
```
