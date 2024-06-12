package startup

import ( 

  "fmt"
  "os"
  "path"

)

func CheckConfigFile(){
  PATH := path.Join(os.Getenv("HOME"), ".config/Walgo/Walgo.toml");
  _, err := os.Stat(PATH);
  
  if err == nil {
    return;
  } else if os.IsNotExist(err) {
      fmt.Println("Config ", PATH, " doesn't exist");
      // Created Walgo.toml file 
      defaultConfig := []byte("[DIRECTORY]\n# Add Path to Wallpaper directory \n #WallpaperPath = \"Enter wallpaper directory path\"\n");
      err = os.WriteFile(PATH, defaultConfig, 0644);
      if err != nil {
          fmt.Println("Error file", err)
          os.Exit(1)
      }
      
      fmt.Println("Config file created at: ", PATH);
      fmt.Println("Rerun Walgo")
      os.Exit(0);
  } else {
    fmt.Println("Error checking file:", err)
    os.Exit(1);
  }
}
