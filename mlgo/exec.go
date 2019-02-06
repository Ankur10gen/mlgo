package mlgo

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	Arr []string
}

// executeLinuxCommand provides the concrete implementation of command on Linux
func (c *Command) executeLinuxCommand() string {

	for i,v := range c.Arr{
		if v == "--dbpath"{
			_,err := os.Stat(c.Arr[i+1])
			if os.IsNotExist(err){
				err := os.MkdirAll(c.Arr[i+1],os.ModePerm)
				if err!=nil{
					log.Fatalln(err)
				}
			}
		}
	}

	var cmd string
	cmd = strings.Join(c.Arr," ")
	log.Printf("Executing: %s\n",cmd)
	stdout, err := exec.Command("sh","-c",cmd).Output()
	if err!=nil{
		log.Fatalln(err)
	}
	return string(stdout)
}

// Execute calls the OS specific command
func (c *Command) Execute() string {
	var out string

	if runtime.GOOS != "windows"{
		out = c.executeLinuxCommand()
		return out
	}
	return ""
}