package main

import (
	"Ansem/internal"
	"context"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"text/tabwriter"
)

type conf struct {
	Directory      string `yaml:"exploits_dir"`
	Tick           int    `yaml:"tick"`
	TeamFile       string `yaml:"team_file"`
	GameServer     string `yaml:"gameserver"`
	Workers        int    `yaml:"workers"`
	SubmissionType string `yaml:"submission_type"`
	FlagRegex      string `yaml:"flag_regex"`
	FlagAccepted   string `yaml:"flag_accepted"`
	Token          string `yaml:"token"`
}

// Function that initialize the config
func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("configs/conf.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()

	//Fix path if the last char is not "/"
	if c.Directory[len(c.Directory)-1] != '/' {
		c.Directory = fmt.Sprintf("%s/", c.Directory)
	}

	//Aligned print
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 8, 0, '\t', 0)
	_, _ = fmt.Fprintf(writer, "Hi, I'm starting with these settings:\n\n"+
		"Exploits Dir:\t%s\n"+
		"Gameserver:\t%s\n"+
		"Teamfile:\t%s\n"+
		"SubmissionType:\t%s\n"+
		"Flag Regex:\t%s\n"+
		"Tick:\t%d\n"+
		"Workers:\t%d\n"+
		"Token:\t%s\n",

		c.Directory, c.GameServer, c.TeamFile, c.SubmissionType, c.FlagRegex, c.Tick, c.Workers, c.Token)
	writer.Flush()

	toSubmit := make(chan string, 20)

	wg := sync.WaitGroup{}
	wg.Add(2)

	exploitCtx := context.Background()
	exploitCtx = context.WithValue(exploitCtx, "exploitDir", c.Directory)
	exploitCtx = context.WithValue(exploitCtx, "tick", c.Tick)
	exploitCtx = context.WithValue(exploitCtx, "fileTeam", c.TeamFile)
	exploitCtx = context.WithValue(exploitCtx, "workers", c.Workers)
	exploitCtx = context.WithValue(exploitCtx, "submit", toSubmit)
	exploitCtx = context.WithValue(exploitCtx, "flagRegex", c.FlagRegex)

	submitterCtx := context.Background()
	submitterCtx = context.WithValue(submitterCtx, "gameServer", c.GameServer)
	submitterCtx = context.WithValue(submitterCtx, "submit", toSubmit)
	submitterCtx = context.WithValue(submitterCtx, "flagRegex", c.FlagRegex)
	submitterCtx = context.WithValue(submitterCtx, "subType", c.SubmissionType)
	submitterCtx = context.WithValue(submitterCtx, "flagAccepted", c.FlagAccepted)
	submitterCtx = context.WithValue(submitterCtx, "token", c.Token)

	go internal.StartExploiter(exploitCtx, &wg)
	go internal.StartSubmitter(submitterCtx, &wg)

	wg.Wait()
}
