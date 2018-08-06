package main

// [START import_statements]
import (
	"fmt"
  "log"
  //"os"
	//"strings"

	//"io/ioutil"
	"net/http"
  //"html/template"
  //"path/filepath"

	"google.golang.org/appengine"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./www")))
  http.HandleFunc("/shop", ServeShop)
  log.Println("Listening...")
  appengine.Main()
}


func ServeShop(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "!!!Shop!!!")
}
