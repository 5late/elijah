package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

const chunkSize = 64000

var path = "/home/all/repos/elijah/"
var filename = "../push/example-file.txt"
var comparer = "../push/example-file2.txt"

func mainCheck() {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	secondData, err := ioutil.ReadFile(comparer)
	if err != nil {
		log.Fatal(err)
	}
	print(string(secondData))
	print(string(data))

	if string(secondData) == string(data) {
		print("same")
	}
}

func Copy(src, target string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func checkForChanges() bool {
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Fetch(&git.FetchOptions{
		RemoteName: "origin",
	})
	return err != nil
}

func main() {
	if checkForChanges() {
		mainCheck()
		return
	}
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return
	}

	// Print the latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(commit)
	Copy(filename, comparer)
}
