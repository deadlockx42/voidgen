//
//   Copyright 2017 Deadlock X42 <deadlock.x42@gmail.com>
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/deadlockx42/voidgen/code"
	"github.com/deadlockx42/voidgen/schema"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s file\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func main() {
	force := false
	flag.BoolVar(&force, "f", false, "force use of existing package directory")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		usage()
	}

	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	g, err := schema.New(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	r, err := schema.Validate(g)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, w := range r.Warnings {
		fmt.Printf("Warning: %s\n", w)
	}
	for _, e := range r.Errors {
		fmt.Printf("Error: %s\n", e)
	}
	if len(r.Errors) != 0 {
		os.Exit(1)
	}
	err = code.Generate(g)
	if err != nil {
		log.Fatal(err.Error())
	}
}
