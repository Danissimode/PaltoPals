// cli.go: Command-line interface for Paltopals, including root command and flags

package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/Danissimode/Palto/config"
	"github.com/Danissimode/Palto/internal"
	"github.com/Danissimode/Palto/logger"
	"github.com/spf13/cobra"
)

var (
	initMessage  string
	taskFileFlag string
)

var rootCmd = &cobra.Command{
	Use:   "Paltopals [request message]",
	Short: "Paltopals - AI-Powered Palto Companion",
	Long:  `Paltopals - AI-Powered Palto Companion`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Printf("Paltopals version: %s\ncommit: %s\nbuild date: %s\n", internal.Version, internal.Commit, internal.Date)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			logger.Error("Error loading configuration: %v", err)
			fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
			os.Exit(1)
		}

		if len(args) > 0 {
			initMessage = strings.Join(args, " ")
		}

		if taskFileFlag != "" {
			content, err := os.ReadFile(taskFileFlag)
			if err != nil {
				logger.Error("Error reading task file: %v", err)
				fmt.Fprintf(os.Stderr, "Error reading task file: %v\n", err)
				os.Exit(1)
			}
			initMessage = string(content)
			logger.Info("Read request from file: %s", taskFileFlag)
		}

		mgr, err := internal.NewManager(cfg)
		if err != nil {
			logger.Error("manager.NewManager failed: %v", err)
			os.Exit(1)
		}
		if initMessage != "" {
			logger.Info("Starting with initial subcommand: %s", initMessage)
		}

		if err := mgr.Start(initMessage); err != nil {
			logger.Error("manager.Start failed: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&taskFileFlag, "file", "f", "", "Read request from specified file")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
}

func Execute() error {
	return rootCmd.Execute()
}
