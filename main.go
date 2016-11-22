package main

import (
	"math/rand"
	"time"

	"github.com/the-anna-project/annactl/command"
)

var (
	gitCommit      string
	goArch         string
	goOS           string
	goVersion      string
	projectVersion string
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	annactlCommand := command.New()

	annactlCommand.VersionCommand().SetGitCommit(gitCommit)
	annactlCommand.VersionCommand().SetGoArch(goArch)
	annactlCommand.VersionCommand().SetGoOS(goOS)
	annactlCommand.VersionCommand().SetGoVersion(goVersion)
	annactlCommand.VersionCommand().SetProjectVersion(projectVersion)

	annactlCommand.New().Execute()
}
