# External Adapter Service Template

This is meant to be an example for how you can set up an external adapter to make use of additional functions for a Chainlink node.

## Dependencies

gorilla/mux

```bash
go get github.com/gorilla/mux
```

## Building

```bash
$ go build -o ea src/*
```

## Running

```bash
$ ./ea
```

## Adding to Chainlink

Make sure your Chainlink node is already running. For example, mine is running at http://localhost:6688 and the adapter is running at http://localhost:3000.

### Add the adapter

```bash
curl -u $USERNAME:$PASSWORD -X POST -H 'Content-Type: application/json' -d '{"name":"ea","url":"http://localhost:3000/"}' http://localhost:6688/v2/bridge_types
```

### Create the JobSpec

```bash
curl -u $USERNAME:$PASSWORD -X POST -H 'Content-Type: application/json' -d '{"initiators":[{"type":"web"}],"tasks":[{"type":"ea"}]}' http://localhost:6688/v2/specs
```

## Using the InputData example

The InputData example allows for data to be passed into the adapter, and the adapter can use that to determine what to do. See the `GetInputData` function to see how that can work.

### Add the BridgeType

```bash
curl -u $USERNAME:$PASSWORD -X POST -H 'Content-Type: application/json' -d '{"name":"inputAdapter","url":"http://localhost:3000/input"}' http://localhost:6688/v2/bridge_types
```

### Create the JobSpec

```bash
curl -u $USERNAME:$PASSWORD -X POST -H 'Content-Type: application/json' -d '{"initiators":[{"type":"web"}],"tasks":[{"type":"inputAdapter"},{"type":"noop"}]}' http://localhost:6688/v2/specs
```

Take note of the "id" field that is returned after running this command.

### Starting a run

Be sure to change the JobID to the given output from the last command.

```bash
curl -u $USERNAME:$PASSWORD -X POST -H 'Content-Type: application/json' -d '{"other": "GetRestData"}' http://localhost:6688/v2/specs/8f7e26344a90473b82eb010a016a8ddd/runs
```

Replace "8f7e26344a90473b82eb010a016a8ddd" with the "id" from the previous command.

## Examples

Here is an example of what the log would look like for a run

```shell
2018/04/14 14:39:41 /input
2018/04/14 14:39:41 Input:
2018/04/14 14:39:41 {"id":"278c97ffadb54a5bbb93cfec5f7b5503","data":{"other":"GetRestData"}}
2018/04/14 14:39:41 Output:
2018/04/14 14:39:41 {"jobRunId":"278c97ffadb54a5bbb93cfec5f7b5503","data":{"value":"30000","last":"3333","other":"GetRestData"},"status":"completed","error":null,"pending":false}
```

And here is what the Chainlink node will log (with debug enabled)

```shell
{"level":"info","ts":1523734781.925597,"caller":"services/job_runner.go:79","msg":"Starting job","job":"8f7e26344a90473b82eb010a016a8ddd","run":"278c97ffadb54a5bbb93cfec5f7b5503","status":"in_progress"}
{"level":"debug","ts":1523734781.9286554,"caller":"services/job_runner.go:114","msg":"Produced task run","taskRun":"TaskRun(5ecc2bb67ace4b96975e03b138713203,inputadapter,completed,)"}
{"level":"debug","ts":1523734781.9290605,"caller":"services/job_runner.go:115","msg":"Task inputadapter ","task":0,"result":"","type":"inputadapter","params":"{\"other\":\"GetRestData\",\"type\":\"inputAdapter\"}","taskrun":"5ecc2bb67ace4b96975e03b138713203","status":""}
{"level":"debug","ts":1523734781.9338725,"caller":"services/job_runner.go:114","msg":"Produced task run","taskRun":"TaskRun(5c2ede41038348bd8d29d85eedd7e75f,noop,completed,)"}
{"level":"debug","ts":1523734781.9344666,"caller":"services/job_runner.go:115","msg":"Task noop ","task":1,"result":"","type":"noop","params":"{\"other\":\"GetRestData\",\"type\":\"noop\"}","taskrun":"5c2ede41038348bd8d29d85eedd7e75f","status":""}
{"level":"info","ts":1523734781.9385347,"caller":"services/job_runner.go:109","msg":"Finished current job run execution","job":"8f7e26344a90473b82eb010a016a8ddd","run":"278c97ffadb54a5bbb93cfec5f7b5503","status":"completed"}
```