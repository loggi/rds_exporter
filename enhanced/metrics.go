// Code generated by go generate; DO NOT EDIT.
package enhanced

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Metrics = map[string]Metric{
	"rdsosmetrics_cpuUtilization_guest": {
		Name: "guest",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_guest",
			"The percentage of CPU in use by guest programs.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_cpuUtilization_nice": {
		Name: "nice",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_nice",
			"The percentage of CPU in use by programs running at lowest priority.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_cpuUtilization_user": {
		Name: "user",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_user",
			"The percentage of CPU in use by user programs.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_cpuUtilization_wait": {
		Name: "wait",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_wait",
			"The percentage of CPU unused while waiting for I/O access.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_cpuUtilization_idle": {
		Name: "idle",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_idle",
			"The percentage of CPU that is idle.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_cpuUtilization_irq": {
		Name: "irq",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_irq",
			"The percentage of CPU in use by software interrupts.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_cpuUtilization_steal": {
		Name: "steal",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_steal",
			"The percentage of CPU in use by other virtual machines.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_cpuUtilization_system": {
		Name: "system",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_system",
			"The percentage of CPU in use by the kernel.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_cpuUtilization_total": {
		Name: "total",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_cpuUtilization_total",
			"The total percentage of the CPU in use. This value includes the nice value.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_avgQueueLen": {
		Name: "avgQueueLen",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_avgQueueLen",
			"The number of requests waiting in the I/O device's queue. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_readIOsPS": {
		Name: "readIOsPS",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readIOsPS",
			"The number of read operations per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_readKb": {
		Name: "readKb",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readKb",
			"The total number of kilobytes read. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_writeLatency": {
		Name: "writeLatency",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeLatency",
			"TODO",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_avgReqSz": {
		Name: "avgReqSz",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_avgReqSz",
			"The average request size, in kilobytes. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_device": {
		Name: "device",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_device",
			"The identifier of the disk device in use. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_rrqmPS": {
		Name: "rrqmPS",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_rrqmPS",
			"The number of merged read requests queued per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_util": {
		Name: "util",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_util",
			"The percentage of CPU time during which requests were issued. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_readThroughput": {
		Name: "readThroughput",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readThroughput",
			"TODO",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_readKbPS": {
		Name: "readKbPS",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readKbPS",
			"The number of kilobytes read per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_tps": {
		Name: "tps",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_tps",
			"The number of I/O transactions per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_readLatency": {
		Name: "readLatency",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readLatency",
			"TODO",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_writeThroughput": {
		Name: "writeThroughput",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeThroughput",
			"TODO",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_await": {
		Name: "await",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_await",
			"The number of milliseconds required to respond to requests, including queue time and service time. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_writeIOsPS": {
		Name: "writeIOsPS",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeIOsPS",
			"The number of write operations per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_writeKb": {
		Name: "writeKb",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeKb",
			"The total number of kilobytes written. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_writeKbPS": {
		Name: "writeKbPS",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeKbPS",
			"The number of kilobytes written per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_wrqmPS": {
		Name: "wrqmPS",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_wrqmPS",
			"The number of merged write requests queued per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_diskIO_diskQueueDepth": {
		Name: "diskQueueDepth",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_diskQueueDepth",
			"TODO",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_fileSys_usedFiles": {
		Name: "usedFiles",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_usedFiles",
			"The number of files in the file system.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_fileSys_maxFiles": {
		Name: "maxFiles",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_maxFiles",
			"The maximum number of files that can be created for the file system.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_fileSys_name": {
		Name: "name",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_name",
			"The name of the file system.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_fileSys_total": {
		Name: "total",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_total",
			"The total number of disk space available for the file system, in kilobytes.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_fileSys_used": {
		Name: "used",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_used",
			"The amount of disk space used by files in the file system, in kilobytes.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_fileSys_mountPoint": {
		Name: "mountPoint",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_mountPoint",
			"The path to the file system.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_fileSys_usedFilePercent": {
		Name: "usedFilePercent",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_usedFilePercent",
			"The percentage of available files in use.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_fileSys_usedPercent": {
		Name: "usedPercent",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_usedPercent",
			"The percentage of the file-system disk space in use.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_loadAverageMinute_one": {
		Name: "one",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_loadAverageMinute_one",
			"The number of processes requesting CPU time over the last minute.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_loadAverageMinute_fifteen": {
		Name: "fifteen",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_loadAverageMinute_fifteen",
			"The number of processes requesting CPU time over the last 15 minutes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_loadAverageMinute_five": {
		Name: "five",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_loadAverageMinute_five",
			"The number of processes requesting CPU time over the last 5 minutes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_network_interface": {
		Name: "interface",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_network_interface",
			"The identifier for the network interface being used for the DB instance.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_network_rx": {
		Name: "rx",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_network_rx",
			"The number of bytes received per second.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_network_tx": {
		Name: "tx",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_network_tx",
			"The number of bytes uploaded per second.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_General_engine": {
		Name: "engine",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_engine",
			"The database engine for the DB instance.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_General_instanceID": {
		Name: "instanceID",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_instanceID",
			"The DB instance identifier.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_General_instanceResourceID": {
		Name: "instanceResourceID",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_instanceResourceID",
			"A region-unique, immutable identifier for the DB instance, also used as the log stream identifier.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_General_numVCPUs": {
		Name: "numVCPUs",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_numVCPUs",
			"The number of virtual CPUs for the DB instance.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_General_timestamp": {
		Name: "timestamp",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_timestamp",
			"The time at which the metrics were taken.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_General_uptime": {
		Name: "uptime",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_uptime",
			"The amount of time that the DB instance has been active.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_General_version": {
		Name: "version",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_version",
			"The version of the OS metrics' stream JSON format.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_hugePagesSize": {
		Name: "hugePagesSize",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesSize",
			"The size for each huge pages unit, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_hugePagesTotal": {
		Name: "hugePagesTotal",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesTotal",
			"The total number of huge pages for the system.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_inactive": {
		Name: "inactive",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_inactive",
			"The amount of least-frequently used memory pages, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_pageTables": {
		Name: "pageTables",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_pageTables",
			"The amount of memory used by page tables, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_slab": {
		Name: "slab",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_slab",
			"The amount of reusable kernel data structures, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_cached": {
		Name: "cached",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_cached",
			"The amount of memory used for caching file system–based I/O.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_hugePagesFree": {
		Name: "hugePagesFree",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesFree",
			"The number of free huge pages.Huge pages are a feature of the Linux kernel.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_hugePagesRsvd": {
		Name: "hugePagesRsvd",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesRsvd",
			"The number of committed huge pages.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_mapped": {
		Name: "mapped",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_mapped",
			"The total amount of file-system contents that is memory mapped inside a process address space, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_total": {
		Name: "total",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_total",
			"The total amount of memory, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_writeback": {
		Name: "writeback",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_writeback",
			"The amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_buffers": {
		Name: "buffers",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_buffers",
			"The amount of memory used for buffering I/O requests prior to writing to the storage device, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_dirty": {
		Name: "dirty",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_dirty",
			"The amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_free": {
		Name: "free",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_free",
			"The amount of unassigned memory, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_active": {
		Name: "active",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_active",
			"The amount of assigned memory, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_memory_hugePagesSurp": {
		Name: "hugePagesSurp",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesSurp",
			"The number of available surplus huge pages over the total.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_vss": {
		Name: "vss",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_vss",
			"The amount of virtual memory allocated to the process, in kilobytes.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_vmlimit": {
		Name: "vmlimit",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_vmlimit",
			"TODO",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_cpuUsedPc": {
		Name: "cpuUsedPc",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_cpuUsedPc",
			"The percentage of CPU used by the process.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_id": {
		Name: "id",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_id",
			"The identifier of the process.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_name": {
		Name: "name",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_name",
			"The name of the process.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_parentID": {
		Name: "parentID",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_parentID",
			"The process identifier for the parent process of the process.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_rss": {
		Name: "rss",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_rss",
			"The amount of RAM allocated to the process, in kilobytes.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_tgid": {
		Name: "tgid",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_tgid",
			"The thread group identifier, which is a number representing the process ID to which a thread belongs.This identifier is used to group threads from the same process.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_processList_memoryUsedPc": {
		Name: "memoryUsedPc",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_memoryUsedPc",
			"The amount of memory used by the process, in kilobytes.",
			[]string{"instance", "region", "id"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_swap_cached": {
		Name: "cached",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_swap_cached",
			"The amount of swap memory, in kilobytes, used as cache memory.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_swap_free": {
		Name: "free",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_swap_free",
			"The total amount of swap memory free, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_swap_total": {
		Name: "total",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_swap_total",
			"The total amount of swap memory available, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_tasks_running": {
		Name: "running",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_running",
			"The number of tasks that are running.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_tasks_sleeping": {
		Name: "sleeping",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_sleeping",
			"The number of tasks that are sleeping.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_tasks_stopped": {
		Name: "stopped",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_stopped",
			"The number of tasks that are stopped.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_tasks_total": {
		Name: "total",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_total",
			"The total number of tasks.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_tasks_zombie": {
		Name: "zombie",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_zombie",
			"The number of child tasks that are inactive with an active parent task.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
	"rdsosmetrics_tasks_blocked": {
		Name: "blocked",
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_blocked",
			"The number of tasks that are blocked.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
	},
}
