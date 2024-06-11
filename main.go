package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/faraday22/Walgo/startup"
)

func clear_screen(){

  cmd := exec.Command("clear");  
  cmd.Stdout = os.Stdout;
  cmd.Run();
}


func main(){
  clear_screen();
  startup.CheckConfigDir();
  startup.CheckConfigFile();
  fmt.Println("Welcome to Walgo!");  

}
