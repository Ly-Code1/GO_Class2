package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func get_content() {
	// json data
	url := "https://jsonplaceholder.typicode.com/users/1"

	res, err := http.Get(url)
	defer res.Body.Close()

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
	}
	

    var data User
    json.Unmarshal(body, &data)
    fmt.Printf("Results: %v\n", data)
    os.Exit(0)
}

func main(){
	get_content()
}