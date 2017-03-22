package templates

type State struct {
	GithubURL string
	Login     string
	Projects  []Project
}

type Project struct {
	URL    string
	Domain string
}
