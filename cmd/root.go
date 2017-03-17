package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "githubin",
	Short: "Easily download binaries from Github",
	Long: `Github lets you store your binaries on their website
(more precisely: on S3). This is a cool way of software distribution.

From the consumer point of view it's not that easy though:
everyone packages software differently.
This is where GithuBin helps: it unifies the way you download binaries`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}