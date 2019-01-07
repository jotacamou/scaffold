package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Project struct {
	Name       string
	Owner      string
	Program    string
	Maintainer string
	// TODO: implement extra variables
	ExtraVars map[string]interface{}
}

var profile, name string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:    "create NAME [flags]",
	Short:  "Create a new development project from template directory",
	Args:   cobra.MinimumNArgs(1),
	PreRun: preRun,
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgHiGreen).SprintFunc()
		fmt.Printf("Creating project from %s template directory...\n", green(profile))
		err := createProject(NewProject())
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&profile, "profile", "p", "", "New project's profile")
}

func preRun(cmd *cobra.Command, args []string) {
	name = args[0]
	validName := "[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]+"
	if ok, _ := regexp.MatchString(validName, name); !ok {
		fmt.Printf("%s: invalid name '%s': should be in owner/program format\n", cmd.Name(), name)
		os.Exit(1)
	}

	if profile == "" {
		fmt.Printf("%s: you must provide a template profile (--profile)\n", cmd.Name())
		os.Exit(1)
	}

	profiles, err := profileList()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ok := false

	for _, templ := range profiles {
		if templ == profile {
			ok = true
			break
		}
	}

	if !ok {
		fmt.Printf("%s: could not find template directory '%s'\n\n", cmd.Name(), profile)
		fmt.Println("Available templates in config:")
		for _, templ := range profiles {
			fmt.Printf("  %s\n", templ)
		}
		os.Exit(1)
	}

}
func NewProject() *Project {
	var username string
	u, err := user.Current()
	if err != nil {
		username = "unknown"
	} else {
		username = u.Username
	}

	items := strings.Split(name, "/")

	return &Project{
		Name:       name,
		Owner:      items[0],
		Program:    items[len(items)-1],
		Maintainer: username,
	}
}

func createProject(p *Project) error {
	templPath := fmt.Sprintf("%s/%s", viper.GetString("templatesDir"), profile)
	if _, err := os.Stat(p.Program); err == nil {
		err := fmt.Sprintf("project '%s' already exists", p.Program)
		return errors.New(err)
	}

	tmpDir, err := ioutil.TempDir(".", fmt.Sprintf(".%s", rootCmd.Name()))
	if err != nil {
		panic(err)
	}

	err = filepath.Walk(templPath, func(path string, info os.FileInfo, err error) error {
		// FIXME: recursive walk to allow directories
		if info.IsDir() {
			return nil
		}

		fmt.Println("processing", path)
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return nil
		}

		f, err := os.Create(filepath.Join(tmpDir, info.Name()))
		defer f.Close()
		if err != nil {
			fmt.Println(err)
			return nil
		}

		err = tmpl.ExecuteTemplate(f, info.Name(), p)
		if err != nil {
			return err
		}

		return nil
	})

	err = os.Rename(tmpDir, p.Program)
	if err != nil {
		return err
	}

	return nil
}

func profileList() ([]string, error) {
	var profiles []string

	files, err := ioutil.ReadDir(viper.GetString("templatesDir"))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if !f.IsDir() {
			fmt.Printf("ignoring directory '%s'\n", f.Name())
			continue
		}
		profiles = append(profiles, f.Name())
	}

	return profiles, nil
}
