package main

import (
	"fmt"
	"log"
	"os"
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

	username := os.Getenv("username")
	password := os.Getenv("password")
	var auth = &http.BasicAuth{
		Username: username,
		Password: password,
	}

	path := "/home/all/repos/elijah/"
	filename := "main.go"
	remotename := "origin"
	commitmessage := ("%v file auto-committed and pushed to GitHub" + filename)

	// Opens existing repository
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	// Adds the file to the staging area

	_, err = w.Add(filename)
	if err != nil {
		log.Fatal(err)
	}

	// We can verify the current status of the worktree using the method Status.

	status, err := w.Status()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(status)

	//Commit the file

	commit, err := w.Commit(commitmessage, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "5|ate",
			Email: "christianrfernandes5@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Prints the current HEAD to verify that all worked well

	obj, err := r.CommitObject(commit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(obj)

	// push using default options (auth and remote name)

	err = r.Push(&git.PushOptions{
		RemoteName: remotename,
		Auth:       auth,
	})
	if err != nil {
		log.Fatal(err)
	}
}
