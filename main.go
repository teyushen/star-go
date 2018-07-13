package main

import (
	"github.com/teyushen/star-go/cmd"
	"log"
	"io/ioutil"
)

func init() {
	log.SetOutput(ioutil.Discard)
}


func main() {

	cmd.Cli()

}