// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/juju/errors"

	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

// lexiconsCmd represents the lexicons command
var lexiconsCmd = &cobra.Command{
	Use:   "lexicons",
	Short: "List lexicons",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := listLexicons(cmd, args); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	listCmd.AddCommand(lexiconsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lexiconsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lexiconsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	lexiconsCmd.Flags().StringP("owner", "o", "", "List all lexicons of the given owner")
	lexiconsCmd.Flags().StringP("type", "t", "", "List all lexicons of the given lexType")
	lexiconsCmd.Flags().StringP("name", "n", "", "Discribe the named lexicon")
	lexiconsCmd.Flags().BoolP("intalled", "i", false, "List Lexicons used in current project")

}

func listLexicons(cmd *cobra.Command, args []string) error {

	owner := cmd.Flag("owner").Value.String()
	lexType := cmd.Flag("type").Value.String()
	name := cmd.Flag("name").Value.String()

	url := "https://raw.githubusercontent.com/codelingo/hub/master/lexicons/lingo_lexicon_type.yaml"
	switch {
	case name != "":

		if owner == "" {
			return errors.New("owner flag must be set")
		}

		if lexType == "" {
			return errors.New("lexType flag must be set")
		}
		url = fmt.Sprintf("https://raw.githubusercontent.com/codelingo/hub/master/lexicons/%s/%s/%s/lingo_lexicon.yaml",
			lexType, owner, name)

	case owner != "":
		if lexType == "" {
			return errors.New("lexType flag must be set")
		}
		url = fmt.Sprintf("https://raw.githubusercontent.com/codelingo/hub/master/lexicons/%s/%s/lingo_owner.yaml",
			lexType, owner)
	case lexType != "":

		url = fmt.Sprintf("https://raw.githubusercontent.com/codelingo/hub/master/lexicons/%s/lingo_lexicons.yaml",
			lexType)
	}
	resp, err := http.Get(url)
	if err != nil {
		return errors.Trace(err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Trace(err)
	}

	fmt.Println(string(data))
	return nil
}
