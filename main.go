package main

import (
	"github.com/teyushen/star-go/cmd"
	"log"
	"io/ioutil"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(ioutil.Discard)
}


func main() {

	cmd.Cli()

}