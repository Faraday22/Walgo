package functions

import (

  "os"
  "os/exec"

)


// In the name clears the screen
func Clear_screen(){

  cmd := exec.Command("clear");  
  cmd.Stdout = os.Stdout;
  cmd.Run();

}

