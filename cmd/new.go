package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new til",
	Long:  "create a new til",
	Run: func(cmd *cobra.Command, args []string) {
		runNewCmd(cmd, args)
	},
}

type Til struct {
	Title string
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func runNewCmd(cobraCmd *cobra.Command, args []string) {
	spew.Dump(args)
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Today I learned ✍️",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt Failed %v", err)
	}

	fmt.Printf("You chose %v", result)

	tmplStruct := Til{Title: result}

	tmpl := template.Must(template.ParseFiles("template.md"))

	// temp file
	tempFile, err := os.Create("/tmp/dat2")
	defer tempFile.Close()
	if err != nil {
		fmt.Printf("Error")
	}

	tmpl.Execute(tempFile, tmplStruct)

	tempFile.Close()
	cmd := exec.Command("nvim", "/tmp/dat2")
	cmd.Stdout, cmd.Stderr, cmd.Stdin = os.Stdout, os.Stderr, os.Stdin
	if err := cmd.Run(); err != nil {
		// fmt.De
	}
}
