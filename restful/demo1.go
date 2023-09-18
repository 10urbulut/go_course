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
	response, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if error != nil {
		fmt.Println(error)
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(body, &todo)

	fmt.Println(todo.Id)
}
func Demo2() {
	todo := Todo{1, 2, "Alışveriş yapılacak", false}

	jsonTodo, _ := json.Marshal(todo)

	response, error := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if error != nil {
		fmt.Println(error.Error())
	}

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

	json.Unmarshal(body, &todo)

	fmt.Println(todo)

}
