package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Please give a parameter")
	}
	parameter1 := os.Args[1]
	if parameter1 == "-l" {
		toLinux(os.Args[2])
	} else if parameter1 == "-w" {
		toWin(os.Args[2])
	} else if parameter1 != "" {
		toLinux(parameter1)
	}
}

func toWin(path string) {
	linuxPath := path
	rootPath := ""
	lastPath := linuxPath
	reg := regexp.MustCompile(`^/mnt/[a-z]`)
	if reg.MatchString(linuxPath) {
		rootPath = fmt.Sprintf("%v:\\", strings.ToUpper(linuxPath[5:6]))
		lastPath = linuxPath[7:]
	}
	reg = regexp.MustCompile(`/`)
	lastPath = reg.ReplaceAllString(lastPath, "\\")
	fmt.Println(rootPath+lastPath)
}

func toLinux(path string) {
	winPath := path
	rootPath := ""
	lastPath := winPath
	reg := regexp.MustCompile(`^[A-Z]:(//|\\)`)
	if reg.MatchString(winPath) {
		rootPath = fmt.Sprintf("/mnt/%v/", strings.ToLower(winPath[0:1]))
		lastPath = winPath[3:]
		if lastPath[0:1] == "/" {
			lastPath = lastPath[1:]
		}
	}
	reg = regexp.MustCompile(`\\`)
	lastPath = reg.ReplaceAllString(lastPath, "/")
	fmt.Println(rootPath+lastPath)
}
