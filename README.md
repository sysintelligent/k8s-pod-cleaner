# k8s-pod-cleaner
This repository implements a simple operator to delete the pods not in running status using Kubebuilder with the help of ChatGPT.

## Description
Kubernetes operators are software extensions that automate the deployment, management, and scaling of applications on a Kubernetes cluster. Developing a Kubernetes operator can be a complex task, but with the help of ChatGPT, you can streamline the process.

To begin developing a Kubernetes operator, you first need to understand the application you want to deploy and manage. This involves identifying the key components and dependencies, as well as the various parameters and configurations required for proper functioning.

Next, you can use ChatGPT to help you write the code for your operator. ChatGPT can assist you in writing the necessary Kubernetes manifests, as well as any custom code needed to handle specific tasks.

Once the code is written, you can use ChatGPT to test and debug your operator. ChatGPT can help you identify and fix any issues with the operator, as well as optimize it for performance and efficiency.

Overall, developing a Kubernetes operator using ChatGPT can significantly streamline the process, making it easier to deploy and manage complex applications on a Kubernetes cluster.

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [Minikube](https://github.com/kubernetes/minikube) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/k8s-pod-cleaner:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/k8s-pod-cleaner:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

## Running

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)