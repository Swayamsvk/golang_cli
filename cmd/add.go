package cmd

import (
	"fmt"
	"strings"

	"github.com/Swayamsvk/task/task-modules/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task on your list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("Added \"%s\"to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
