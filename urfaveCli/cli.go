package urfaveCli

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func Main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) {
		if c.NArg() != 0 {
			fmt.Printf("未找到命令: %s\n运行命令 %s ; help 获取帮助\n", c.Args().Get(0), app.Name)
			return
		}
		var prompt string
		prompt = app.Name + " > "
	L:
		for {
			var input string
			fmt.Print(prompt)
			//   fmt.Scanln(&input)

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // use `for scanner.Scan()` to keep reading
			input = scanner.Text()
			//fmt.Println("captured:",input)
			switch input {
			case "close":
				fmt.Println("close.")
				break L
			default:
			}
			//fmt.Print(input)
			cmdArgs := strings.Split(input, " ")
			//fmt.Print(len(cmdArgs))
			if len(cmdArgs) == 0 {
				continue
			}

			s := []string{app.Name}
			s = append(s, cmdArgs...)

			c.App.Run(s)

		}

		return
	}

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
