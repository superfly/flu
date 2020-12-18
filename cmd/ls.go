package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

var recursive bool
var sortapps bool

func init() {
	lsCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recur into any non-dotted directory to search heirachies")
	lsCmd.Flags().BoolVarP(&sortapps, "sortapps", "s", false, "Sort results by app name")
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all the apps in the children of the current directory",
	Long: `List all the apps  in the children of the current directory.
	Scan all directories for fly.toml. Where it exists, parse it and report the app's name`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := appscan(); err != nil {
			return err
		}
		return nil
	},
}

func appscan() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Appname", "Directory Name"})
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

	appentries, err := scanapps(table, "./")

	if err != nil {
		return err
	}

	if sortapps {
		sort.Slice(appentries, func(i, j int) bool {
			return appentries[i].appname < appentries[j].appname
		})
	}

	for _, v := range appentries {
		table.Append([]string{v.appname, v.appdir})
	}
	table.Render()

	return nil
}

type AppEntry struct {
	appname string
	appdir  string
}

func scanapps(table *tablewriter.Table, dirname string) ([]AppEntry, error) {
	return scanappsprefixed("", table, dirname)
}

func scanappsprefixed(prefix string, table *tablewriter.Table, dirname string) ([]AppEntry, error) {
	appentries := make([]AppEntry, 0)

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return appentries, err
	}

	for _, f := range files {
		if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
			tomlName := filepath.Join(dirname, f.Name(), "fly.toml")
			if fileExists(tomlName) {
				flyToml, err := toml.LoadFile(tomlName)
				if err != nil {
					return appentries, err
				} else {
					appNameHolder := flyToml.Get("app")
					if appNameHolder == nil {
						appentries = append(appentries, AppEntry{"Bad fly.toml (no app=)", prefix + f.Name()})
					} else {
						appName := appNameHolder.(string)
						appentries = append(appentries, AppEntry{appName, prefix + f.Name()})
					}
				}
			} else {
				if recursive {
					subapps, err := scanappsprefixed(prefix+f.Name()+"/", table, filepath.Join(dirname, f.Name()))
					if err != nil {
						return appentries, err
					}
					appentries = append(appentries, subapps...)
				}
			}
		}
	}

	return appentries, nil
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
