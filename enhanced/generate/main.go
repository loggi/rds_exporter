// The following directive is necessary to make the package coherent:
// This program generates metrics.go. It can be invoked by running
// go generate
package main

import (
	"log"
	"os"
	"text/template"

	"github.com/prometheus/client_golang/prometheus"
)

type Group struct {
	Name    string
	metrics []Metric
}

type Metric struct {
	Group string
	Name  string
	Help  string
}

func (m Metric) FqName() string {
	switch m.Group {
	case "cpuUtilization":
	}
	return prometheus.BuildFQName("rdsosmetrics", m.Group, m.Name)
}

func (m Metric) Labels() []string {
	labels := []string{
		"instance",
		"region",
	}
	switch m.Group {
	case "processList",
		"network",
		"diskIO",
		"fileSys":
		labels = append(labels, "id")
	}

	return labels
}

func (m Metric) ConstLabels() map[string]string {
	switch m.Group {
	case "cpuUtilization":
	}

	switch m.Name {
	case "CPUUtilization":
		return map[string]string{
			"cpu":  "All",
			"mode": "idle",
		}
	}

	return nil
}

func (g Group) Metrics() []Metric {
	return g.metrics
}

var (
	// http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html
	docs = map[string]map[string]string{
		"General": {
			"engine":             "The database engine for the DB instance.",
			"instanceID":         "The DB instance identifier.",
			"instanceResourceID": "A region-unique, immutable identifier for the DB instance, also used as the log stream identifier.",
			"numVCPUs":           "The number of virtual CPUs for the DB instance.",
			"timestamp":          "The time at which the metrics were taken.",
			"uptime":             "The amount of time that the DB instance has been active.",
			"version":            "The version of the OS metrics' stream JSON format.",
		},
		"cpuUtilization": {
			"guest":  "The percentage of CPU in use by guest programs.",
			"idle":   "The percentage of CPU that is idle.",
			"irq":    "The percentage of CPU in use by software interrupts.",
			"nice":   "The percentage of CPU in use by programs running at lowest priority.",
			"steal":  "The percentage of CPU in use by other virtual machines.",
			"system": "The percentage of CPU in use by the kernel.",
			"total":  "The total percentage of the CPU in use. This value includes the nice value.",
			"user":   "The percentage of CPU in use by user programs.",
			"wait":   "The percentage of CPU unused while waiting for I/O access.",
		},
		"diskIO": {
			"avgQueueLen":     "The number of requests waiting in the I/O device's queue. This metric is not available for Amazon Aurora.",
			"avgReqSz":        "The average request size, in kilobytes. This metric is not available for Amazon Aurora.",
			"await":           "The number of milliseconds required to respond to requests, including queue time and service time. This metric is not available for Amazon Aurora.",
			"device":          "The identifier of the disk device in use. This metric is not available for Amazon Aurora.",
			"readIOsPS":       "The number of read operations per second. This metric is not available for Amazon Aurora.",
			"readKb":          "The total number of kilobytes read. This metric is not available for Amazon Aurora.",
			"readKbPS":        "The number of kilobytes read per second. This metric is not available for Amazon Aurora.",
			"rrqmPS":          "The number of merged read requests queued per second. This metric is not available for Amazon Aurora.",
			"tps":             "The number of I/O transactions per second. This metric is not available for Amazon Aurora.",
			"util":            "The percentage of CPU time during which requests were issued. This metric is not available for Amazon Aurora.",
			"writeIOsPS":      "The number of write operations per second. This metric is not available for Amazon Aurora.",
			"writeKb":         "The total number of kilobytes written. This metric is not available for Amazon Aurora.",
			"writeKbPS":       "The number of kilobytes written per second. This metric is not available for Amazon Aurora.",
			"wrqmPS":          "The number of merged write requests queued per second. This metric is not available for Amazon Aurora.",
			"readLatency":     "TODO",
			"writeLatency":    "TODO",
			"writeThroughput": "TODO",
			"readThroughput":  "TODO",
			"diskQueueDepth":  "TODO",
		},
		"fileSys": {
			"maxFiles":        "The maximum number of files that can be created for the file system.",
			"mountPoint":      "The path to the file system.",
			"name":            "The name of the file system.",
			"total":           "The total number of disk space available for the file system, in kilobytes.",
			"used":            "The amount of disk space used by files in the file system, in kilobytes.",
			"usedFilePercent": "The percentage of available files in use.",
			"usedFiles":       "The number of files in the file system.",
			"usedPercent":     "The percentage of the file-system disk space in use.",
		},
		"loadAverageMinute": {
			"fifteen": "The number of processes requesting CPU time over the last 15 minutes.",
			"five":    "The number of processes requesting CPU time over the last 5 minutes.",
			"one":     "The number of processes requesting CPU time over the last minute.",
		},
		"memory": {
			"active":         "The amount of assigned memory, in kilobytes.",
			"buffers":        "The amount of memory used for buffering I/O requests prior to writing to the storage device, in kilobytes.",
			"cached":         "The amount of memory used for caching file system–based I/O.",
			"dirty":          "The amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes.",
			"free":           "The amount of unassigned memory, in kilobytes.",
			"hugePagesFree":  "The number of free huge pages.Huge pages are a feature of the Linux kernel.",
			"hugePagesRsvd":  "The number of committed huge pages.",
			"hugePagesSize":  "The size for each huge pages unit, in kilobytes.",
			"hugePagesSurp":  "The number of available surplus huge pages over the total.",
			"hugePagesTotal": "The total number of huge pages for the system.",
			"inactive":       "The amount of least-frequently used memory pages, in kilobytes.",
			"mapped":         "The total amount of file-system contents that is memory mapped inside a process address space, in kilobytes.",
			"pageTables":     "The amount of memory used by page tables, in kilobytes.",
			"slab":           "The amount of reusable kernel data structures, in kilobytes.",
			"total":          "The total amount of memory, in kilobytes.",
			"writeback":      "The amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes.",
		},
		"network": {
			"interface": "The identifier for the network interface being used for the DB instance.",
			"rx":        "The number of bytes received per second.",
			"tx":        "The number of bytes uploaded per second.",
		},
		"processList": {
			"cpuUsedPc":    "The percentage of CPU used by the process.",
			"id":           "The identifier of the process.",
			"memoryUsedPc": "The amount of memory used by the process, in kilobytes.",
			"name":         "The name of the process.",
			"parentID":     "The process identifier for the parent process of the process.",
			"rss":          "The amount of RAM allocated to the process, in kilobytes.",
			"tgid":         "The thread group identifier, which is a number representing the process ID to which a thread belongs.This identifier is used to group threads from the same process.",
			"vss":          "The amount of virtual memory allocated to the process, in kilobytes.",
			"vmlimit":      "TODO",
		},
		"swap": {
			"cached": "The amount of swap memory, in kilobytes, used as cache memory.",
			"free":   "The total amount of swap memory free, in kilobytes.",
			"total":  "The total amount of swap memory available, in kilobytes.",
		},
		"tasks": {
			"blocked":  "The number of tasks that are blocked.",
			"running":  "The number of tasks that are running.",
			"sleeping": "The number of tasks that are sleeping.",
			"stopped":  "The number of tasks that are stopped.",
			"total":    "The total number of tasks.",
			"zombie":   "The number of child tasks that are inactive with an active parent task.",
		},
	}
)

func main() {
	f, err := os.Create("metrics.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	groups := []Group{}
	for groupName, doc := range docs {
		metrics := []Metric{}
		for metricName, metricHelp := range doc {
			metric := Metric{
				Group: groupName,
				Name:  metricName,
				Help:  metricHelp,
			}
			metrics = append(metrics, metric)
		}
		group := Group{
			Name:    groupName,
			metrics: metrics,
		}
		groups = append(groups, group)
	}

	packageTemplate.Execute(f, struct {
		Groups []Group
	}{
		Groups: groups,
	})
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package enhanced

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Metrics = map[string]Metric{
{{- range .Groups }}
{{- range .Metrics }}
	"{{.FqName}}" : {
		Name: "{{.Name}}",
		Desc: prometheus.NewDesc(
			"{{.FqName}}",
			"{{.Help}}",
			{{printf "%#v" .Labels}},
			{{printf "%#v" .ConstLabels}},
		),
	},
{{- end }}
{{- end }}
}
`))
