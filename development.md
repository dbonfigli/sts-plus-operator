## How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster.

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html).


## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.

**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

**Note:** Run `make --help` for more information on all potential `make` targets.

### Running on the cluster

1. Deploy the requirements for tests in the cluster, i.e. cert-manager and the prometheus-operator:
```sh
make deploy-test-requirements
```

2. Build and push your image to the location specified by `IMG`:

```sh
IMAGE=<some-registry>/sts-phased-rollout-operator:tag
make docker-build docker-push IMG=$IMAGE
```
Or, if you use kind you can push the image directly to the cluster with:
```sh
IMAGE=<some-registry>/sts-phased-rollout-operator:tag
make docker-build IMG=$IMAGE
kind load docker-image $IMAGE
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=$IMAGE
```

4. Install Instances of Custom Resources:

```sh
kubectl apply -k config/samples/
```

### Running locally

Test it without running the operator in the cluster:

1. Deploy the requirements for tests in the cluster, i.e. cert-manager and the prometheus-operator:
```sh
make deploy-test-requirements
```

2. Port forward the prometheus endpoint so that it is reachable from your local development:
```sh
kubectl port-forward service/prometheus-prometheus -n monitoring 9090:9090
```

3. Install the CRDs into the cluster:

```sh
make install
```

4. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
# disable webhooks, otherwise you will need valid certs files ca.crt, tls.crt and tls.key in the /tmp/k8s-webhook-server/serving-certs/ directory
export ENABLE_WEBHOOKS=false
make run #or `make run-debug` for debug logs
```

5. Run the sample:
```sh
kubectl apply -k config/samples/local-development
```


### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller from the cluster:

```sh
make undeploy
```
