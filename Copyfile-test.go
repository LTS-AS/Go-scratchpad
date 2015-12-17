package main

 import (
     "io"
     "fmt"
     "os"
 )

 func main () {
     // open files r and w
     r, err := os.Open("input.txt")
     if err != nil {
         panic(err)
     }
     defer r.Close()

     w, err := os.Create("output.txt")
     if err != nil {
         panic(err)
     }
     defer w.Close()
	 
	
     // do the actual work
     n, err := io.Copy(w, r)
     if err != nil {
         panic(err)
     }


    fmt.Printf("Copied %v bytes\n", n)
	fmt.Printf("Copied input.txt, and created output.txt")
	
	 err2 := os.Remove("input2.txt")
	 if err2 != nil {
		 panic(err2)
	 }
 }
