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

const (
	ConfigDirName  = ".git-sync"
	ConfigFileName = "config.yaml"
)

func main() {
	if err := control(); err != nil {
		log.Fatal(err)
	}
}

func control() error {
	customizeMsg := flag.Bool("c", false, "customize message")
	flag.Parse()

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	conf, err := loadConfig(home, ConfigDirName, ConfigFileName)
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

		commands := buildCommands(path, customizeMsg, commitMsg)

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

func loadConfig(home, configDirName, configFileName string) (*model.Config, error) {
	confDir := fmt.Sprintf("%s/%s", home, configDirName)
	if !file.Exists(confDir) {
		if err := os.Mkdir(confDir, 0755); err != nil {
			return nil, fmt.Errorf("fail in crating .git-sync dir: %v", err)
		}
	}
	b, err := os.ReadFile(fmt.Sprintf("%s/%s", confDir, configFileName))
	if err != nil {
		return nil, err
	}
	var conf model.Config
	if err := yaml.Unmarshal(b, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func buildCommands(path string, customizeMsg *bool, commitMsg string) [][]string {
	commands := [][]string{
		{"git", "-C", path, "pull"},
		{"git", "-C", path, "add", "."},
		{"git", "-C", path, "commit", "-m", commitMsg},
		{"git", "-C", path, "push"},
	}
	if *customizeMsg {
		commands[2] = []string{"git", "-C", path, "commit"}
	}
	return commands
}
