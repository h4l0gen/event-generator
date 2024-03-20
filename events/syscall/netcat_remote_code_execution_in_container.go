// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2024 The Falco Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package syscall

import (
    "os/exec"

    "github.com/falcosecurity/event-generator/events"
)

var _ = events.Register(NetcatRemoteCodeExecutionInContainer)

func NetcatRemoteCodeExecutionInContainer(h events.Helper) error {
    if h.InContainer() {
        // Launch netcat (nc) with the -e flag for remote code execution
        cmd := exec.Command("nc", "-e", "/bin/bash")

        h.Log().Info("Netcat runs inside container that allows remote code execution")
        err := cmd.Run()
        if err != nil {
            h.Log().WithError(err).Error("Failed to launch netcat (nc) for remote code execution")
            return err
        }
    } 

    return nil
}