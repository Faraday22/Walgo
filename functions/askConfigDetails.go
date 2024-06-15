package functions

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/faraday22/Walgo/startup"
)

func ConfigDetails() {
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
  fmt.Println("");
  fmt.Printf("Config directory path is: %s\n", config.DIRECTORY.WallpaperPath);
  fmt.Println("");
  

  // Asks for custom directory if the path is default in the config file
  if config.DIRECTORY.WallpaperPath == "Enter wallpapers directory path" {
    Clear_screen();
    CheckUserWallDir();
    config.DIRECTORY.WallpaperPath = UserInputDir; 
    os.Remove(configPath)
    startup.CheckConfigFile(config.DIRECTORY.WallpaperPath)
  }
}
