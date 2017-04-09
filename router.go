package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

type Route struct {
    Rpath   string  `json:"path"`
    Rmethod string  `json:"method"`
    Rmode   string  `json:"mode"`
    Rfile   string  `json:"file"`
}

func getRoutes() []Route {
    pwd, _ := os.Getwd()
    body, e := ioutil.ReadFile(pwd+"/routes.json")
    if e != nil {
        fmt.Println(e.Error())
        os.Exit(1)
    }

    var c []Route
    json.Unmarshal(body, &c)
    return c
}

func handleReq(req string) string{
    params := strings.Split(req, " ")
    routes := getRoutes()
    for _, p := range routes {
      if p.Rpath == params[1]{
        res := controllerAction(p.Rfile, p.Rmethod, p.Rmode)
        return res
      }
    }
    return "File Not Found."
}
