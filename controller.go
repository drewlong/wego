package main

import (
  "os"
  "io/ioutil"
)

func controllerAction(file string, method string, mode string) string{
  pwd, _ := os.Getwd()
  rawFile, e := ioutil.ReadFile(pwd+"/templates/"+file)
  if e != nil{
    return "Error: No Path to Action."
  }
  body := string(rawFile)
  return body
}
