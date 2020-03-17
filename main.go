package main
	
import (
	"fmt"
  "net/http"
)

func getOrgStats(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}


func main() {
	http.HandleFunc("/getOrgStats", getOrgStats)

	http.ListenAndServe(":8090", nil)
}