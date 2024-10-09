package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

// The argument to newPost can be decoupled from fs.File to a more generic io.Reader
func newPost(postBody io.Reader) (Post, error) {
	// A scanner provides an interface for reading data such as a file of newline-delimited
	// lines of test
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		// Scan reads a line, Text extracts the text data
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	// postData, err := io.ReadAll(postFile) // reads the entire content of each file
	// if err != nil {
	// 	return Post{}, err
	// }

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")
	body := readBody(scanner)

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore the next line

	buf := bytes.Buffer{} // For the content of the body, we write the data into a buffer
	for scanner.Scan() {  // Scan returns bool, so we can use it as 'for' loop condition
		fmt.Fprintln(&buf, scanner.Text()) // Scan removes newline, hence we need to reinsert a newline using Fprintln
	}
	return strings.TrimSuffix(buf.String(), "\n") // To remove the last newline
}
