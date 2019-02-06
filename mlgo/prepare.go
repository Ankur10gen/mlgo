package mlgo

import (
	"fmt"
	"strconv"
)

// calculatedReplicaSet is a placeholder
// It is populated dynamically based on the number of replica set members
type calculatedReplicaSet struct {
	port    uint32
	dbpath  string
	logpath string
}

// prepareStandalone prepares the command for a standalone mongod instance
func prepareStandalone(opts *StandAloneOpts) Command {
	var cmd Command

	cmd.Arr = []string{
		"mongod",
		"--port",
		strconv.Itoa(int(opts.Port)),
		"--dbpath",
		opts.DbPath,
		"--logpath",
		opts.LogPath,
		"--fork",
	}

	return cmd
}

// prepareReplicaSetMember prepares the command for a replica set member
func prepareReplicaSetMember(stOpts *StandAloneOpts, replOpts *ReplicaSetOpts, prep *calculatedReplicaSet) Command {

	var cmd Command

	cmd.Arr = []string{
		"mongod",
		"--port",
		strconv.Itoa(int(prep.port)),
		"--dbpath",
		prep.dbpath,
		"--logpath",
		prep.logpath,
		"--fork",
		"--replSet",
		replOpts.Name,
	}

	return cmd
}

// prepareReplicaSetArbiter prepares the command for a replica set arbiter
func prepareReplicaSetArbiter(stOpts *StandAloneOpts, replOpts *ReplicaSetOpts, prep *calculatedReplicaSet) Command {

	var cmd Command

	cmd.Arr = []string{
		"mongod",
		"--port",
		strconv.Itoa(int(prep.port)),
		"--dbpath",
		prep.dbpath,
		"--logpath",
		prep.logpath,
		"--fork",
		"--replSet",
		replOpts.Name,
	}

	return cmd
}

// prepareReplicaSetInitiate initiates a replica set
func prepareReplicaSetInitiate(stOpts *StandAloneOpts) Command  {
	var cmd Command
	cmd.Arr = []string{
		"mongo",
		"--port",
		strconv.Itoa(int(stOpts.Port)),
		"--eval",
		"'rs.initiate()'",
	}
	return cmd
}

// prepareMemberInitiate prepares the rs.add() command for a data member/arbiter
func prepareMemberInitiate(stOpts *StandAloneOpts, memberPort uint32, isArbiter bool) Command  {
	var cmd Command
	var addArb string = fmt.Sprintf("'rs.add(\"localhost:%d\",%t)'",memberPort,isArbiter)

	cmd.Arr = []string{
		"mongo",
		"--port",
		strconv.Itoa(int(stOpts.Port)),
		"--eval",
		addArb,
	}
	return cmd
}

