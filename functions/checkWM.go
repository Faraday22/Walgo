package functions

import (

  "os/exec"
  "bytes"
)


var windowManagers = [1]string{"sway"}

func findRunningWM(windowManagers string) bool {
  // Scans process to find out which WM is running
   cmd := exec.Command("pgrep", windowManagers)
    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    return err == nil

}




var CurrentRunningWM string = ""

// Finds which WM is running from the list
// in windowManagers array
func WmEnvironmentSearch() {

    for _, windowManagersProcess := range windowManagers {
      if findRunningWM(windowManagers[0]) {
          // exports varible to be used in changeWallpaper.go
          CurrentRunningWM = windowManagersProcess
        }
    }

}
