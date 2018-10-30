package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hashicorp/terraform/terraform"
	tfversion "github.com/hashicorp/terraform/version"
)

func main() {
	flag.Usage = func() {
		fmt.Printf(`Usage: tfplan2json
terraform version: %s

Expects a tfplan via stdin
`, tfversion.Version)
	}
	flag.Parse()

	stat, err := os.Stdin.Stat()
	if err != nil {
		flag.Usage()
		log.Fatalf("Error encountered when checking stdin: %s", err)
	}

	if stat.Size() <= 0 {
		flag.Usage()
		log.Fatal("You must pass a plan file via stdin!")
	}

	plan, err := terraform.ReadPlan(os.Stdin)
	if err != nil {
		flag.Usage()
		log.Fatalf("Error reading plan! %s\n\nAre you sure the plan file was passed?", err)
	}

	encoded, err := json.MarshalIndent(plan, "", "  ")
	if err != nil {
		log.Fatalf("Error attempting to serialize plan to JSON! %s", err)
	}

	io.WriteString(os.Stdout, string(encoded))
}
