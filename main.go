package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// docbaseコマンドの本体
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
		targetPath := filepath.Join(".git", "objects", sha1[0:2], sha1[2:])

		buff, err := ioutil.ReadFile(targetPath)
		if err != nil {
			log.Fatal(err)
		}

		b := bytes.NewReader(buff)
		r, err := zlib.NewReader(b)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		if catFileCmdFlagT == true {
			// https://stackoverflow.com/questions/9644139/from-io-reader-to-string-in-go
			buf := new(strings.Builder)
			n, err := io.Copy(buf, r)
			if err != nil {
				log.Fatal(err, n)
			}

			header := strings.SplitN(buf.String(), "\x00", 2)[0]
			objectType := strings.SplitN(header, " ", 2)[0]
			fmt.Println(objectType)
		} else if catFileCmdFlagP == true {
			io.Copy(os.Stdout, r)
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
