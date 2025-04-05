package cmd

import (
	"fmt"
	"gh-api/internal/events"
	"gh-api/internal/output"
	"github.com/spf13/cobra"
	"time"
)

var period time.Duration
var days int

var rootCmd = &cobra.Command{
	Use:   "gh-api [username]",
	Short: "Tool to check latest activity of some user",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		day, _ := cmd.Flags().GetBool("day")
		week, _ := cmd.Flags().GetBool("week")
		month, _ := cmd.Flags().GetBool("month")
		year, _ := cmd.Flags().GetBool("year")
		switch {
		case day:
			period = 24 * time.Hour
		case week:
			period = 7 * 24 * time.Hour
		case month:
			period = 30 * 24 * time.Hour
		case year:
			period = 365 * 24 * time.Hour
		default:
			period = time.Duration(days) * 24 * time.Hour
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		events, err := events.FetchEvents(args[0])
		if err != nil {
			fmt.Printf("Error fetching events: %s\n", err)
			return
		}
		if err = output.PrintEvents(events, period); err != nil {
			fmt.Printf("Error printing events: %s\n", err)
			return
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("day", "d", false, "Period: day")
	rootCmd.PersistentFlags().BoolP("week", "w", false, "Period: week")
	rootCmd.PersistentFlags().BoolP("month", "m", false, "Period: month")
	rootCmd.PersistentFlags().BoolP("year", "y", false, "Period: month")
	rootCmd.PersistentFlags().IntVar(&days, "period", 0, "Period: number of days")
	rootCmd.MarkFlagsOneRequired("day", "week", "month", "year", "period")
	rootCmd.MarkFlagsMutuallyExclusive("day", "week", "month", "year", "period")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
