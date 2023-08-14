package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/pmezard/go-difflib/difflib"
)

// View Command
var viewCmd = &cobra.Command{
	Use: "view",
	Short: "View the kubeconfig file",
	Example: "kubetune view",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory", err)
			return
		}
		filePath := filepath.Join(homeDir, ".kube", "config") // Construct the path to the kubeconfig file
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading kubeconfig file:", err)
			return
		}
		fmt.Println(string(data)) // Prints the content of the file
	},
}

// Change Command
var changeCmd = &cobra.Command{
	Use: "change",
	Short: "Change the kubeconfig file",
	Example: "kubetune change",
	Run: func(cmd *cobra.Command, args []string) {
		// Get home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory", err)
			return
		}
		filePath := filepath.Join(homeDir, ".kube", "config")
		backupPath := filePath + ".backup"

		// Read the current kubeconfig file
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading kubeconfig file:", err)
			return
		}

		// Save a copy of the current kubeconfig file as a backup
		err = os.WriteFile(backupPath, data, 0644)
		if err != nil {
			fmt.Println("Error creating backup file:", err)
			return
		}

		// Select Editor
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vi" // Default to vi if no editor is set
		}

		command := exec.Command(editor, filePath)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err = command.Run()
		if err != nil {
			fmt.Println("Error opening editor:", err) 
			return
		}

		fmt.Println("Changes saved to kubeconfig file.")
	},
}

// Diff Command
var diffCmd = &cobra.Command{
	Use: "diff",
	Short: "Track changes in a diff format for the kubeconfig file",
	Example: "kubetune diff",
	Run: func(cmd *cobra.Command, args []string) {
		// Get home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory", err)
			return
		}
	
		// Set paths
		filePath := filepath.Join(homeDir, ".kube", "config")
		backupPath := filePath + ".backup"

		// Get current kubeconfig file
		current, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading kubeconfig file:", err)
			return
		}

		// Create backup kubeconfig 
		backup, err := os.ReadFile(backupPath)
		if err != nil {
			fmt.Println("Error reading backup file:", err)
			return
		}

		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(string(backup)),
			B:        difflib.SplitLines(string(current)),
			FromFile: "Previous",
			ToFile:   "Current",
			Context:  3,
		}

		diffText, err := difflib.GetUnifiedDiffString(diff)
		if err != nil {
			fmt.Println("Error generating diff:", err)
			return
		}

		fmt.Println(diffText)
	},
}

// Main Function
func main() {
	var rootCmd = &cobra.Command{Use: "kubetune"}

	rootCmd.AddCommand(viewCmd, changeCmd, diffCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}