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