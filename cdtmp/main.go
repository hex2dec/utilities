package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path"
	"strings"
	"time"
)

func main() {
	username := getCurrentUsername()
	if username == "" {
		return
	}

	dir := getTempDirectoryName(username)
	if dir == "" {
		return
	}

	dir = makeTempDirectory(dir)
	if dir == "" {
		return
	}

	// TODO: copy command into clipboard

	fmt.Printf("go into working directory, run below command:\n\n\tcd %s\n\n", dir)
}

func getCurrentUsername() string {
	cu, err := user.Current()
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return ""
	}
	return cu.Username
}

func getRandomString(size int) string {
	chatSet := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	var out strings.Builder
	for i := 0; i < size; i++ {
		c := chatSet[rand.Intn(len(chatSet))]
		out.WriteString(string(c))
	}
	return out.String()
}

func getTempDirectoryName(name string) string {
	suffix := getRandomString(8)
	dir := strings.Join([]string{name, suffix}, "-")
	return dir
}

func makeTempDirectory(name string) string {
	tmp := os.TempDir()
	dir := path.Join(tmp, name)
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return ""
	}
	return dir
}
