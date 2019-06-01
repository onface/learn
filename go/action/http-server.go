package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

type ExampleData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonController(w http.ResponseWriter, r *http.Request) {
	var dataVO = ExampleData{Name: "nimo", Age: 21}
	dataJO, _ := json.Marshal(dataVO)
	_, _ = fmt.Fprintf(w, string(dataJO))
}

func queryController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析请求
	queryJSON, _ := json.Marshal(r.URL.Query())
	_, _ = fmt.Fprintf(w, string(queryJSON))
}
func postController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析请求
	value := r.PostFormValue("value")
	_, _ = fmt.Fprintf(w, "POST value: " + value)
}
func htmlController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprintf(w, `
	<ul>
		<li><a  target="_blank" href="/json">/json</a></li>
		<li><a  target="_blank" href="/query?a=1&b=some">/query?a=1&b=some</a></li>
	</ul>
	<form method="post" action="/post" >
		<input name="value" value="onface" />
		<button type="submit" >submit</button>
	</form>
`)
}
func main() {
	http.HandleFunc("/json", jsonController)
	http.HandleFunc("/query", queryController)
	http.HandleFunc("/post", postController)
	http.HandleFunc("/", htmlController)
	err := http.ListenAndServe(":1313", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}