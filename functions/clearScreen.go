package functions

import (

  "os"
  "os/exec"

)

func Clear_screen(){

  cmd := exec.Command("clear");  
  cmd.Stdout = os.Stdout;
  cmd.Run();

}

