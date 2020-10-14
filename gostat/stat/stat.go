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

func CatMemory() interface{} {
	read, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	return read
}

func CatCPU() interface{} {
	read, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	return read
}

func CatDisk() interface{} {
	read, err := disk.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	return read
}

func CatUptime() interface{} {
	read, err := uptime.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	return read
}

func CatNetwork() interface{} {
	read, err := network.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	return read
}

func CatLoad() interface{} {
	read, err := loadavg.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	return read
}

func CatAll() interface{} {
	var aio map[string]interface{}
	aio = make(map[string]interface{})
	aio["cpu"] = CatCPU()
	aio["mem"] = CatMemory()
	aio["disk"] = CatDisk()
	aio["uptime"] = CatUptime()
	aio["load"] = CatLoad()
	aio["network"] = CatNetwork()
	return aio
}

func RetJSON(in interface{}) []byte {
	json, err := json.Marshal(in)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return json
}
