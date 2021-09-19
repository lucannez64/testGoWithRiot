package rest


import (
  "net/http"
  "fmt"
  "log"
)

func Server(address string){
  fmt.Printf("Starting server at %s\n", address)
  if err := http.ListenAndServe(address, nil); err != nil {
      log.Fatal(err)
  }
}
