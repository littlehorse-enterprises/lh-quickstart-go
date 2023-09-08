<p align="center">
<img alt="LH" src="https://littlehorse.dev/img/logo.jpg" width="50%">
</p>


## Quickstart for Go

### Prerequisites

- `docker`.
- `go`.

### Running Locally

Install `lhctl`:

```
go install github.com/littlehorse-enterprises/littlehorse/lhctl@latest
```

Verify the installation:

```
lhctl
```

Start a LH Server with:

```
docker run --name littlehorse -d -p 2023:2023 public.ecr.aws/littlehorse/lh-standalone:latest
```

When you run the LH Server according to the command above, the API Host is `localhost` and the API Port is `2023`.
Now configure `~/.config/littlehorse.config`:

```
LHC_API_HOST=localhost
LHC_API_PORT=2023
```

You can confirm that the Server is running via:

```
lhctl search wfSpec
```

Result:

```
{
  "results": []
}
```

Now let's run the example

Deploy the `TaskDef` and run the TaskWorker:

```
go run ./src/worker
```

Now, in a separate terminal, register the WfSpec:

```
go run ./src/deploy
```


In another terminal, use `lhctl` to run the workflow:

```
# Here, we specify that the "name" variable = "Obi-Wan"
lhctl run my-workflow name Obi-Wan
```

The workflow can also be run using the sdk:

```
go run ./src/runworkflow
```

Now let's inspect the result:

```
# This call shows the workflow specification
lhctl get wfSpec my-workflow

# This call shows the result
lhctl get wfRun <wf run id>

# This will show you all nodes in tha run
lhctl get nodeRun <wf run id> 0 1

# This shows the task run information
lhctl get taskRun <wf run id> <task run global id>
```

> More Go examples [here](https://github.com/littlehorse-enterprises/littlehorse/tree/master/sdk-go/examples).
