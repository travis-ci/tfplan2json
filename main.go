package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hashicorp/terraform/plans/planfile"
	"github.com/hashicorp/terraform/version"
)

func main() {
	flag.Usage = func() {
		fmt.Printf(`Usage: tfplan2json <plan-file>
terraform version: %s
`, version.String())
	}
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		log.Fatal("You must pass a plan file as first argument!")
	}

	planFile := os.Args[1]

	fi, err := os.Stat(planFile)
	if err != nil {
		flag.Usage()
		log.Fatalf("Error encountered when checking plan file: %s", err)
	}

	if fi.Size() <= 0 {
		flag.Usage()
		log.Fatal("Plan file %s is empty", planFile)
	}

	plan, err := planfile.Open(planFile)
	if err != nil {
		flag.Usage()
		log.Fatalf("Error reading plan! %s", err)
	}

	encoded, err := json.MarshalIndent(plan, "", "  ")
	if err != nil {
		log.Fatalf("Error attempting to serialize plan to JSON! %s", err)
	}

	io.WriteString(os.Stdout, string(encoded))
}
