package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var learnCmd = &cobra.Command{
	Use:   "learned [what you learned]",
	Short: "Write what you learned",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		learned()
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)
}

func learned() {
	fmt.Print("Successfully learned ")
	d := color.New(color.FgGreen, color.Bold)
	d.Printf(" ✔\n")
}
