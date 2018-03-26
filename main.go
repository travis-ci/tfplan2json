package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/hashicorp/terraform/terraform"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("Error encountered when checking stdin: %s", err)
	}

	if stat.Size() <= 0 {
		log.Fatal("You must pass a plan file via stdin!")
	}

	plan, err := terraform.ReadPlan(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading plan! %s\n\nAre you sure the plan file was passed?", err)
	}

	encoded, err := json.MarshalIndent(plan, "", "  ")
	if err != nil {
		log.Fatalf("Error attempting to serialize plan to JSON! %s", err)
	}

	io.WriteString(os.Stdout, string(encoded))
}
