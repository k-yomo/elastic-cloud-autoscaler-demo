# elastic-cloud-autoscaler-demo
Demo of https://github.com/k-yomo/elastic-cloud-autoscaler

⚠️ Running the demo will cost money.

## Prerequisites
- Elastic Cloud subscription
- Go v.1.19.x
- Terraform v1.3.6
- [k6](https://k6.io/docs/get-started/installation/)
- [direnv](https://direnv.net/)

## Setup
Create Elastic Cloud's API key and set it to env variable.
You can create API key in the below page.
https://cloud.elastic.co/deployment-features/keys
```shell
$ export EC_API_KEY=
```

2. Run `setup` command.  It'll create required deployments and elasticsearch index with 1.7M documents.
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
topology size updated from: 65536 => to 131072
replica num updated from: 1 => to 3
reason: current or desired topology size '64g' is less than min topology size '128g'
==============================================
```

## Cleanup
`cleanup` command will destroy created resources
```shell
$ make cleanup
```
