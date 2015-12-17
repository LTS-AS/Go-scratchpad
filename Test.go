package main

import (
  "io/ioutil"; 
  "fmt"
  "time"
  "strings"
  "os"
  "path/filepath"
  )


func main() {
  contents,_ := ioutil.ReadFile("Gruppenavn.yml")
  fmt.Println(string(contents))
  fmt.Println(time.Now())
  ioutil.WriteFile("filename.yml", contents, 0644)
  fmt.Printf("%q\n", strings.Split(string(contents), "\r\n  "))
  os.Mkdir("." + string(filepath.Separator) + "Testmappe", 0777)
  os.Chdir(Testmappe)
}
