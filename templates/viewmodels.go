package templates

type State struct {
	GithubURL string
	Login     string
	Projects  []Project
}

type Project struct {
	ID         string
	URL        string
	Domain     string
	AutoDeploy bool
}
