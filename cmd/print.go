package cmd

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newPrintCmd())
}

func newPrintCmd() *cobra.Command {
	var message, color string

	printCmd := &cobra.Command{
		Use:   "print",
		Short: "Print a colored message",
		Example: `  tw print --message "Hello World" --color "#FF5733"
  					tw print -m "Hello World" -c "#00FF00"`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if message == "" {
				return errors.New("message cannot be empty")
			}

			style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
			fmt.Println(style.Render(message))
			return nil
		},
	}

	printCmd.Flags().StringVarP(&message, "message", "m", "", "Message to print")
	printCmd.Flags().StringVarP(&color, "color", "c", "10", "Text color")

	return printCmd
}
