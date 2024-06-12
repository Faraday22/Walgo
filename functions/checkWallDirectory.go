package functions

import (
  "fmt"
)

/* 

  Begining of finding directory to set in condif file if its default text

*/

var correctInput = false
func CheckUserWallDir() {
  if correctInput == false {
    fmt.Println("What is the Directory you keep you wallpapers in?");
    var userInputDir string = "";
    fmt.Scanln(&userInputDir);

    fmt.Println("Is this the correct directory?");
    fmt.Println("Y/n");
    var isCorrectDir string = "";
    fmt.Scanln(&isCorrectDir);
    if isCorrectDir == "Y" || isCorrectDir == "y" {
        // Continues
    } else if isCorrectDir == "N" || isCorrectDir == "n" {
      CheckUserWallDir();
    } else {
      fmt.Println("That isn't a option");
      CheckUserWallDir();
    }
  }
} 
