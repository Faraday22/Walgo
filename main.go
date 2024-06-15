package main

import (

  "github.com/faraday22/Walgo/startup"
  "github.com/faraday22/Walgo/functions" 
  
)

func main(){
  functions.Clear_screen();
  startup.CheckConfigDir();
  startup.CheckConfigFile();
  
  functions.ConfigDetails();
  functions.WmCommands();
}
