# On-Premise-Deployments

> Tiger: A tool to manage and deploy enterprise applications to k8s clusters

Tiger is a simple command line tool developed in Golang using [cobra](https://github.com/spf13/cobra) and [kubernetes/client-go](https://github.com/kubernetes/client-go). Tiger has the ability to add/manage Kubernetes cluster configs. It also provides the functionality to deploy enterprise application images on separately managed Kubernetes clusters. These Kubernetes clusters may be managed by enterprise-clients or the enterprise itself.

## Development

### Requirements

Tiger has been developed on Go-1.13.4. Instruction and binaries to install it can be found [here](https://golang.org/doc/install). Dependencies have been vendored. Tiger requires `GO111MODULES` to be set to `on`. Do so by, `export GO111MODULES=on`. For Kubernetes deployment testing, [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) has been used. The Kubernetes cli tool, [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), has been used for additional debugging.

### Quick Setup

1. Install Go, Minikube and Kubectl.
2. Set `export GO111MODULES=on`.
3. Clone project.
3. Start minikube cluster by, `minikube start`
4. Create a symlink of minikube kubeconfig in the `contexts` directory of the project. `ln -s $HOME/.kube/config $GOPATH/src/github.com/ankitkataria/on-premise-deployments/contexts/minikube-prod.yml`. This is the default configuration used when no `cluster-config` has been provided.
5. Build binary, `make clean && make build`. The binary will be available in `bin` directory.
6. Optionally, add `bin` to current $PATH, `export PATH=$PATH:$GOPATH/github.com/ankitkataria/on-premise-deployments/bin`.


### Usage

1. Tiger help.

```bash
$ tiger help

Tiger: On Premise Deployments

Usage:
  tiger [flags]
  tiger [command]

Available Commands:
  add         Add and configure new clients
  deploy      Manage applications on required client clusters
  help        Help about any command
  version     Displays the version of the current Tiger build

Flags:
  -h, --help   help for tiger

Use "tiger [command] --help" for more information about a command.
```

2. Add a new Kubernetes config

```bash
$ tiger add --clusters=cluster1,cluster2 --users=dev,prod --contexts=frontend,backend enterprise-client
#Configuration will be available to manage at ./contexts/enterprise-client-prod.yml

$ pwd && vim ./contexts/enterprise-client-prod.yml
# Sample output should look like:
# /home/kataria/go/src/github.com/ankitkataria/on-premise-deployments
#./contexts/minikube-example.yml
```

3. Creating a deployment on Kubernetes cluster

```bash
$ tiger deploy create --image=<image-name> --replicas=<replica-number> --cluster-config=./contexts/enterprise-client-prod.yml test-deployment
# if flag is not present, --cluster-config defaults to ./contexts/minikube-prod.yml for create, update and delete

# Verify on minikube using - 
$ kubectl get pods
```

4. Updating deployments on a cluster

```bash
$ tiger deploy update --image=<updated-image> --replicas=<replica-number> ---cluster-config=./contexts/enterprise-client-prod.yml test-deployment

# Verify on minikube using -
$ kubectl describe deployment test-deployment
```

5. Deleting deployments on a cluster

```bash
$ tiger deploy delete --cluster-config=./contexts/enterprise-client-prod.yml test-deployment

# Verify on minikube using -
$ kubectl get deployments
```

### Project structure

- `cmd`
	- contains `tiger` subcommands `add`, `deploy`, `version`
	- `deploy`
		- contains `deploy` subcommands `create`, `update`, `delete`


- `pkg`
	- `utils`
		- `ops` - contains k8s client-go functions to handle creatation, updations and deletion of deployments
		- `config` - contains go-interface mappings for k8s configs

- `contexts`
	- contains added k8s config files

### Future Work

- [ ] Make Tiger more flexible in handling Kubernetes configurations
- [ ] Add [Prometheus](https://prometheus.io/) for Metrics

## License

This project is under MIT License. 
