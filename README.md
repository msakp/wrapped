# wrapped
<h3>first flexible go struct wrapper</h3>

## Installation
`go get -u github.com/msakp/wrapped@latest`


## Quick Start

for given structs **Book** and **BookView**:
```go
type Book struct {
	ID                int
	Title             string
	Cost              int
	SomeUnwantedField int
}

type BookView struct {
	Id             int
	Title          string
	Cost           int
	OtherField     string
}

```
```go
book := Book{
		ID: 10,
		Title: "Book",
		Cost: 200,
		SomeUnwantedField: 1000,
}

bookV := wrapped.To[BookView](&book)

// can acess bookV fields and methods
_ = bookV.Cost
bookV.OtherField = "yo"

fmt.Printf("%#v\n", bookV)

*/
```

## TODO
- multiple input struct support 
- tag support for more flexibility
