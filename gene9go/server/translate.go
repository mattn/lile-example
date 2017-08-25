package server

import (
	"bufio"
	"os"
	"strings"

	"github.com/mattn/lile-example/gene9go"
	context "golang.org/x/net/context"
)

func translate(word string) (string, error) {
	f, err := os.Open("gene.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	found := ""
	for scanner.Scan() {
		first := scanner.Text()
		if !scanner.Scan() {
			break
		}

		if strings.ToLower(first) == strings.ToLower(word) {
			found = scanner.Text()
			break
		}
	}
	return found, scanner.Err()
}

func (s Gene9goServer) Translate(ctx context.Context, r *gene9go.Request) (*gene9go.Response, error) {
	text, err := translate(r.GetWord())
	if err != nil {
		return nil, err
	}
	return &gene9go.Response{Text: text}, nil
}
