# goupdate will update all of your Go packages

[![wercker status](https://app.wercker.com/status/ee5eb248e1a1fbc02c83626807e511e8/m "wercker status")](https://app.wercker.com/project/bykey/ee5eb248e1a1fbc02c83626807e511e8)

This is a simple tool for updating all of your Go packages. 

goupdate grabs your GOPATH from your environment and then runs through your library of local packages hosted on github.com, bitbucket.org, code.google.com, and gopkg.in and tries to update them by executing "go get -u" on all of the packages. 

## Usage
````bash
  go get github.com/ehabit/goupdate
  goupdate
````

## Features
+ Updates Go packages hosted on github.com, bitbucket.org, code.google.com, and gopkg.in.
+ Spins up 4 workers to parallelize package updates.