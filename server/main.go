package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wanton-idol/TO-DO-APP/routers"
)

func main() {
	fmt.Println("TODO App")
	r := routers.Router()

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Server is started at port :3000...")
}
