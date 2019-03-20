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

package common

// Scheduler
const SchedulerName = "yunikorn"
const DefaultPolicyGroup = "queues"

// Cluster
const ClusterId = "my-kube-cluster"
const ClusterVersion = "0.1"
const DefaultNodeAttributeHostNameKey = "si.io/hostname"
const DefaultNodeAttributeRackNameKey = "si.io/rackname"
const DefaultRackName = "/rack-default"

// Job
const LabelJobId = "jobId"
const LabelQueueName = "queue"
const JobDefaultQueue = "root"
const DefaultPartition = "default"

// Resource
const Memory = "memory"
const CPU = "vcore"

// Spark
const SparkLabelAppId = "spark-app-id"
const SparkLabelRole = "spark-role"
const SparkLabelRoleDriver = "driver"