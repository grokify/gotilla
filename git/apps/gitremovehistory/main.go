package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/grokify/simplego/cmd/cmdutil"
	"github.com/grokify/simplego/io/ioutilmore"
	"github.com/jessevdk/go-flags"
)

// CLI App for:
// https://help.github.com/en/articles/removing-sensitive-data-from-a-repository

type cliOptions struct {
	File string `short:"f" long:"file" description:"Git filepath" required:"true"`
	Exec []bool `short:"e" long:"exec" description:"execute" required:"false"`
}

const (
	GitCmdFilterBranchFormat string = `git filter-branch --force --index-filter "git rm --cached --ignore-unmatch %s" --prune-empty --tag-name-filter cat -- --all`
	GitCmdForceAll           string = `git push origin --force --all`
	GitCmdForceTags          string = `git push origin --force --tags`
	GitCmdDeleteRefs         string = `git for-each-ref --format="delete %(refname)" refs/original | git update-ref --stdin`
	GitCmdReflogExpire       string = "git reflog expire --expire=now --all"
	GitCmdGcPruneNow         string = "git gc --prune=now"
)

func main() {
	opts := cliOptions{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if 1 == 0 {
		opts.File, err = filepath.Abs(opts.File)
		if err != nil {
			log.Fatal(err)
		}
	}

	isFile, err := ioutilmore.IsFile(opts.File)
	if err != nil {
		log.Fatal(err)
	} else if !isFile {
		log.Fatal(fmt.Sprintf("[%s] is not a file.", opts.File))
	}

	cmds := []string{
		GitCmdFilterBranchFormat,
		GitCmdForceAll,
		GitCmdForceTags,
		GitCmdDeleteRefs,
		GitCmdReflogExpire,
		GitCmdGcPruneNow}
	cmds[0] = fmt.Sprintf(GitCmdFilterBranchFormat, opts.File)

	l := len(cmds)
	for i, cmd := range cmds {
		fmt.Printf("[%v/%v] %v\n", i+1, l, cmd)
		if len(opts.Exec) > 0 {
			stdout, stderr, err := cmdutil.ExecSimple(cmd)
			if err != nil {
				log.Fatal(err)
			}
			stdoutString := strings.TrimSpace(stdout.String())
			stderrString := strings.TrimSpace(stderr.String())
			if len(stderrString) > 0 {
				log.Fatal(stderrString)
			} else if len(stdoutString) > 0 {
				fmt.Printf(stdoutString)
			}
		}
	}

	fmt.Println("DONE")
}
