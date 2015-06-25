package command

import (
	"github.com/codegangsta/cli"
	"github.com/shibukawa/git4go"
	"fmt"
	"os"
)

func CmdLsTree(c *cli.Context) {
	repo, err := git4go.OpenRepositoryExtended(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(c.Args()) == 0 {
		cli.ShowSubcommandHelp(c)
	} else {
		ref, err := repo.DwimReference(c.Args().First())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		resolved, err := ref.Resolve()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		commit, err := repo.LookupCommit(resolved.Target())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		tree, err := commit.Tree()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, entry := range tree.Entries {
			fileMode := fmt.Sprintf("%06o", int(entry.Filemode))
			fmt.Printf("%s %s %s\t%s\n", fileMode, entry.Type.String(), entry.Id.String(), entry.Name)
		}
	}
}
