//  
//  Copyright 2023 PayPal Inc.
//  
//  Licensed to the Apache Software Foundation (ASF) under one or more
//  contributor license agreements.  See the NOTICE file distributed with
//  this work for additional information regarding copyright ownership.
//  The ASF licenses this file to You under the Apache License, Version 2.0
//  (the "License"); you may not use this file except in compliance with
//  the License.  You may obtain a copy of the License at
//  
//     http://www.apache.org/licenses/LICENSE-2.0
//  
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//  
//  Package utility provides the utility interfaces for mux package
//  
package stats

import (
	"fmt"
	"os"

	"juno/cmd/storageserv/stats/shmstats"
	"juno/pkg/cmd"
)

var _ cmd.ICommand = (*CmdStorageStats)(nil)

type CmdStorageStats struct {
	cmd.Command
	optPid      int
	optWorkerId string
}

func (c *CmdStorageStats) Init(name string, desc string) {
	c.Command.Init(name, desc)
	c.IntOption(&c.optPid, "p|pid", -1, "specify the pid of the storageserv")
	c.StringOption(&c.optWorkerId, "w|worker-id", "", "specify worker id. print stats for all workers, if \"all\"")
}

func (c *CmdStorageStats) Parse(args []string) (err error) {
	if err = c.Option.Parse(args); err != nil {
		return
	}
	if c.optPid == -1 {
		err = fmt.Errorf("specify a valid storageserv pid")
		return
	}
	return
}

func (c *CmdStorageStats) Exec() {
	if err := shmstats.InitForRead(c.optPid); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	shmstats.PrettyPrint(os.Stdout, c.optWorkerId)
	shmstats.Finalize()
}
