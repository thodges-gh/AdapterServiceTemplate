# External Adapter Service Template

This is meant to be an example for how you can set up an external adapter to make use of additional functions for a Chainlink node.

## Building

```bash
$ go build -o ea src/*
```

## Running

```bash
$ ./ea
```

## Adding to Chainlink

Make sure your Chainlink node is already running. For example, mine is running at localhost:6688 with a `$USERNAME` of "chainlink" and `$PASSWORD` of "twochains" (the defaults).

### Add the adapter

```bash
curl -u chainlink:twochains -X POST -H 'Content-Type: application/json' -d '{"name":"ea","url":"http://localhost:3000/"}' http://localhost:6688/v2/bridge_types
```

### Create the job

```bash
curl -u chainlink:twochains -X POST -H 'Content-Type: application/json' -d '{"initiators":[{"type":"web"}],"tasks":[{"type":"ea"}]}' http://localhost:6688/v2/specs
```


### Starting a run

Be sure to change the JobID to the given output from the last command.

```bash
curl -u chainlink:twochains -X POST http://localhost:6688/v2/specs/a0de434162de4e37817d9f0b9c12da3d/runs
```

## Examples

Here is an example of what the log would look like for a run

```shell
2018/03/24 15:17:52 Input:
2018/03/24 15:17:52 {"id":"87a196436d084c608a748f6f98d0c1d6","data":{}}
2018/03/24 15:17:52 Output:
2018/03/24 15:17:52 {"id":"87a196436d084c608a748f6f98d0c1d6","data":{"value":"true","last":"1111","other":"crypto"}}
```

And here is what the Chainlink node will output (with debug enabled)

```shell
{"level":"info","ts":1521923157.4895928,"caller":"web/router.go:57","msg":"Web request","method":"POST","status":200,"path":"/v2/specs/b32bddfc9eb4419f9f95dad9bfbe4cff/runs","query":"","body":"","clientIP":"::1","comment":"","servedAt":"2018/03/24 - 15:25:57","latency":"232.138µs"}
{"level":"info","ts":1521923157.511245,"caller":"services/job_runner.go:54","msg":"Starting job","job":"b32bddfc9eb4419f9f95dad9bfbe4cff","run":"b8004e2989e24e1d8e4449afad2eb480","status":"in progress"}
{"level":"debug","ts":1521923157.5142593,"caller":"services/job_runner.go:71","msg":"Produced task run","tr":"TaskRun(a742a5cb75544ce3bce936a24524d0c8,ea,completed,)"}
{"level":"info","ts":1521923157.5196233,"caller":"services/job_runner.go:81","msg":"Task ea finished","task":0,"result":"","type":"ea","params":"{\"type\":\"ea\"}","taskrun":"a742a5cb75544ce3bce936a24524d0c8","status":""}
{"level":"info","ts":1521923157.5196896,"caller":"services/job_runner.go:97","msg":"Finished current job run execution","job":"b32bddfc9eb4419f9f95dad9bfbe4cff","run":"b8004e2989e24e1d8e4449afad2eb480","status":"completed"}

```