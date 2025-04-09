package cmd

import (
	"fmt"
	"github.com/gingerlavender/GitHub-User-Activity-Fetcher/internal/events"
	"github.com/gingerlavender/GitHub-User-Activity-Fetcher/internal/output"
	"github.com/spf13/cobra"
	"time"
)

var period time.Duration
var days int
var eventType string
var token string

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
		drawPlot, _ := cmd.Flags().GetBool("plot")
		events, err := events.FetchEvents(args[0], token)
		if err != nil {
			fmt.Printf("Error fetching events: %s\n", err)
			return
		}
		if err = output.PrintEvents(&events, period, eventType); err != nil {
			fmt.Printf("Error printing events: %s\n", err)
			return
		}
		if drawPlot {
			if err = output.DrawEventsPlot(&events, period, eventType); err != nil {
				fmt.Printf("Error drawing plot: %s\n", err)
				return
			}
			fmt.Printf("\nPlot has been successfully saved!\n")
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("day", "d", false, "Period: day")
	rootCmd.PersistentFlags().BoolP("week", "w", false, "Period: week")
	rootCmd.PersistentFlags().BoolP("month", "m", false, "Period: month")
	rootCmd.PersistentFlags().BoolP("year", "y", false, "Period: month")
	rootCmd.PersistentFlags().Bool("plot", false, "Draw a plot of user activity (in HTML format)")
	rootCmd.PersistentFlags().IntVar(&days, "period", 0, "Period: number of days")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Your API authorization token")
	rootCmd.PersistentFlags().StringVar(&eventType, "eventType", "", "Show only specific event (full name like \"PushEvent\")")
	rootCmd.MarkFlagsOneRequired("day", "week", "month", "year", "period")
	rootCmd.MarkFlagsMutuallyExclusive("day", "week", "month", "year", "period")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
