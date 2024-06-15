package functions

import (
  "fmt"
)

var correctInput = false
var UserInputDir string = ""
func CheckUserWallDir() {
  if correctInput == false {
    fmt.Println("What is the Directory you keep you wallpapers in?");
    fmt.Scanln(&UserInputDir);

    fmt.Println("Is this the correct directory?");
    fmt.Println("Y/n");
    var isCorrectDir string = "";
    fmt.Scanln(&isCorrectDir);
    if isCorrectDir == "Y" || isCorrectDir == "y" {
        Clear_screen()
        // Warning Message
        fmt.Println("You will not be able to change the\n directory again unless through the config file\n located at ~/.config/Walgo")       
        fmt.Println("")
    } else if isCorrectDir == "N" || isCorrectDir == "n" {
      CheckUserWallDir();
    } else {
      fmt.Println("That isn't a option");
      CheckUserWallDir();
    }
  }
}
