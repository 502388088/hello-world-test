package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"log"
)

var muxRouter *mux.Router;

func main() {

	fmt.Println("=======================================");
	fmt.Println("=======Starting Hello World Test=======");
	fmt.Println("=======================================");

	muxRouter = mux.NewRouter();

	muxRouter.Methods("GET").Path("/hello").HandlerFunc(HandleHelloWorld);

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080";
		log.Printf("Port number is not set using default: %s\n", port );
	}

	err:= http.ListenAndServe(":"+port, muxRouter)
	if(err != nil){
		log.Println(err);
		os.Exit(1);
	}
}


func HandleHelloWorld (w http.ResponseWriter, r *http.Request) {
	defer  r.Body.Close();

	w.Header().Set("Content-Type", "text/plain");
	w.WriteHeader(http.StatusOK);

	fmt.Fprint(w, "Hello World!");
}