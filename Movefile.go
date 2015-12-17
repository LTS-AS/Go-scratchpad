 package main

   import (
       "fmt"
       "os"
   )

   func main() {

       err :=  os.Rename("\Folder_A\file.txt", "\Folder_B\file.txt")

       if err != nil {
           fmt.Println(err)
           return
       }
  }
