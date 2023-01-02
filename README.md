# elastic-cloud-autoscaler-demo
Demo of https://github.com/k-yomo/elastic-cloud-autoscaler

⚠️ Running the demo will cost money.

## Prerequisites
- Elastic Cloud subscription
- Go v.1.19.x
- Terraform v1.3.6
- [k6](https://k6.io/docs/get-started/installation/)
- [direnv](https://direnv.net/)
- [Git LFS](https://git-lfs.com/)

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
3. Run `index-products` command.  It'll index 1.7M documents. The data [Shopping Queries Dataset](https://github.com/amazon-science/esci-data) published from Amazon.
```shell
$ make index-products
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

To increase CPU util, `make loadtest` command will send search requests using k6.
It'll gradually increase traffic and then decrease, so you can see the scale-out and scale-in.
```shell
$ make loadtest
```

## Cleanup
`cleanup` command will destroy created resources
```shell
$ make cleanup
```
