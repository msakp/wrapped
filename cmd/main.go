package main

import (
	"encoding/json"
	"fmt"

	"github.com/msakp/wrapped"
)


type BookModel struct{
	ID int64
	Title string
	Price int32
	Selled int64
	Limit int64
}

type BookView struct{
	Id int64
	Title string
	Price int32
	Selled int64
	Limit int64
}


type BookStat struct{
	Id int64
	SalesPercent float64

}

func dump(data any){
    b,_:=json.MarshalIndent(data, "", "  ")
		fmt.Println("===================================")
		fmt.Printf("RESULT: \n%s\n", string(b))
}


func main(){
	book := BookModel{
		ID: 10,
		Title: "new book",
		Price: 20,
		Selled: 210,
		Limit: 3000,
	}


	fmt.Println("Hello world!")
	bookV := wrapped.To[BookView](&book)
	dump(&bookV)
}


















