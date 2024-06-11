package main

import (
  "fmt"
  "os"
  "os/exec"
  "path"
)


func main(){
  exec.Command("clear")  
  checkConfigDir();
  fmt.Println("Welcome to Walgo!");
  

}


func checkConfigDir() {
  PATH := path.Join(os.Getenv("HOME"), ".config/Walgo");
  _, err := os.Stat(PATH);
  
  if err == nil {
    return;
  } else if os.IsNotExist(err) {
      fmt.Println("Config ", PATH, " doesn't exist");
      err = os.MkdirAll(PATH, 0775);
      if err != nil {
          fmt.Println("Error creating directory", err)
          os.Exit(1)
      }
      
      fmt.Println("Config directory created at: ", PATH);
      os.Exit(0);
  } else {
    fmt.Println("Error checking file:", err)
    os.Exit(1);
  }

}

