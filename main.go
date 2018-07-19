package main

import (
	"log"
	"io/ioutil"
	"github.com/teyushen/star-go/cmd"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(ioutil.Discard)
}


func main() {
	cmd.Cli()

}
