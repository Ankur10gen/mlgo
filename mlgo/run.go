package mlgo

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// StartStandalone starts the standalone mongod instance
func StartStandalone(opts *StandAloneOpts)  {
	var cmd Command
	cmd = prepareStandalone(opts)
	cmd.Execute()
}

// StartReplicaSet starts the entire replica set
func StartReplicaSet(stOpts *StandAloneOpts, replOpts *ReplicaSetOpts) {
	var cmd Command

	var prep = &calculatedReplicaSet{
		stOpts.Port,
		"",
		"",
	}
	prep.port = stOpts.Port - 1

	for i := 0; i < replOpts.Members; i++ {
		//TODO: Path for windows
		prep.dbpath = fmt.Sprintf("%s/%s/%d", stOpts.DbPath, replOpts.Name, i)
		prep.logpath = fmt.Sprintf("%s/mongod.log ", prep.dbpath)
		prep.port = prep.port + 1
		cmd = prepareReplicaSetMember(stOpts, replOpts, prep)
		cmd.Execute()
		if i==0{
			cmd = prepareReplicaSetInitiate(stOpts)
			cmd.Execute()
			// Wait till isMaster is true
			for !isMaster(prep.port) {
				fmt.Println("Waiting for primary")
				time.Sleep(time.Second * 1)
			}
		} else {
			cmd = prepareMemberInitiate(stOpts,prep.port,false)
			cmd.Execute()
		}
	}

	if replOpts.Arbiter {
		prep.port = prep.port + 1
		prep.dbpath = fmt.Sprintf("%s/%s/arb", stOpts.DbPath, replOpts.Name)
		prep.logpath = fmt.Sprintf("%s/mongod.log ", prep.dbpath)
		cmd = prepareReplicaSetArbiter(stOpts, replOpts, prep)
		cmd.Execute()
		cmd = prepareMemberInitiate(stOpts,prep.port,true)
		cmd.Execute()
	}
}

// isMaster checks if the member is primary
func isMaster(port uint32) bool {
	var cmd Command

	cmd.Arr = []string{
		"mongo",
		"--port",
		strconv.Itoa(int(port)),
		"--eval",
		"'db.isMaster()'",
	}

	out := cmd.Execute()
	filter := fmt.Sprintf("\"ismaster\" : true")

	return strings.Contains(out, filter)
}