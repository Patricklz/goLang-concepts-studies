package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// create will return the app ready to run
func Create() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search IP's and names of servers from internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Internet address ip search",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Server address ip search",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
