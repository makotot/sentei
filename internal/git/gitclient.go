package gitclient

import (
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

func (gitRepo *GitClient) GetBranches() []string {
	cmd := exec.Command("git", "branch", "--list")
	cmd.Dir = gitRepo.Path
	if err := cmd.Run(); err != nil {
		return nil
	}
	output, err := cmd.Output()
	if err != nil {
		return nil
	}
	return strings.Fields(string(output))
}
