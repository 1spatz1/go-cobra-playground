package stringer

import (
	"fmt"
	"github.com/1spatz1/go-cobra-playground/stringer/pkg/stringer"
	"github.com/spf13/cobra"
)

var primeCmd = &cobra.Command{
	Use:     "prime",
	Aliases: []string{},
	Short:   "Checks if the input is a prime",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args[0]
		res := stringer.IsPrime(i)
		if res {
			fmt.Printf("%s is a prime number", i)
		} else {
			fmt.Printf("%s is not a prime number", i)
		}
	},
}

func init() {
	rootCmd.AddCommand(primeCmd)
}
