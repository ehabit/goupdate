// GoUpdate written by Glen Baker glen@ehab.it
// GoUpdate is protected under MIT License for Open Source Software.

package main

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"io/ioutil"
	"os"
	"sync"
)

var GOPATH = os.Getenv("GOPATH")

var GITHUBPATH = GOPATH + "/src/github.com"
var BITBUCKETPATH = GOPATH + "/src/bitbucket.org"
var GOOGLECODEPATH = GOPATH + "/src/code.google.com"
var GOPKGPATH = GOPATH + "/src/gopkg.in"

var HOSTS = []string{
	GITHUBPATH,
	BITBUCKETPATH,
	GOOGLECODEPATH,
	GOPKGPATH,
}

var packagesUpdated int
var problemPackages []string

func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			fmt.Println(err)
			return false
		}
	}

	if f.IsDir() {
		return true
	}
	return false
}

func ParseFilenames(path string) []string {
	var filenames []string
	contents, _ := ioutil.ReadDir(path)
	for _, f := range contents {
		filenames = append(filenames, f.Name())
	}
	return filenames
}

func IsGoFile(filename string) bool {
	if filename[len(filename)-3:] == ".go" {
		return true
	}
	return false
}

func CheckDirForGo(path string) bool {
	if IsDir(path) {
		for _, f := range ParseFilenames(path) {
			if IsGoFile(f) {
				return true
			}
		}
	}
	return false
}

func UpdatePackage(path string) {
	if IsDir(path) {
		err := sh.Command("go", "get", "-u", sh.Dir(path)).Run()
		if err != nil {
			problemPackages = append(problemPackages, path+" "+err.Error())
			fmt.Println("Package not updated:", path, err)
		} else {
			packagesUpdated += 1
			fmt.Println("Updated package:", path)
		}
	}
}

func UpdatePackages(hostPath string) {
	paths := make(chan string, 64)

	// spawn four workers to update packages
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for path := range paths {
				UpdatePackage(path)
			}
			wg.Done()
		}()
	}

	userPaths := ParseFilenames(hostPath)
	for _, user := range userPaths {
		// one directory deep
		packagePaths := ParseFilenames(hostPath + "/" + user)
		if hostPath == GOPKGPATH {
			path := hostPath + "/" + user
			if CheckDirForGo(path) {
				paths <- path
			}
		}
		for _, pack := range packagePaths {
			// two directories deep
			if hostPath == GOOGLECODEPATH {
				subpackPaths := ParseFilenames(hostPath + "/" + user + "/" + pack)
				for _, subpack := range subpackPaths {
					// three directories deep
					path := hostPath + "/" + user + "/" + pack + "/" + subpack
					if CheckDirForGo(path) {
						paths <- path
					} else {
						// four directories deep
						supSubPaths := ParseFilenames(path)
						for _, supSubPath := range supSubPaths {
							if CheckDirForGo(path + "/" + supSubPath) {
								paths <- supSubPath
							}
						}
					}
				}
			} else if hostPath == GITHUBPATH || hostPath == BITBUCKETPATH {
				path := hostPath + "/" + user + "/" + pack
				if CheckDirForGo(path) {
					paths <- path
				}
			}
		}
	}

	close(paths)
	wg.Wait()
}

func UpdatePackagesOnHosts(hosts []string) {
	fmt.Println("Updating Go packages hosted on github.com, bitbucket.org, code.google.com and gopkg.in")

	for _, host := range hosts {
		UpdatePackages(host)
	}
}

func UpdateCount() int {
	return packagesUpdated
}

func ReportUpdateStats() {
	fmt.Println("Total packages updated:", UpdateCount())
	fmt.Println("Total package update errors:", len(problemPackages))
	for i, pack := range problemPackages {
		fmt.Println(i+1, pack)
	}
}

func main() {
	UpdatePackagesOnHosts(HOSTS)

	fmt.Println("\nWe have updated every G-bomb package we could find.\nGood luck out there you sexy gopher!\n")

	ReportUpdateStats()
}
