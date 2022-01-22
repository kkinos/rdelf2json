package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kinpoko/rdelf2json/elftojson"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rdelf2json [file name]",
	Short: "Convert ELF headers to json",
	Long:  "Convert ELF headers to json",

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		json, err := elftojson.ELFToJson(b)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
		return nil

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
