package main
import (
	"fmt"
	"log"
	"net/http"
)
func formHandler(w http.ResponseWriter, r  *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w,"address = %s\n", address)
}

func HelloHandler(w http.ResponseWriter, r *http.Request){
    if r.URL.Path!= "/hello"{
	http.Error(w, "404 not found", http.StatusNotFound)

	return
}
if r.Method != "GET"{
	http.Error(w, "method is not supported", http.StatusNotFound)
	return
  }
  fmt.Fprint(w, "hello!")

}
func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/form", fileServer)  // ✅ Correct
	http.Handle("/hello", fileServer) // ✅ Correct
	

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil{
		log.Fatal(err)
	}
}New commit at 2025-02-04 23:52:09
