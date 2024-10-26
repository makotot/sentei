/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	gitclient "github.com/makotot/sentei/internal/git"
	"github.com/makotot/sentei/internal/tui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "sentei",
	Short:   "A tool to interactively select branches to be deleted.",
	Version: "0.0.2",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		repo := gitclient.GitClient{Path: wd}
		isrepo := repo.CheckIsGitRepo()

		if !isrepo {
			fmt.Println("Not a git repository.")
			return
		}

		branches, err := repo.GetBranches()

		if err != nil {
			fmt.Println(err)
			return
		}

		selected, err := tui.Form(branches)

		if err != nil {
			fmt.Println(err)
			return
		}

		result, err := repo.DeleteBranches(selected)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Successfully delete these branches: ", result)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sentei.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
