package main

import (
	"log"

	"github.com/WoofinaS/gobooru/gel"
	"github.com/spf13/pflag"
)

var (
	key, name, tag, folder string
	c                      gel.Client
)

func init() {
	pflag.CommandLine.SortFlags = false
	pflag.StringVarP(&folder, "folder", "o", "./", "specify a output folder")
	pflag.StringVarP(&tag, "tag", "t", "", "Specify a tag to download")
	pflag.StringVar(&key, "key", "", "Specify a api key to avoid throttling")
	pflag.StringVar(&name, "name", "", "Specify a username to avoid throttling")
	pflag.Parse()
}

func main() {
	c = gel.NewClient(key, name)
	filter := gel.PostFilter{Tags: []string{tag}}
	pagenum := 0
	for i := 0; i <= pagenum; i++ {
		results, err := c.SearchPosts(filter)
		if err != nil {
			log.Fatal(err)
		}
		pagenum = int(results.Offset/100 + 1)
		for _, post := range results.Posts {
			log.Println("Downloading: " + post.FileName)
			gel.DownloadPost(post, "./")
		}
	}
}
