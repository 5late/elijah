# Elijah

**Automate pushing files to github in Go**

## How to use

- Create a ``.env`` file with a ``username`` and ``password`` field. 
    - The ``username`` field contains the username you wish to use, while the ``password`` field contains a personal access token.
- Change the ``path`` variable to match your github repo path
- Change the filename variable to the file you would like to push to github
- ***Profit***

### Uses [go-git](https://github.com/go-git/go-git)

**A lot of the code can be found in the [_examples](https://github.com/go-git/go-git/tree/master/_examples) folder**

- I have added auth and got it to properly work for me, as well as removed some of the functions to make it simpler.

## License

Licensed under Apache 2.0. Visit the [LICENSE] file to learn more.
