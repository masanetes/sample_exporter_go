# Sample Exporter Go

sample implementation of prometheus exporter in golang

# Quick Start

start applications(prometheus, blackbox_exporter, sample_exporter_go)

```
$ docker compose up -d
```

```
$ curl http://localhost:19115/metrics
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 8.86e-05
go_gc_duration_seconds{quantile="0.25"} 0.0001178
go_gc_duration_seconds{quantile="0.5"} 0.000248
go_gc_duration_seconds{quantile="0.75"} 0.0005785
go_gc_duration_seconds{quantile="1"} 0.0009721
go_gc_duration_seconds_sum 0.0062885
go_gc_duration_seconds_count 18
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 10
```

```
$ curl http://localhost:19115/probe?target=https://kubernetes.io
# HELP http_status http status
# TYPE http_status gauge
http_status 200
```


### Clean Up

```
$ docker compose stop && docker compose rm 
```
