package main

import (
	"fmt"
	"os/exec"

	"github.com/yanzay/huho/templates"
	"github.com/yanzay/log"
)

func deploy(path string, project templates.Project) error {
	url := fmt.Sprintf("https://%s", project.URL)
	volume := fmt.Sprintf("%s:/www", path)
	cmd := exec.Command("docker", "run", "--rm", "-v", volume, "yanzay/hugo-builder", url)
	res, err := cmd.CombinedOutput()
	log.Debug(string(res))
	return err
}
