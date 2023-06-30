package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the application command line ready to be running
func Generate() *cli.App{
	app := cli.NewApp()
	app.Name = "Website Analyser"
	app.Usage = "Search IPs and Name Servers on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "amazon.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Lookup IPs from adresses on the internet",
			Flags: flags,
			Action: lookupIps,
		},
		{
			Name: "servers",
			Usage: "Lookup for server names on the internet",
			Flags: flags,
			Action: lookupServers,
		},
		
	}

	

	return app
}

func lookupIps(c *cli.Context){
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for i, ip := range ips{
		n := i +1
		fmt.Println("Server", n,":",ip)
	}

}

func lookupServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for i, server := range servers{
		n := i + 1
		fmt.Println("Nameserver", n,":", server.Host)
	}
}