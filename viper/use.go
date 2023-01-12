package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"os"
)

// urfave/cli框架 && viper 的使用
func main() {
	app := cli.NewApp()

	// base application info
	app.Name = "tpcheer"
	app.Author = "tpcheer"
	app.Version = "0.0.1"
	app.Copyright = "tpcheer team reserved"
	app.Usage = "tpcheer server"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "./conf/config.toml",
			Usage: "load configuration from `FILE`",
		},
		cli.BoolFlag{
			Name:  "cpuprofile",
			Usage: "enable cpu profile",
		},
	}
	app.Action = serve
	app.Run(os.Args)
}

func serve(c *cli.Context) error {
	viper.SetConfigType("toml")
	viper.SetConfigFile(c.String("config")) // 采用了 cli.StringFlag中name的默认value
	viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("ddd---")
	})

	fmt.Println("aa-----")
	fmt.Println(viper.GetString("core.subanm"))
	return nil
}
