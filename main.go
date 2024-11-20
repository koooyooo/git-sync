package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/koooyooo/git-sync/model"
	"github.com/koooyooo/git-sync/util/file"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := control(); err != nil {
		log.Fatal(err)
	}
}

func control() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	confDir := fmt.Sprintf("%s/.git-sync", home)
	if !file.Exists(confDir) {
		_ = os.Mkdir(confDir, 0755)
	}
	b, err := os.ReadFile(fmt.Sprintf("%s/config.yaml", confDir))
	if err != nil {
		return err
	}
	var conf model.Config
	if err := yaml.Unmarshal(b, &conf); err != nil {
		return err
	}
	for _, d := range conf.Dirs {
		fmt.Printf("target: [%s] %s \n", d.Name, d.Path)
		path := d.Path
		if !file.Exists(path) {
			fmt.Printf("unexistence of path: %s\n", path)
			continue
		}

		path = strings.ReplaceAll(path, "~", home)
		path = strings.ReplaceAll(path, "$HOME", home)
		path = strings.ReplaceAll(path, "${HOME}", home)

		commands := [][]string{
			{"git", "-C", path, "pull"},
			{"git", "-C", path, "add", "."},
			{"git", "-C", path, "commit", "-m", "add new commit"},
			{"git", "-C", path, "push"},
		}
		for _, c := range commands {
			cmd := exec.Command(c[0], c[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	}
	return nil
}
