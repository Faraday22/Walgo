package main

import (
  "fmt"
  "os/exec"
  "github.com/faraday22/Walgo/startup"
)

func clear_screen(){

  exec.Command("clear").Run();  

}


func main(){
  clear_screen();
  startup.CheckConfigDir();
  startup.CheckConfigFile();
  fmt.Println("Welcome to Walgo!");  

}
