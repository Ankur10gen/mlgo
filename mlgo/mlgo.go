// Package mlgo prepares and runs all the necessary commands on operating system
// for initiating a deployment
package mlgo

// Spin checks if the user wants to spin a deployment
type Spin struct {
	Spind bool `long:"spin" description:"Spin a deployment"`
}

// StandAloneOpts are the basic requirements for any deployment
type StandAloneOpts struct {
	Port uint32 `long:"port" description:"Port number of process" default:"27017"`
	DbPath string `long:"dir" description:"DBPath directory" default:"/data/db"`
	LogPath string `long:"logpath" description:"Path of log file" default:"/data/db/mongod.log"`
	Repl bool `long:"repl" description:"For standalone replica set"`
}

// ReplicaSetOpts are basic requirements for a replica set
// Used in collaboration with StandAloneOpts
type ReplicaSetOpts struct {
	Name string `long:"replSetName" description:"Name of the replica set" default:"rs0"`
	Members int `long:"num" description:"Number of replica set members" default:"3"`
	Arbiter bool `long:"arb" description:"Add an arbiter as well"`
}



