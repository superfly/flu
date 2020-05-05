package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	toml "github.com/pelletier/go-toml"
)


func main() {
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if f.IsDir() {
				tomlName:=f.Name()+"/fly.toml"
				if fileExists(tomlName) {
					flyToml,err:=toml.LoadFile(tomlName)
					if err != nil {
						fmt.Println("Error ", err.Error())
					} else {
						appNameHolder:=flyToml.Get("app")
						if appNameHolder==nil {
							fmt.Printf("%20s %32s\n",f.Name(),"Bad fly.toml (no app=)")
						} else {
							appName := appNameHolder.(string)
							fmt.Printf("%20s %32s\n", f.Name(), appName)
						}
					}
				}
			}
		}
	}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}