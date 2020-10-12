package os

import (
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/memory"
)

func catMemory() {
	mem, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	//debug output
	fmt.Printf("memory total: %d bytes\n", mem.Total)
	fmt.Printf("memory used: %d bytes\n", mem.Used)
	fmt.Printf("memory cached: %d bytes\n", mem.Cached)
	fmt.Printf("memory free: %d bytes\n", mem.Free)

}
