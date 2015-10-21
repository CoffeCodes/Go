package main
import ( "fmt"
"encoding/json"
"net/http"
"github.com/gorilla/mux"
)


func Index(w http.ResponseWriter,r *http.Request){

fmt.Fprintln(w,"!Welcome");
}

func TodoIndex(w http.ResponseWriter,r *http.Request){
todos := Todos {

Todo{Name : "This is Go Lang"}
Todo{Name : "SO whats?"}


}
if err := json.NewEncoder(w).Encode(todos);err != nil{
panic(err)
}


}
func TodoShow(w http.ResponseWriter,r *http.Request){
vars := mux.Vars(r);
todoId := vars["todoId"]
fmt.Fprintln(w,"Todo Show",todo)


}



