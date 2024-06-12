package startup

import ( 

  "fmt"
  "os"
  "path"

)

func CheckConfigDir() {
  // Sets path for the directory of applications config
  PATH := path.Join(os.Getenv("HOME"), ".config/Walgo");
  _, err := os.Stat(PATH);
  
  // If it exist exits program
  if err == nil {
    return;
  } else if os.IsNotExist(err) {
      // If the directory doesn't exist its created
      fmt.Println("Config ", PATH, " doesn't exist");
      err = os.MkdirAll(PATH, 0775);
      if err != nil {
          fmt.Println("Error creating directory", err)
          os.Exit(1)
      }
      
      fmt.Println("Config directory created at: ", PATH);
      // Calls other function to find if config file
      // is generated
      CheckConfigFile();
      fmt.Println("Rerun Walgo")
      os.Exit(0);
  } else {
    fmt.Println("Error checking file:", err)
    os.Exit(1);
  }

}
