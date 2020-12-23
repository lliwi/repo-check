package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var frepo []string

func main() {

	if len(os.Args) > 3 {
		printHelp()

	} else {
		if os.Args[1] == "-u" {
			fmt.Println("Start checking repos on url" + os.Args[2])
			url := os.Args[2]
			if strings.HasPrefix(url, "http") == false {
				checkhttp("http://" + url)
				checkhttp("https://" + url)
			} else {
				checkhttp(url)
			}

		} else if os.Args[1] == "-f" {
			fmt.Println("Start checking repos from file")
			file := os.Args[2]
			readFile(file)

		} else {
			printHelp()

		}
	}

}

func printHelp() {

	colorBlue := "\033[34m"
	colorReset := "\033[0m"

	fmt.Println(string(colorBlue), "Use: ")
	fmt.Println(string(colorBlue), "    repo_check -h           --> to view this help")
	fmt.Println(string(colorBlue), "    repo_check -u  url      --> check single url")
	fmt.Println(string(colorBlue), "    repo_check -f filepath  --> check multiple urls from file", string(colorReset))

}

func readFile(file string) {
	colorGreen := "\033[32m"
	colorReset := "\033[0m"

	readFile, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	for _, eachline := range fileTextLines {

		if strings.HasPrefix(eachline, "http") == false {
			checkhttp("http://" + eachline)
			checkhttp("https://" + eachline)
		} else {
			checkhttp(eachline)
		}
	}
	fmt.Println(string(colorGreen), "Repos encontrados:")
	for _, repo := range frepo {
		if repo != "" {
			fmt.Println(repo)
		}
	}
	fmt.Println("All done :).", string(colorReset))

}

func checkhttp(url string) {

	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"

	var repos = []string{"/.git/", "/.snv/", "/CVS/", "/.bzr/", "/.hg/"}
	found := 0
	urlfound := ""
	url2f := ""

	isUrl, err := regexp.MatchString(`^(https?://)?([-a-z0-9]+)[.]([-a-z0-9]+)?(.([-a-z0-9]+))`, url)

	if err != nil {
		log.Fatal(err)
	}

	if isUrl != false {

		fmt.Println(string(colorYellow), "*** Start cheching "+url+" ***", string(colorReset))
		for _, repo := range repos {
			fmt.Println("[i] " + repo + "for " + url)
			if webIsReachable(url+repo) == true {

				found = found + 1
				urlfound = "[!] " + url + repo + " repo found"
				url2f = "[!] " + url + repo
			}
		}
		if found < 2 {
			fmt.Println(string(colorGreen), urlfound, string(colorReset))
			frepo = append(frepo, url2f)
		}

	}
}

func webIsReachable(web string) bool {
	response, errors := http.Get(web)

	if errors != nil {
		_, netErrors := http.Get(web)

		if netErrors != nil {

		}

		return false
	}

	if response.StatusCode == 200 {
		return true
	}

	return false

}
