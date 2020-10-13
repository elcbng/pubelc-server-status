package utilities

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/memory"
)

func catMemory() []byte {
	mem, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	json, err := json.Marshal(mem)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return json
}
