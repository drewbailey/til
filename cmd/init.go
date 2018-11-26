package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
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
	// initialize a new repository in a location passed in via args

	dir, err := ioutil.TempDir("", "til")
	fs := osfs.New(dir)
	dot, _ := fs.Chroot(".git")

	st := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())
	repo, err := git.Init(st, nil)
	// repo
	// st := storage.Storer{}
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Printf("Err %v", err)
	}

}
