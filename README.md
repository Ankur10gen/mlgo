# Usage

```
$ go run main.go --help
Usage:
  mlgo [OPTIONS]

purpose:
      --spin         Spin a deployment

standalone:
      --port=        Port number of process (default: 27017)
      --dir=         DBPath directory (default: /data/db)
      --logpath=     Path of log file (default: /data/db/mongod.log)
      --repl         For standalone replica set

replicaset:
      --replSetName= Name of the replica set (default: rs0)
      --num=         Number of replica set members (default: 3)
      --arb          Add an arbiter as well

Help Options:
  -h, --help         Show this help message
```

# Standalone

```
go run main.go --spin
```

# Replicaset

Default 3 members

```
go run main.go --spin --port 27018 --dir /data/db --repl
```

# 3 member replicaset plus

```
go run main.go --spin --port 27018 --dir /data/db --repl
```

# 2 node replicaset plus arbiter (PSA)

```
go run main.go --spin --port 27018 --dir /data/db --repl --replSetName project0 --num 2 --arb
```

# Project Structure
## Work in Progress POC

* main.go is the main driver file which calls the parser
* It also takes decisions based on inputs
* File mlgo/mlgo.go has the command line options
* prepare.go prepares the commands for use by methods in run.go
* exec.go runs the necessary operating system calls
