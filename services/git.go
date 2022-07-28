package services

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
)

// function to clone exploitdb repo to current working directory
func git_clone_exploitdb(git_dir string) {

	fmt.Printf("[%s][%s] Cloning exploitdb repository\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.BlueString("INIT"))
	r, err := git.PlainClone(git_dir, false, &git.CloneOptions{
		URL: "https://github.com/offensive-security/exploitdb",
	})

	if err != nil {
		log.Error(err)
	}
	log.Info(r)
	fmt.Printf("[%s][%s] Cloned exploitdb repository\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))

}

// function to pull latest changes into policy-template folder
func git_pull_exploitdb(git_dir string) {

	fmt.Printf("[%s][%s] Fetching updates from exploitdb repository\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.BlueString("INIT"))
	r, err := git.PlainOpen(git_dir)
	if err != nil {
		log.Error(err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Error(err)
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		log.Debug(err)
	}

	fmt.Printf("[%s][%s] Fetched updates from exploitdb repository\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))

}

// Function to Create connection to kubernetes cluster

func Git_Operation(git_dir string) {

	//check if the policy-template directory exist
	// if exist pull down the latest changes
	// else clone the policy-templates repo
	if _, err := os.Stat(git_dir); !os.IsNotExist(err) {

		git_pull_exploitdb(git_dir)

	} else {

		git_clone_exploitdb(git_dir)

	}

}
