# elastic-cloud-autoscaler-demo
Demo of https://github.com/k-yomo/elastic-cloud-autoscaler

⚠️ Running the demo will cost money.

## Prerequisites
- Go v.1.19.x
- Terraform v1.3.6
- Elastic Cloud subscription

## Setup
Create Elastic Cloud's API key and set it to env variable.
You can create API key in the below page.
https://cloud.elastic.co/deployment-features/keys
```shell
$ export EC_API_KEY=
```

2. Run `setup` command.  It'll create required deployments and elasticsearch index.
```shell
$ make setup
```

## Usage
```shell
$ make run
```

It'll check if scaling is needed and apply the change if any per minute based on the given configuration.
When scaling is applied, it'll print scaling operation like below.
```
==============================================
scaling direction: SCALING_OUT
topology size updated from: 1024 => to 65536
replica num updated from: 1 => to 0
reason: current or desired topology size '1g' is less than min topology size '64g'
==============================================
```

## Cleanup
`cleanup` command will destroy created resources
```shell
$ make cleanup
```
