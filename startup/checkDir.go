package startup

import ( 

  "fmt"
  "os"
  "path"

)

func CheckConfigDir() {
  PATH := path.Join(os.Getenv("HOME"), ".config/Walgo");
  _, err := os.Stat(PATH);
  
  if err == nil {
    return;
  } else if os.IsNotExist(err) {
      fmt.Println("Config ", PATH, " doesn't exist");
      err = os.MkdirAll(PATH, 0775);
      if err != nil {
          fmt.Println("Error creating directory", err)
          os.Exit(1)
      }
      
      fmt.Println("Config directory created at: ", PATH);
      CheckConfigFile();
      fmt.Println("Rerun Walgo")
      os.Exit(0);
  } else {
    fmt.Println("Error checking file:", err)
    os.Exit(1);
  }

}
