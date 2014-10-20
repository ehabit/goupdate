# goupdate will update all of your Go packages

This is a simple tool for updating all of your Go packages. 

goupdate grabs your GOPATH from your environment and then runs through your library of local packages hosted on github.com, bitbucket.org, code.google.com, and gopkg.in and tries to update them by executing "go get -u" on all of the packages. 

## Usage
````bash
  go get github.com/ehabit/goupdate
  goupdate
````