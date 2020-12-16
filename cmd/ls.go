package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all the apps in current directory",
	Long:  `All software has versions. This is Hugo's`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := appscan(); err != nil {
			return err
		}
		return nil
	},
}

func appscan() error {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			tomlName := f.Name() + "/fly.toml"
			if fileExists(tomlName) {
				flyToml, err := toml.LoadFile(tomlName)
				if err != nil {
					return err
				} else {
					appNameHolder := flyToml.Get("app")
					if appNameHolder == nil {
						fmt.Printf("%20s %32s\n", f.Name(), "Bad fly.toml (no app=)")
					} else {
						appName := appNameHolder.(string)
						fmt.Printf("%20s %32s\n", f.Name(), appName)
					}
				}
			}
		}
	}

	return nil
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
