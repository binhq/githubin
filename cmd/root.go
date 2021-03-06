package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logger = logrus.New()

var (
	version bool
	quiet   bool
)

var rootCmd = &cobra.Command{
	Use:   "gitbin",
	Short: "Easily download binaries from Github",
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Printf("GithuBin version %s, build %s (at %s)\n", Version, CommitHash, BuildDate)
			return
		}

		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if quiet {
			logger.Level = logrus.ErrorLevel
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Print version information and quit")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Do not show log output")
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
