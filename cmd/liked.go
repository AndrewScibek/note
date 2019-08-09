package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var likedCmd = &cobra.Command{
	Use:   "liked [what you liked]",
	Short: "Write what you liked",
	Run: func(cmd *cobra.Command, args []string) {
		liked()
	},
}

func init() {
	rootCmd.AddCommand(likedCmd)
}

func liked() {
	fmt.Print("Successfully liked ")
	d := color.New(color.FgGreen, color.Bold)
	d.Printf("✔\n")
}
