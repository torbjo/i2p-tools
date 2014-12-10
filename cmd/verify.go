package cmd

import (
	"fmt"
	"os"

	"github.com/MDrollette/go-i2p/su3"
	"github.com/codegangsta/cli"
)

func NewSu3VerifyCommand() cli.Command {
	return cli.Command{
		Name:        "verify",
		Usage:       "Verify a Su3 file",
		Description: "Verify a Su3 file",
		Action:      su3VerifyAction,
		Flags:       []cli.Flag{},
	}
}

func su3VerifyAction(c *cli.Context) {
	file, err := os.Open(c.Args().Get(0))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	su3File, err := su3.ReadSu3(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(su3File.String())

	if err := su3File.VerifySignature(); nil != err {
		panic(err)
	}

	fmt.Println("Verified signature.")
}
