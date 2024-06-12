package mytool

import (
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var applyCmd = &cobra.Command{
	Use:           "test",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "Subcommand description",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	// Sub command args
	var argOne string
	applyCmd.Flags().StringVar(&argOne, "ignore-state", "default-value", "Argument description")
}
