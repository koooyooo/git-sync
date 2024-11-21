package main

import (
	"flag"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/koooyooo/git-sync/model"
	"github.com/koooyooo/git-sync/util/file"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	if err := control(); err != nil {
		log.Fatal(err)
	}
}

func control() error {
	flag.Parse()
	specMsg := flag.Bool("m", false, "specify message")

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	conf, err := loadConfig(home, err)
	if err != nil {
		return err
	}
	for _, d := range conf.Dirs {
		fmt.Printf("target: [%s] %s \n", d.Name, d.Path)

		// fix Path
		path := d.Path
		path = strings.ReplaceAll(path, "~", home)
		path = strings.ReplaceAll(path, "$HOME", home)
		path = strings.ReplaceAll(path, "${HOME}", home)

		if !file.Exists(path) {
			fmt.Printf("unexistence of path: %s\n", path)
			continue
		}

		timeStr := time.Now().Format("2006-01-02 15:04:05")
		commitMsg := fmt.Sprintf("update at: %s", timeStr)

		commands := buildCommands(path, specMsg, commitMsg)

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

func buildCommands(path string, specMsg *bool, commitMsg string) [][]string {
	commands := [][]string{
		{"git", "-C", path, "pull"},
		{"git", "-C", path, "add", "."},
		{"git", "-C", path, "commit", "-m", commitMsg},
		{"git", "-C", path, "push"},
	}
	if *specMsg {
		commands[2] = []string{"git", "-C", path, "commit", "-m", commitMsg}
	}
	return commands
}

func loadConfig(home string, err error) (model.Config, error) {
	confDir := fmt.Sprintf("%s/.git-sync", home)
	if !file.Exists(confDir) {
		_ = os.Mkdir(confDir, 0755)
	}
	b, err := os.ReadFile(fmt.Sprintf("%s/config.yaml", confDir))
	if err != nil {
		return model.Config{}, err
	}
	var conf model.Config
	if err := yaml.Unmarshal(b, &conf); err != nil {
		return model.Config{}, err
	}
	return conf, nil
}
