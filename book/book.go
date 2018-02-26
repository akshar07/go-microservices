package book

import(
	"encoding/json"
	"net/http"
)

type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string	`json:"isbn"`
}

func (b Book) ToJSON() []byte {
	ToJSON,err:=json.Marshal(b)
	if err !=nil{
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Book{
	book:=Book{}
	err:= json.Unmarshal(data,&book)
	if err !=nil{
		panic(err)
	}
	return book
}

var Books=[]Book{
	Book{Title:"CloudNative Go" ,Author: "Akshar Takle",ISBN:"123456"},
	Book{Title:"CloudNative Primer", Author: "Akshar Takle",ISBN:"123133456"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request){
	b,err:=json.Marshal(Books)
	if err!=nil{
		panic(err)
	}
	w.Header().Add("Content-Type","application/json; charset-utf-8")
	w.Write(b)
}