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

var taskName bool
var taskCompleted bool
var listFlag bool

// todolistCmd represents the todolist command
var todolistCmd = &cobra.Command{
	Use:   "todo [task name]",
	Short: "Use Todo [name] to register a task, to complete a task add --done at end",
	Long: `
	Registering a task:
	todo [task name]
	Completing a task:
	todo [task name] --done
	todo [task name] -d
	`,
	Run: func(cmd *cobra.Command, args []string) {

		filename := "tasks.json"

		list, errRead := app.ReadTasksFromFile(filename)
		if errRead != nil {
			fmt.Println(errRead)
		}
		for _, task := range list {
			fmt.Printf("Task: %s | Status: %t\n", task.Name, task.Status)
		}
		if len(args) > 0 && !listFlag {

			if len(args) < 1 {
				fmt.Println("argumento [task name] is required")
				return
			}
			err := app.AddOrUpdateTask(filename, args[0], taskCompleted)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Task added or updated successfully!")
		}

	},
}

func init() {
	rootCmd.AddCommand(todolistCmd)

	todolistCmd.Flags().BoolVarP(&taskCompleted, "done", "d", false, "Complete a task")
	todolistCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "List Tasks")
}
