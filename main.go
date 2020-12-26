package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/toricor/go-mini-git/repository"

	"github.com/spf13/cobra"
)

type ObjectType int

const (
	Unknown ObjectType = iota
	Blob
	Commit
	Tree
	Tag
)

var objectTypeLabelMap = map[string]ObjectType{
	"blob":   Blob,
	"commit": Commit,
	"tree":   Tree,
	"tag":    Tag,
}

var objectTypeMap = map[ObjectType]string{
	Blob:   "blob",
	Commit: "commit",
	Tree:   "tree",
	Tag:    "tag",
}

func (t ObjectType) String() string {
	if o, ok := objectTypeMap[t]; ok {
		return o
	}
	return "Unknown ObjectType"
}

type GitObject struct {
	objectType ObjectType
	size       uint64
	content    string
}

func buildGitObject(uncompressed string) GitObject {
	splited := strings.SplitN(uncompressed, "\x00", 2)
	header := splited[0]
	content := splited[1]
	objectType := strings.SplitN(header, " ", 2)[0]
	size, _ := strconv.ParseUint(strings.SplitN(header, " ", 2)[1], 10, 64)

	return GitObject{
		objectType: objectTypeLabelMap[objectType],
		size:       size,
		content:    content,
	}
}

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
		gitObject := buildGitObject(uncompressed)

		if catFileCmdFlagT == true {
			fmt.Println(gitObject.objectType)
		} else if catFileCmdFlagP == true {
			fmt.Println(gitObject.content)
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
