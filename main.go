package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	var (
		oauth string
		org   string
	)
	flag.StringVar(&oauth, "oauth", "", "oauth token to use")
	flag.StringVar(&org, "org", "", "org name")
	flag.Parse()

	if len(oauth) == 0 || len(org) == 0 {
		fmt.Println(color.RedString("oauth and org need to be set."))
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: oauth},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.ListByOrg(org, nil)
	if err != nil {
		panic(err)
	}

	c := color.New(color.FgRed)

	for _, r := range repos {
		name := *r.Name
		forkCount := *r.ForksCount
		if forkCount > 0 {
			c.Printf("%s [%d] {%s} \n", name, forkCount, *r.Owner.Login)
			forks, _, _ := client.Repositories.ListForks(*r.Owner.Login, name, &github.RepositoryListForksOptions{Sort: "newest"})
			for _, f := range forks {
				c.Printf("\t %s \n", *f.Owner.Login)
			}
		}
	}
}
