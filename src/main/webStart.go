package main
import (
	"net/http"
	"io"
	"log"
)
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, 世界!\n")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("D:/Projects/GOP/test/static/")))
	//http.ListenAndServe(":8080", nil)
	http.HandleFunc("/hello",HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}