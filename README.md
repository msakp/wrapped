# wrapped
<h3>first flexible go struct wrapper</h3>

## Installation
`go get -u github.com/msakp/wrapped@latest`


## Quick Start

for given structs **Book** and **BookView**:
```go
type Book struct{
	ID     int
	Title  string
	Cost   int
}

type BookView struct{
	Id     int
	Title  string
	cost   int
}

```
```go
book := Book{ID: 10, Title: "Book", Cost: 200}

bookV := wrapped.To[BookView](&book)

fmt.Printf("bookView: %#v\n", bookV)
/*
bookView: &main.BookView{Id:10, Title:"Book", cost:200}
*/
```

## TODO
- multiple input struct support 
- tag support for more flexibility
