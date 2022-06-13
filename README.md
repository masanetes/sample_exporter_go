# Sample Exporter Go

sample implementation of prometheus exporter in golang

# Quick Start

start applications(prometheus, blackbox_exporter, my_exporter)

```
$ docker compose up -d
```

```
$ curl http://localhost:19115/metrics
# HELP google_status Number of goroutines that currently exist.
# TYPE google_status gauge
google_status 200
# HELP threads Number of OS threads created.
# TYPE threads gauge
threads 5
```


clean

```
$ docker compose stop && docker compose rm 
```
