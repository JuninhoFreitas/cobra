/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"

	"github.com/JuninhoFreitas/cobra/app"
	"github.com/spf13/cobra"
)

var a, b float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Cli controller of Calculator",
	Long:  `Cli controller of Calculator`,
	Run: func(cmd *cobra.Command, args []string) {
		calc := app.NewCalc()
		calc.A = a
		calc.B = b
		fmt.Println(calc.Sum())
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().Float64VarP(&a, "a", "a", 0, "A value")
	cliCmd.Flags().Float64VarP(&b, "b", "b", 0, "B value")
}
