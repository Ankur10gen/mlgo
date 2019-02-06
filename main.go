package main

import (
	"log"
	"os"

	"github.com/Ankur10gen/mlgo/mlgo"
	"github.com/jessevdk/go-flags"
)

var parser = flags.NewNamedParser("mlgo",flags.Default)

var purpose mlgo.Spin
var standaloneOpts mlgo.StandAloneOpts
var replOpts mlgo.ReplicaSetOpts

func main()  {
	// option groups
	parser.AddGroup("purpose","purpose",&purpose)
	parser.AddGroup("standalone","standalone",&standaloneOpts)
	_,err := parser.AddGroup("replicaset","replicaset",&replOpts)
	if err!=nil{
		log.Println(err)
	}

	if _,err := parser.Parse(); err!=nil{
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			log.Println(err)
			os.Exit(1)
		}
	}

	// if --spin is set then we check other options else TODO
	if purpose.Spind {
		switch {
		case !standaloneOpts.Repl:
			// Spin a standalone
			mlgo.StartStandalone(&standaloneOpts)
		case standaloneOpts.Repl:
			mlgo.StartReplicaSet(&standaloneOpts,&replOpts)
		}
	}
}
