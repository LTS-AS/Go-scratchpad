package main

import (
    "io/ioutil"
)

func main() {
    // read whole the file
    b, err := ioutil.ReadFile("Gruppenavn.yml")
    if err != nil {
        panic(err)
    }

    // write whole the body
    err = ioutil.WriteFile("Output.yml", b, 0644)
    if err != nil {
        panic(err)
    }
}
