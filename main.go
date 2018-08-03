package main

// [START import_statements]
import (
	//"fmt"
  "log"
  //"os"
	"strings"

	"io/ioutil"
	"net/http"
  //"html/template"
  //"path/filepath"

	"google.golang.org/appengine"
)

type MyHandler struct {
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./www")))
  //http.HandleFunc("/", serveTemplate)
  log.Println("Listening...")
  appengine.Main()
}


func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    log.Println(path)
    data, err := ioutil.ReadFile(string(path))

    if err == nil {
        var contentType string

        if strings.HasSuffix(path, ".css") {
            contentType = "text/css"
        } else if strings.HasSuffix(path, ".html") {
            contentType = "text/html"
        } else if strings.HasSuffix(path, ".js") {
            contentType = "application/javascript"
        } else if strings.HasSuffix(path, ".png") {
            contentType = "image/png"
        } else if strings.HasSuffix(path, ".svg") {
            contentType = "image/svg+xml"
        } else {
            contentType = "text/plain"
        }

        w.Header().Add("Content Type", contentType)
        w.Write(data)
    } else {
        w.WriteHeader(404)
        w.Write([]byte("404 My dear - " + http.StatusText(404)))
    }
}
