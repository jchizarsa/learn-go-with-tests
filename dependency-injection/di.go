package dependencyinjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name) // Defaults the Writer to stdout, don't want this
	fmt.Fprintf(writer, "Hello, %s", name) // Allows us to pass in a writer, want this
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func main() {
	// Greet(os.Stdout, "Elodie") // Need to change type from *bytes.Buffer -> io.Writer
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
