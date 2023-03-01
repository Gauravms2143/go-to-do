/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/Gauravms2143/go-to-do/gotodoapp/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo.",
	Long:  `Add will create new Todo item to our todo list.`,
	Run:   addRun,
}

/**
*addRun : This function will called when `app.exe add` command is fire.
*arg []string is slice which is array underlying representation, store all
* input arguments.
 */
func addRun(cmd *cobra.Command, arg []string) {
	//items := []todo.Item{}
	//Read data from file.
	//dataFile is flag, This is custome flag which we have created in
	//application(for more detail see root.go file).
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("Internal Error : %v", err)
		os.Exit(-1)
	}

	for _, value := range arg {
		items = append(items, todo.Item{Text: value})
	}
	//fmt.Printf("%#v\n", items)
	err1 := todo.SavedItems(dataFile, items)
	if err1 != nil {
		fmt.Errorf("%v", err1)
	}
}

/**
* init() is special function,its called after package variable declaration.
* each package may have multiple init().
* init() execution order is unguaranteed.
 */
func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
