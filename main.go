package main

import (
	"fmt"
	"os"

	"github.com/toricor/go-mini-git/object"
	"github.com/toricor/go-mini-git/repository"

	"github.com/spf13/cobra"
)

// Command
var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("the stupid content tracker")
	},
}

var catFileCmdFlagT bool
var catFileCmdFlagP bool
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "it's git cat-file ",
	Long:  "Provide content or type and size information for repository objects",
	Run: func(cmd *cobra.Command, args []string) {
		sha1 := args[0]
		targetPath := repository.GetGitObjectFilePath(sha1)
		uncompressed := repository.GetUncompressedContent(targetPath)
		gitObject := object.BuildGitObject(uncompressed)

		if catFileCmdFlagT == true {
			fmt.Println(gitObject.ObjectType)
		} else if catFileCmdFlagP == true {
			fmt.Println(gitObject.Content)
		} else {
			fmt.Println("unreachable")
		}
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	catFileCmd.Flags().BoolVarP(&catFileCmdFlagT, "type", "t", false, "show type: cat-file -t ${sha1_hash}")
	catFileCmd.Flags().BoolVarP(&catFileCmdFlagP, "pretty-print", "p", false, "pretty print: cat-file -p ${sha1_hash}")
	rootCmd.AddCommand(catFileCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
