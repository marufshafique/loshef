package cmd

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the loshef programm by passing file names :)",
	Long: `Start command will start finding difference between files:

This command will take two file names as argument. Please pass
primary file name first and sceondary file name as sceond argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		pryName := args[0]
		sdyName := args[1]
		// open json file
		pmr, err := os.Open(pryName)
		scr, serr := os.Open(sdyName)

		if err != nil {
			log.Fatal("File not found :(")
		}
		defer pmr.Close()
		if serr != nil {
			log.Fatal("File not found :(")
		}
		defer scr.Close()

		pmrValue, _ := ioutil.ReadAll(pmr)
		scrValue, _ := ioutil.ReadAll(scr)

		var result map[string]string
		var scrResult map[string]string
		diffs := make(map[string]string)
		json.Unmarshal([]byte(pmrValue), &result)
		json.Unmarshal([]byte(scrValue), &scrResult)

		for rk := range result {
			_, ok := scrResult[rk]
			if !ok {
				scrResult[rk] = result[rk]
				diffs[rk] = result[rk]
			}
		}
		file, _ := json.MarshalIndent(scrResult, "", "  ")
		diffFile, _ := json.MarshalIndent(diffs, "", "  ")
		ioutil.WriteFile("results.json", file, 0466)
		ioutil.WriteFile("diffs.json", diffFile, 0466)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
