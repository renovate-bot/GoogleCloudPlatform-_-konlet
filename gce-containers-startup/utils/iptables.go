// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("iptables command output: %s\n", string(outs))
	}
}

//COS milestone 113 started to use iptables-nft vs iptables-legacy
//This function reads the env variable passed to konlet from the host OS
//with the host OS iptables version and makes the switch to legacy when needed.
func InitIpTables() error {
        log.Print("Determining the iptables version")
        var iptables = os.Getenv("HOST_IPTABLES")
        if iptables == "" {
                return errors.New("HOST_IPTABLES environment variable is not set - cannot determine the version to be used in konlet")
        } 
        if iptables == "legacy" {
                log.Print("Detected legacy iptables on the host OS. Switching to legacy iptables.")
                var cmd = exec.Command("update-alternatives", "--set", "iptables", "/usr/sbin/iptables-legacy")

                var output, err = cmd.CombinedOutput()

                if err != nil {
                        return err
                }
                log.Printf("%s\n", output)
        } else {
                log.Print("Detected nf_tables on the host OS. Staying on the nf_tables.")
        } 
        return nil
}

func OpenIptablesForProtocol(protocol string) error {
        log.Printf("Updating IPtables firewall rules - allowing %s traffic on all ports", protocol)
        // TODO: Make it use osCommandRunner.
        var cmd = exec.Command("iptables", "-A", "INPUT", "-p", protocol, "-j", "ACCEPT")
        var output, err = cmd.CombinedOutput()

        if err != nil {
                return err
        }
        log.Printf("%s\n", output)

        cmd = exec.Command("iptables", "-A", "FORWARD", "-p", protocol, "-j", "ACCEPT")
        output, err = cmd.CombinedOutput()

        log.Printf("%s\n", output)

        if err != nil {
                return err
        }
        return nil
}

func OpenIptables() error {
	var err = OpenIptablesForProtocol("tcp")
	if err != nil {
		return err
	}
	err = OpenIptablesForProtocol("udp")
	if err != nil {
		return err
	}
	err = OpenIptablesForProtocol("icmp")
	if err != nil {
		return err
	}

	return nil
}
