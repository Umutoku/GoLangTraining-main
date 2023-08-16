package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	var todo Todo
	json.Unmarshal([]byte(body), &todo)
	fmt.Println(bodyString)
	fmt.Println(todo)

}

func Demo2() {
	todo := Todo{
		UserId:    1,
		Id:        15,
		Title:     "deniyoruz",
		Completed: false,
	}
	jsonTodo, _ := json.Marshal(todo)

	response, _ := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json", bytes.NewBuffer(jsonTodo))

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	var todor Todo
	json.Unmarshal([]byte(body), &todor)
	fmt.Println(bodyString)
	fmt.Println(todor)

}
