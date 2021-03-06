// Copyright 2015 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package versioncli implements a subcommand for obtaining version with the CLI.
//
package versioncli

import (
	"fmt"

	"github.com/maruel/subcommands"
)

// CmdVersion returns a generic "version" subcommand printing the version given.
func CmdVersion(version string) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "version <options>",
		ShortDesc: "prints version number",
		LongDesc:  "Prints the tool version number.",
		CommandRun: func() subcommands.CommandRun {
			return &versionRun{version: version}
		},
	}
}

type versionRun struct {
	subcommands.CommandRunBase
	version string
}

func (c *versionRun) Run(a subcommands.Application, args []string, _ subcommands.Env) int {
	if len(args) != 0 {
		fmt.Fprintf(a.GetErr(), "%s: position arguments not expected\n", a.GetName())
		return 1
	}
	fmt.Println(c.version)
	return 0
}
