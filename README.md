<p align="center">
<img alt="LittleHorse Logo" src="https://littlehorse.io/img/logo-wordmark-white.png" width="50%">
</p>

# LittleHorse Go QuickStart

- [LittleHorse Go QuickStart](#littlehorse-go-quickstart)
- [Prerequisites](#prerequisites)
  - [Setup Go](#setup-go)
  - [LittleHorse CLI](#littlehorse-cli)
  - [Local LH Server Setup](#local-lh-server-setup)
  - [Verifying Setup](#verifying-setup)
- [Running the Example](#running-the-example)
  - [Register Workflow](#register-workflow)
  - [Run Workflow](#run-workflow)
  - [Run Task Worker](#run-task-worker)
- [Next Steps](#next-steps)

**Get started in under 5 minutes, or your money back!** :wink:

This repo contains a minimal example to get you started using LittleHorse in Go. [LittleHorse](https://littlehorse.io) is a high-performance orchestration engine which lets you build workflow-driven microservice applications with ease.

You can run this example in two ways:

1. Using a local deployment of a LittleHorse Server (instructions below, requires one `docker` command).
2. Using a LittleHorse Server deployed in a cloud sandbox (to get one, contact `info@littlehorse.io`).

In this example, we will run a classic "Greeting" workflow as a quickstart. The workflow takes in one input variable (`input-name`), and calls a `greet` Task Function with the specified `input-name` as input.

# Prerequisites

Your system needs:

- `go`
- `brew` (to install `lhctl`) (Mac/Linux/WSL)
- `docker` (to run the LH Server) or access to a LH Cloud Sandbox.

## Setup Go

To add the LittleHorse Go Client to your project, you can use the following command:

```sh
go get github.com/littlehorse-enterprises/littlehorse
```

## LittleHorse CLI

Install the LittleHorse CLI:

```sh
brew install littlehorse-enterprises/lh/lhctl
```

## Local LH Server Setup

If you have obtained a private LH Cloud Sandbox, you can skip this step and just follow the configuration instructions you received from the LittleHorse Team (remember to set your environment variables!).

To run a LittleHorse Server locally in one command, you can run:

```sh
docker run --name littlehorse -d -p 2023:2023 -p 8080:8080 ghcr.io/littlehorse-enterprises/littlehorse/lh-standalone:0.11.2
```

Using the local LittleHorse Server takes about 15-25 seconds to start up, but it does not require any further configuration. Please note that the `lh-standalone` docker image requires at least 1.5GB of memory to function properly. This is because it runs kafka, the LH Server, and the LH Dashboard (2 JVM's and a NextJS app) all in one container.

## Verifying Setup

At this point, whether you are using a local Docker deployment or a private LH Cloud Sandbox, you should be able to contact the LH Server:

```sh
->lhctl version
lhctl version: 0.11.2 (Git SHA homebrew)
Server version: 0.11.2
```

**You should also be able to see the dashboard** at `https://localhost:8080`. It should be empty, but we will put some data in there soon when we run the workflow!

If you _can't_ get the above to work, please let us know on our [Community Slack Workspace](https://launchpass.com/littlehorsecommunity). We'll be happy to help.

# Running the Example

Without further ado, let's run the example start-to-finish.

## Register Workflow

First, we run `src/register`, which does two things:

1. Registers a `TaskDef` named `greet` with LittleHorse.
2. Registers a `WfSpec` named `quickstart` with LittleHorse.

A [`WfSpec`](https://littlehorse.dev/docs/concepts/workflows) specifies a process which can be orchestrated by LittleHorse. A [`TaskDef`](https://littlehorse.dev/docs/concepts/tasks) tells LittleHorse about a specification of a task that can be executed as a step in a `WfSpec`.

```sh
go run ./src/register_metadata.go
```

You can inspect your `WfSpec` with `lhctl` as follows. It's ok if the response doesn't make sense, we have a UI coming really soon which visualizes it for you!

```sh
lhctl get wfSpec quickstart
```

Now, go to your dashboard in your browser (`http://localhost:8080`) and refresh the page. Click on the `quickstart` WfSpec. You should see something that looks like a flow-chart. That is your Workflow Specification!

## Run Workflow

Now, let's run our first `WfRun`! Use `lhctl` to run an instance of our `WfSpec`.

```sh
# Run the 'quickstart' WfSpec, and set 'input-name' = "obi-wan"
lhctl run quickstart input-name obi-wan
```

The response prints the initial status of the `WfRun`. Pull out the `id` and copy it!

Let's look at our `WfRun` once again:

```sh
lhctl get wfRun <wf_run_id>
```

If you would like to see it on the dashboard, refresh the `WfSpec` page and scroll down. You should see your ID under the `RUNNING` column. Please double-click on your `WfRun` id, and it will take you to the `WfRun` page.

Note that the status is `RUNNING`! Why hasn't it completed? That's because we haven't yet started a worker which executes the `greet` tasks. Want to verify that? Let's search for all tasks in the queue which haven't been executed yet. You should see an entry whose `wfRunId` matches the Id from above:

```sh
lhctl search taskRun --taskDefName greet --status TASK_SCHEDULED
```

You can also see the `TaskRun` node on the workflow. It's highlighted, meaning that it's already running! If you click on it, you can see that it is in the `TASK_SCHEDULED` status.

## Run Task Worker

Now let's start our worker, so that our blocked `WfRun` can finish:

```sh
go run ./src/workers.go
```

Once the worker starts up, please open another terminal and inspect our `WfRun` again:

```sh
lhctl get wfRun <wf_run_id>
```

Voila! It's completed. You can also verify that the Task Queue is empty now that the Task Worker executed all of the tasks:

```sh
lhctl search taskRun --taskDefName greet --status TASK_SCHEDULED
```

Please refresh the dashboard, and you can see the `WfRun` has been completed!

# Next Steps

If you've made it this far, then it's time you become a full-fledged LittleHorse Knight!

Want to do more cool stuff with LittleHorse and Go? You can find more Go examples [here](https://github.com/littlehorse-enterprises/littlehorse/tree/master/sdk-go). This example only shows rudimentary features like tasks and variables. Some additional features not covered in this quickstart include:

- Conditionals
- Loops
- External Events (webhooks/signals etc)
- Interrupts
- User Tasks
- Multi-Threaded Workflows
- Workflow Exception Handling

We also have quickstarts in [Java](https://github.com/littlehorse-enterprises/lh-quickstart-java) and [Python](https://github.com/littlehorse-enterprises/lh-quickstart-python). Support for .NET is coming soon.

Our extensive [documentation](www.littlehorse.dev) explains LittleHorse concepts in detail and shows you how take full advantage of our system.

Our LittleHorse Server is free for production use under the SSPL license. You can find our official docker image at the [AWS ECR Public Gallery](https://gallery.ecr.aws/littlehorse/lh-server). If you would like enterprise support, or a managed service (either in the cloud or on-prem), contact `info@littlehorse.io`.

Happy riding!
