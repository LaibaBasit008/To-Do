package main

import ("net/http"
"log"
"html/template"
)

var tmpl *template.Template
type Todo struct {
	Item string
	Done bool
}
type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request){
	data := PageData{
		Title:"THINGS 2 DO",
		Todos: []Todo{
			{Item:"Water the plants", Done: true},
			{Item:"Jump off the cliff ^", Done: false},
		},
	}
	tmpl.Execute(w, data)
}
func main(){
	mux := http.NewServeMux()
	tmpl= template.Must(template.ParseFiles("templates/index.gohtml"))
	fd := http.FilseServer(http.Dir("./static"))
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":5010",mux))

}