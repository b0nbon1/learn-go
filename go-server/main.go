package main

import (
  "fmt"
  "log"
  "net/http"
)

func formHandler(res http.ResponseWriter, req *http.Request) {
  if err := req.ParseForm(); err != nil {
    fmt.Fprintf(res, "ParseForm() err: %v", err)
    return
  }
  fmt.Fprintf(res, "POST request successful\n")
  name := req.FormValue("name")
  address := req.FormValue("address")
  age := req.FormValue("age") 

  fmt.Fprintf(res, "Name = %s\n", name)
  fmt.Fprintf(res, "Address = %s\n", address)

  fmt.Fprintf(res, "Age = %s\n", age)
}

func hellohandler(res http.ResponseWriter, req *http.Request) {
  if req.URL.Path != "/home" {
    http.Error(res, "404 not found", http.StatusNotFound)
    return
  }

  if req.Method != "GET" {
    http.Error(res, "method is not supported", http.StatusMethodNotAllowed)
    return
  }

  fmt.Fprint(res, "hello world, Gophers!")
}

func main(){
  fileserver := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileserver)
  http.HandleFunc("/form", formHandler)
  http.HandleFunc("/hello", hellohandler)

  fmt.Printf("Starting server at port 8080\n")

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
