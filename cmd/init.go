package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new TIL repo",
	Long:  "init is used to set up your til repository",
	Run: func(cmd *cobra.Command, args []string) {
		runInit(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	validateDir := func(input string) error {
		return os.Mkdir(input, os.ModePerm)
	}

	prompt := promptui.Prompt{
		Label:    "Where would you like to save your TILs❓",
		Validate: validateDir,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error creating new directory %v", err)
		return
	}

	fmt.Printf("📓 Your TILs will be curated here: %v", result)
}
