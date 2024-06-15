package functions

import (

  "fmt"
  "os/exec"
  "bytes"
)


var windowManagers = [1]string{"sway"}

func findRunningWM(windowManagers string) bool {

   cmd := exec.Command("pgrep", windowManagers)
    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    return err == nil

}





func WmCommands() {

    for _, windowManagersProcess := range windowManagers {
      if findRunningWM(windowManagers[0]) {
          /* put function that takes windowManagersProcess
          *  used the varible and has if BLANK = WM then
          *  uses the specifi windowManagers command to change the wallpaper
          */ 
          fmt.Printf("Process %s is running.\n", windowManagersProcess)
        }
    }

}
