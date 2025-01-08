package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"math/rand"

	"github.com/spf13/cobra"
)

var (
	maxCommits int
	frequency  int
	remoteRepo string
)

var highActivityCmd = &cobra.Command{
	Use:   "high-activity",
	Short: "Generate high activity",
	Long:  "Generate high activity in your github",
	Run:   run,
}

func init() {
	highActivityCmd.Flags().IntVar(&maxCommits, "max-commits", 20, "Maximum number of commits per day")
	highActivityCmd.Flags().IntVar(&frequency, "frequency", 1, "Frequency of commits")
	highActivityCmd.Flags().StringVar(&remoteRepo, "remote", "", "Remote repository URL")
}

func run(cmd *cobra.Command, args []string) {
	if remoteRepo != "" {
		runCommand("git remote add origin " + remoteRepo)
	}

	today := time.Now()
	for i := 0; i < frequency; i++ {
		numCommits := contributionsPerDay(maxCommits)
		for j := 0; j < numCommits; j++ {
			date := today.Add(time.Duration(j) * time.Minute)
			contribute(date)
		}
	}
}

func contribute(date time.Time) {
	file, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	message := fmt.Sprintf("Contribution: %s\n", date.Format("2006-01-02 15:04"))
	if _, err := file.WriteString(message); err != nil {
		fmt.Println(err)
		return
	}

	runCommand("git add README.md")
	runCommand(fmt.Sprintf("git commit --date=\"%s\" -m \"%s\"", date.Format("2006-01-02 15:04"), message))
}

func runCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func contributionsPerDay(maxCommits int) int {
	if maxCommits < 1 {
		return 1
	}
	if maxCommits > 20 {
		return 20
	}
	return rand.Intn(maxCommits) + 1
}
