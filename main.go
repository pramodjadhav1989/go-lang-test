package main

import ("fmt"
"net/http"
"log"
"encoding/json"
)

type Stock struct{
Id int `json:id`
StkCode string `json:stkCode`
}

var Stocks[] Stock




func getHello(w http.ResponseWriter,r *http.Request){
w.Header().Add("Content-Type","application/json")
json.NewEncoder(w).Encode("Hello word")
}

func addStock(w http.ResponseWriter,r *http.Request){
	w.Header().Add("Content-Type","application/json")
	var stock Stock
	_=json.NewDecoder(r.Body).Decode(&stock)
	Stocks=append(Stocks,stock)
	json.NewEncoder(w).Encode(Stocks)
}

func getAllStock(w http.ResponseWriter,r *http.Request){
w.Header().Add("Content-Type","application/json")
json.NewEncoder(w).Encode(Stocks)
}

func rout(){
http.HandleFunc("/Hello",getHello)
http.HandleFunc("/Stock/getAll",getAllStock)
http.HandleFunc("/Stock/Manage",addStock)
log.Fatal(http.ListenAndServe(":8080",nil))
}

func main(){

	fmt.Println("hi...")
	Stocks=append(Stocks,Stock{Id:1,StkCode:"XYZ"})
	rout()

}