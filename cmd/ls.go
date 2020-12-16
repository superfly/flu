package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
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
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Directory Name", "Appname"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	err := scanapps(table, "./")

	if err != nil {
		return err
	}

	table.Render()

	return nil
}

func scanapps(table *tablewriter.Table, dirname string) error {
	return scanappsprefixed("", table, dirname)
}

func scanappsprefixed(prefix string, table *tablewriter.Table, dirname string) error {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			tomlName := filepath.Join(dirname, f.Name(), "fly.toml")
			if fileExists(tomlName) {
				flyToml, err := toml.LoadFile(tomlName)
				if err != nil {
					return err
				} else {
					appNameHolder := flyToml.Get("app")
					if appNameHolder == nil {
						table.Append([]string{prefix + f.Name(), "Bad fly.toml (no app=)"})
					} else {
						appName := appNameHolder.(string)
						table.Append([]string{prefix + f.Name(), appName})
					}
				}
			} else {
				err = scanappsprefixed(prefix+f.Name()+"/", table, filepath.Join(dirname, f.Name()))
				if err != nil {
					return err
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
