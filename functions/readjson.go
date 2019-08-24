package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var users Users

func CheckJson() {
	jsonFile, err := os.Open("./jsonfiles/users.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened users.json")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}

type Users struct {
	Users []User `json:"users"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twittr   string `json:twitter`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

func ReadJson() {
	jsonFile, err := os.Open("./jsonfiles/users.json")

	if err != nil {
		fmt.Println(err)
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &users)

		for i := 0; i < len(users.Users); i++ {
			fmt.Println("User Type: " + users.Users[i].Type)
			fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
			fmt.Println("User Name: " + users.Users[i].Name)
			fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
		}
	}

}
