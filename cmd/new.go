package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

func init() {
	rootCmd.AddCommand(newCmd)
}

func runNewCmd(cobraCmd *cobra.Command, args []string) {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt Failed %v", err)
	}

	fmt.Printf("You chose %v", result)

	cmd := exec.Command("nvim", "~/testfiletestfile.md")
	cmd.Stdout, cmd.Stderr, cmd.Stdin = os.Stdout, os.Stderr, os.Stdin
	if err := cmd.Run(); err != nil {
		// fmt.De
	}
}
