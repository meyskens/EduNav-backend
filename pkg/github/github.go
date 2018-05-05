package github

import (
	"time"

	"../config"
	gh "github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	client      *gh.Client
	clientSetup bool
)

// Contributor contains the name and email for each contriutor
type Contributor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarURL"`
}

// GetContributors gets all contributors of the EduNav project
func GetContributors() []Contributor {
	contributorsChan := make(chan []Contributor)
	go getContributorsForRepo("meyskens", "EduNav-backend", contributorsChan)
	go getContributorsForRepo("meyskens", "EduNav-scan", contributorsChan)
	go getContributorsForRepo("meyskens", "EduNav", contributorsChan)

	contributors := []Contributor{}

	for i := 0; i < 3; i++ {
		c := <-contributorsChan
		contributors = append(contributors, c...)
	}

	contributors = filterDuplicates(contributors)
	return contributors
}

func getContributorsForRepo(org, name string, out chan []Contributor) {
	if !clientSetup {
		setUpClient()
		clientSetup = true
	}
	contributors := []Contributor{}
	stats, _, err := client.Repositories.ListContributorsStats(context.Background(), org, name)
	if _, ok := err.(*gh.AcceptedError); ok {
		time.Sleep(1 * time.Second)
		go getContributorsForRepo(org, name, out) // repeat till we got something
		return
	}
	if err != nil {
		out <- contributors
		return
	}
	for _, ghContributor := range stats {
		username := ghContributor.Author.GetLogin()
		info, _, err := client.Users.Get(context.Background(), username)
		if err != nil {
			continue
		}
		contributor := Contributor{
			ID:        int64(info.GetID()),
			Name:      info.GetName(),
			Email:     info.GetEmail(),
			AvatarURL: info.GetAvatarURL(),
		}
		contributors = append(contributors, contributor)
	}
	out <- contributors
}

func setUpClient() {
	conf := config.GetConfiguration()
	if conf.GitHubToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: conf.GitHubToken},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		client = gh.NewClient(tc)
	} else {

		client = gh.NewClient(nil)
	}

}

func filterDuplicates(in []Contributor) []Contributor {
	hadID := map[int]bool{}
	uniqueContributors := []Contributor{}
	for _, contibutor := range in {
		if _, exists := hadID[contibutor.ID]; !exists {
			uniqueContributors = append(uniqueContributors, contibutor)
			hadID[contibutor.ID] = true
		}
	}
	return uniqueContributors
}
