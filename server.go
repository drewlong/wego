package main

import(
  "net/http"
  "fmt"
  "flag"
)

func SetPort() string{
  portPtr := flag.String("p", "7000", "Port to run the API server.")

  flag.Parse()

  return fmt.Sprintf(":%s", *portPtr)
}

func takeAction(w http.ResponseWriter, r *http.Request){
  req := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
  resBody := handleReq(req)
  fmt.Fprintf(w, resBody)
  fmt.Println(req)
}

func main(){
  port := SetPort()
  http.HandleFunc("/", takeAction)
  status := http.ListenAndServe(port, nil)
  if status != nil{
    fmt.Println(status)
  }
}
