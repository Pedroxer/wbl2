package main

import (
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

// saveSinglePage save a page. Just copy body to html file
func saveSinglePage(path string) error {
	resp, err := http.Get(path) // Receive page
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create("index.html") // Create file and write page's body to it
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

type Link struct {
	Path     string
	Dir      string
	Filename string
}

func (l *Link) getFilePath() string {
	dir := strings.TrimSuffix(l.Dir, "/")
	osDir, _ := os.Getwd()
	return "file://" + osDir + "/" + dir + "/" + l.Filename
}
func newLink(u *url.URL) *Link {
	if u.Path == rootPath {
		return &Link{Path: u.Path, Dir: rootDir, Filename: "index.html"}
	}

	if strings.HasSuffix(u.Path, "/") {
		l := Link{Path: u.String(), Dir: u.Host + u.Path, Filename: u.RequestURI()[len(u.Path):]}
		if l.Filename == "" {
			l.Filename = "index.html"
		}
		return &l
	}
	s := strings.Split(u.Host+u.Path, "/")
	l := Link{u.String(), strings.TrimSuffix(u.Host+u.Path, s[len(s)-1]), s[len(s)-1] + u.RequestURI()[len(u.Path):]}
	l.Filename = strings.ReplaceAll(l.Filename, "/", "")
	if !strings.HasSuffix(l.Filename, ".html") {
		l.Filename = "'" + l.Filename + "'"
	}
	return &l
}

func download(l *Link, lvl int) error {
	if lvl < 0 {
		return nil
	}

	fmt.Println("download: ", l.Path)

	resp, err := http.Get(l.Path)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	body = bytes.ReplaceAll(body, []byte("&amp;"), []byte("&"))

	paths, err := getLinks(body, l.Path)
	if err != nil {
		return err
	}

	sort.Slice(paths, func(i, j int) bool { return len(paths[i]) > len(paths[j]) })

	pathPref := strings.TrimSuffix(l.Path, "/")

	var links []*Link
	for _, p := range paths {
		if strings.HasPrefix(p, "/") {
			p = pathPref + p
		}
		u, err := url.Parse(p)
		if err != nil {
			return err
		}
		links = append(links, newLink(u))
	}
	replaceLinks(&body, paths, links)

	err = savePage(body, l)
	if err != nil {
		return err
	}
	for _, link := range links {
		err := download(link, lvl-1)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveWebSite(path string) error {
	uLinks[path] = struct{}{}
	u, err := url.Parse(path)
	if err != nil {
		return err
	}

	rootPath = u.Scheme + "://" + u.Host
	rootDir = u.Host

	fmt.Println("downloading web site:", rootPath)

	link := newLink(u)

	err = download(link, level)
	if err != nil {
		return err
	}

	return nil
}
func getLinks(body []byte, path string) ([]string, error) {
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var links []string
	f(doc, &links, path)

	return links, nil

}

func f(n *html.Node, links *[]string, path string) {
	path = strings.TrimSuffix(path, "/")

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && (strings.HasPrefix(a.Val, rootPath) || strings.HasPrefix(a.Val, "/")) {
				absPath := a.Val
				if strings.HasPrefix(absPath, "/") {
					absPath = path + absPath
				}
				if _, ok := uLinks[absPath]; !ok {
					uLinks[absPath] = struct{}{}
					*links = append(*links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c, links, path)
	}
}

func replaceLinks(body *[]byte, paths []string, links []*Link) {
	for i, l := range links {
		fmt.Println(l.Path)
		fmt.Println(l.Dir)
		fmt.Println(l.Filename)
		fmt.Println(l.getFilePath())
		*body = bytes.ReplaceAll(*body, []byte(paths[i]), []byte(l.getFilePath()))
	}
}

func savePage(body []byte, l *Link) error {
	err := os.MkdirAll(l.Dir, os.ModePerm)
	if err != nil {
		return err
	}
	oldDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(oldDir)

	err = os.Chdir(l.Dir)
	if err != nil {
		return err
	}

	file, err := os.Create(l.Filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return nil
}

var uLinks = make(map[string]struct{})

var rootDir string
var rootPath string

var level int
var recursive bool

func init() {
	flag.IntVar(&level, "l", 1, "depth for recursive download")
	flag.BoolVar(&recursive, "r", false, "enable recursive download")
	flag.Parse()
}

func main() {
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "no url provided")
		return
	}
	path := flag.Arg(0)

	if recursive {
		if level < 1 {
			fmt.Fprintf(os.Stderr, "level of recursion starts with 1, found: %d", level)
			return
		}

		if err := saveWebSite(path); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println("downloading complete")
	} else {
		fmt.Println("downloading web-page:", path)
		if err := saveSinglePage(path); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println("downloaded")
	}

}
