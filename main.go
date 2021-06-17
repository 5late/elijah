package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/joho/godotenv"
)

const chunkSize = 64000

var filename = "example-file.txt"
var comparer = "example-file2.txt"

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

func deepCompare(file1, file2 string) bool {
	// Check file size ...

	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				log.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}

func main() {
	if deepCompare(filename, comparer) {
		return
	}
	godotenv.Load(".env")

	username := os.Getenv("username")
	password := os.Getenv("password")
	var auth = &http.BasicAuth{
		Username: username,
		Password: password,
	}

	path := "/home/all/repos/elijah/"

	remotename := "origin"
	commitmessage := (filename + " | file auto-committed and pushed to GitHub ~5late")

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

	// push using default options (auth and remote name defined at top of file)

	err = r.Push(&git.PushOptions{
		RemoteName: remotename,
		Auth:       auth,
	})
	if err != nil {
		log.Fatal(err)
	}
	Copy(filename, comparer)
}
