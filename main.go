package main

import (
	"fmt"
	"os"
	"github.com/faraday22/Walgo/startup"
  "github.com/faraday22/Walgo/functions" 
  
  "github.com/BurntSushi/toml"

)

func main(){
  functions.Clear_screen();
  startup.CheckConfigDir();
  startup.CheckConfigFile();
  // Begins set up for setting data to wallpaper config
  type Config struct {
    DIRECTORY struct {
      WallpaperPath string `toml:"WallpaperPath"`
    }  
  }
  // Sets home directory value
  homeDir, err := os.UserHomeDir();
  if err != nil {
    fmt.Println(err)
    return;
  }
  // Reads info from config file
  configPath := fmt.Sprintf("%s/.config/Walgo/Walgo.toml", homeDir)
  // data is var for info in Walgo Config Config wallpaper path
  data, err := os.ReadFile(configPath)
  if err != nil {
    fmt.Println(err);
    return;
  }
  
  // Decodes values
  var config Config
  err = toml.Unmarshal(data, &config)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Welcome to Walgo!");  
  fmt.Printf("Config file path is: %s\n", config.DIRECTORY.WallpaperPath)
  

  // Asks for custom directory if the path is default in the config file
  if config.DIRECTORY.WallpaperPath == "Enter wallpaper directory path" {
    functions.Clear_screen();
    functions.CheckUserWallDir();
  } 

}
