/*
Copyright 2019 The Unity Scheduler Authors

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

package main

import (
    "flag"
    "os"
)

func init() {
    flag.Set("logtostderr", "true")
}

var (
    endpoint = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
)

func main() {
    flag.Parse()

    handle()
    os.Exit(0)
}

func handle() {
    scheduler := &SimpleScheduler{}
    scheduler.Run("tcp://localhost:3333")
}
