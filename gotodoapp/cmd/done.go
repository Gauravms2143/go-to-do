/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/Gauravms2143/go-to-do/gotodoapp/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark item as done",
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	//Read data from file.
	items, err := todo.ReadItems(dataFile)

	//we will set todo status based on label so convert string
	// label to integer.
	i, err := strconv.Atoi(args[0])

	//for invalid input genertae error.
	if err != nil {
		log.Fatalln(args[0], "is not valid label.\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v \n", items[i-1].Text, "marked done")

		sort.Sort(todo.ByPri(items))
		todo.SavedItems(dataFile, items)
	} else {
		log.Println("doesn't match any items.")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
