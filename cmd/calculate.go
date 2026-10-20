package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCalculateCmd())
}

func newCalculateCmd() *cobra.Command {
	var a, b float64

	calCmd := &cobra.Command{
		Use:   "calculate",
		Short: "Perform mathematical operations",
		Example: `  tw calculate add -a 5 -b 3
  tw calculate divide -a 10 -b 4`,
	}

	calCmd.PersistentFlags().Float64VarP(&a, "a", "a", 0, "First number")
	calCmd.PersistentFlags().Float64VarP(&b, "b", "b", 0, "Second number")
	_ = calCmd.MarkPersistentFlagRequired("a")
	_ = calCmd.MarkPersistentFlagRequired("b")

	calCmd.AddCommand(
		newOperationCmd("add", "Add two numbers", func() (float64, error) {
			return a + b, nil
		}),
		newOperationCmd("subtract", "Subtract b from a", func() (float64, error) {
			return a - b, nil
		}),
		newOperationCmd("multiply", "Multiply two numbers", func() (float64, error) {
			return a * b, nil
		}),
		newOperationCmd("divide", "Divide a by b", func() (float64, error) {
			if b == 0 {
				return 0, errors.New("cannot divide by zero")
			}
			return a / b, nil
		}),
	)

	return calCmd
}

func newOperationCmd(use, short string, op func() (float64, error)) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := op()
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}
}
