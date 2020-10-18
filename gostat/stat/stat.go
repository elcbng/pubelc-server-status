package stat

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/disk"
	"github.com/mackerelio/go-osstat/loadavg"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/network"
	"github.com/mackerelio/go-osstat/uptime"
)

var read interface{}
var err error

func Catter(target string) []byte {
	switch target {
	case "memory":
		read, err = memory.Get()
		Er(err)
	case "cpu":
		read, err = cpu.Get()
		Er(err)
	case "disk":
		read, err = disk.Get()
		Er(err)
	case "uptime":
		read, err = uptime.Get()
		Er(err)
	case "network":
		read, err = network.Get()
		Er(err)
	case "load":
		read, err = loadavg.Get()
		Er(err)
	}
	dat := RetJSON(read)
	return dat
}

func CatAll() []byte {
	var aio = make(map[string]interface{})
	aio["cpu"], err = cpu.Get()
	Er(err)
	aio["mem"], err = memory.Get()
	Er(err)
	aio["disk"], err = disk.Get()
	Er(err)
	aio["uptime"], err = uptime.Get()
	Er(err)
	aio["load"], err = network.Get()
	Er(err)
	aio["network"], err = loadavg.Get()
	Er(err)
	dat := RetJSON(aio)
	return dat
}

func RetJSON(in interface{}) []byte {
	re, err := json.Marshal(in)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return re
}

func Er(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
