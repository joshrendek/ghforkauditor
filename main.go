package main

import (
	"flag"
	"fmt"
	"github.com/apcera/termtables"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type repoInfo struct {
	Name      string
	Forks     int
	ForkUsers []string
	URL       string
}

func main() {
	var (
		oauth string
		org   string
	)
	flag.StringVar(&oauth, "oauth", "", "oauth token to use")
	flag.StringVar(&org, "org", "", "org name")
	flag.Parse()

	if len(oauth) == 0 || len(org) == 0 {
		fmt.Println("oauth and org need to be set.")
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: oauth},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	repositories := []repoInfo{}
	repos := []github.Repository{}

	for {
		reps, resp, err := client.Repositories.ListByOrg(org, opt)
		if err != nil {
			panic(err)
		}

		repos = append(repos, reps...)

		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}

	for _, r := range repos {
		name := *r.Name
		forkCount := *r.ForksCount
		isPrivate := *r.Private
		url := fmt.Sprintf("https://github.com/%s/%s", org, name)
		if forkCount > 0 && isPrivate {
			forks, _, _ := client.Repositories.ListForks(*r.Owner.Login, name, &github.RepositoryListForksOptions{Sort: "newest"})
			users := []string{}
			for _, f := range forks {
				users = append(users, *f.Owner.Login)

			}
			tmp := repoInfo{Name: name, Forks: forkCount, ForkUsers: users, URL: url}
			if len(users) > 0 {
				repositories = append(repositories, tmp)
			}
		}
	}
	table := termtables.CreateTable()
	table.AddHeaders("Repository", "URL", "Forks", "Users")
	for _, r := range repositories {
		table.AddRow(r.Name, r.URL, r.Forks, strings.Join(r.ForkUsers, ", "))
	}
	fmt.Println(table.Render())
}
