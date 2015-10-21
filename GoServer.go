package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func ReplyNameGorilla(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"] // variable name is case sensitive
	w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
}

func ReplyTodo(w http.ResponseWriter, r *http.Request) {

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)

}

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}
type Todos []Todo

func main() {
	mx := mux.NewRouter()
	mx.HandleFunc("/", SayHelloWorld)
	mx.HandleFunc("/todo", ReplyTodo)
	mx.HandleFunc("/{name}", ReplyNameGorilla) // variable name is case sensitive

	http.ListenAndServe(":8080", mx)
}
