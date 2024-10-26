package gitclient

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type GitClient struct {
	Path string
}

func (gitRepo *GitClient) CheckIsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = gitRepo.Path
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func (gitRepo *GitClient) GetBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	branches := strings.Split(string(output), "\n")

	for i, branch := range branches {
		if strings.HasPrefix(branch, "*") {
			branches[i] = strings.Trim(branch, "*")
		}
		branches[i] = strings.TrimSpace(branches[i])
	}

	defaultBranch := getDefaultBranch()

	for i := len(branches) - 1; i >= 0; i-- {
		if branches[i] == "" || branches[i] == defaultBranch {
			branches = append(branches[:i], branches[i+1:]...)
		}
	}

	return branches, nil
}

func (client *GitClient) DeleteBranches(branches []string) (string, error) {
	if len(branches) == 0 {
		return "", errors.New("no branches selected to delete")
	}

	branchesStr := strings.Join(branches, " ")
	fmt.Println("Deleting branches: ", branchesStr)
	cmd := exec.Command("git", append([]string{"branch", "-D"}, branches...)...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Output: %s\n", string(output))
		return "", err
	}

	return strings.Join(branches, ", "), nil
}

func getDefaultBranch() string {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(output))
}
