package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/joho/godotenv"
)

// Example of how to open a repository in a specific path, and push to
// its default remote (origin).
func main() {
	godotenv.Load(".env")

	path := "/home/all/repos/elijah/"
	username := os.Getenv("username")
	fmt.Println(username)
	password := os.Getenv("password")
	var auth = &http.BasicAuth{
		Username: username,
		Password: password,
	}

	// Opens an already existing repository.
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	// ... we need a file to commit so let's create a new file inside of the
	// worktree of the project using the go standard library.

	filename := filepath.Join(path, "example-git-file")
	err = ioutil.WriteFile(filename, []byte("hello world!"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Adds the new file to the staging area.

	_, err = w.Add("example-git-file")
	if err != nil {
		log.Fatal(err)
	}

	// We can verify the current status of the worktree using the method Status.

	status, err := w.Status()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit Since version 5.0.1, we can omit the Author signature, being read
	// from the git config files.

	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "5|ate",
			Email: "christianrfernandes5@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Prints the current HEAD to verify that all worked well.

	obj, err := r.CommitObject(commit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(obj)

	// push using default options
	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	if err != nil {
		log.Fatal(err)
	}
}
