package main

import client "vct/httpClient"

func main() {
	cli := client.Setup(true, false)

	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}

	if err := client.DoOps(cli); err != nil {
		panic(err)
	}

	c := client.Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	cli = client.Setup(true, true)

	if err := client.DoOps(cli); err != nil {
		panic(err)
	}

	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}
}
