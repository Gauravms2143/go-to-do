/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/Gauravms2143/go-to-do/gotodoapp/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos.",
	Long:  `Listing the todos.`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	//fmt.Println("list called")
	//dataFile is flag, This is custome flag which we have created in
	//application(for more detail see root.go file).
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}
	//Sort Todo.
	sort.Sort(todo.ByPri(items))

	//We'll use tab writer to formate todos.
	//creates a new tabwriter with(Parameter Explanation)-
	//3: 			A minimum cell width of 3,
	//0: 			No padding,
	//1: 			One cell of padding between columns,
	//' ': 			A space character as the padding character,
	//0: 			No additional formatting flags.
	//os.Stdout: 	The tabwriter is created to write to the standard output.
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	//This loop iterates through the list of to-do items, and
	//writes each item's priority and text to the tabwriter w.
	//The Fprintln function is used to write the formatted text to
	//the tabwriter, followed by a newline character ("\n").
	for _, todo := range items {
		if todo.Done == doneOpt {
			fmt.Fprintln(w, todo.Label()+"\t"+todo.Prettydone()+"\t"+todo.PrettyP()+"\t"+todo.Text+"\t")
		}
	}

	// flushes any buffered data from the tabwriter w, ensuring
	// that all output has been written to the standard output.
	w.Flush()

}

var (
	doneOpt bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "shown 'Done' Todos.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
