# scabarrus.com/k8s.webhook
This repository contains some projects that help to understand webhook plugin that cant be embeded in a K8S cluster to manage:

![alt text](https://github.com/scabarrus/scabarrus.com/blob/master/k8s_webhook.PNG)

4 Webhooks will be developed to provide understanding of how your K8S cluster can be customized:
* Authn webhook for user authentication
* Authz webhook to control permission
* mutating webhook to apply modify resource on fly
* validate webhook to check compliance rules of resources managed

An additional microservice called user-management provide Endpoint to register in postgres user, groups and roles.

![alt text](https://github.com/scabarrus/scabarrus.com/blob/master/user-management.PNG)

## Lab environment
My lab environment is a minimal K8S cluster with one Master and one Worker (poor cluster :-)).

## Module webhook k8s

This module contains three 4 main project:
- user-management 
- auhtn webhook 
- authz webhook (not yet)
- validate webhook (not yet)
- mutating webhook (not yet)

## User-management
This microservices provides endpoint to manage users and groups.
There is for now no authentication required because it's just as an example.

### Docker image 
To build the image from docker command:
```
# docker build -t user-management:1.0 -f scabarrus.com/k8s.webhook/deployment/Dockerfile . 
```

If you are on remote VM you can save this image with following command:
```
docker save -o user-management.tar user-management:1.0
```

If you want to load the image proviously on the worker (no registry for the moment):<br>
```
# docker load -i /tmp/user-management
352733f0fa4c: Loading layer [==================================================>] 2.293 MB/2.293 MB
a725aa4d4c34: Loading layer [==================================================>] 61.37 MB/61.37 MB
64bee900ead7: Loading layer [==================================================>] 34.24 MB/34.24 MB
Loaded image: user-management:1.0
```

### Deploy application
```
# kubectl apply -f scabarrus.com/k8s.webhook/deployment/usr-mgt/configmap.yaml
# kubectl apply -f scabarrus.com/k8s.webhook/deployment/usr-mgt/deployment.yaml
# kubectl expose deploy user-management-deploy --name user-management-svc  --type=NodePort
```

## Authn webhook
This webhook just query user-management microservice with bearer token sent when a user wants to interact with api-server.

### Build the image from source
```
# docker build -t authn-webhook:1.0 -f scabarrus.com/k8s.webhook/deployment/authn/Dockerfile .
```

### Save image
```
# docker save -o authn-webhook.tar authn-webhook:1.0
```

### Load image
```
# docker load -i authn-webhook.tar 
1fb22cc88687: Loading layer [==================================================>] 2.366 MB/2.366 MB
468ce57c4203: Loading layer [==================================================>] 127.1 MB/127.1 MB
8c1406fa31b0: Loading layer [==================================================>]  66.8 MB/66.8 MB
Loaded image: authn-webhook:1.0
```

### Deploy webhook
```
# kubectl apply -f scabarrus.com/k8s.webhook/deployment/authn/configmap.yaml
# kubectl apply -f scabarrus.com/k8s.webhook/deployment/authn/deployment.yaml
# kubectl expose deploy authn-webhook-svc --name authn-webhook  --type=NodePort
```
## Configure your cluster
```
kubeadm init --config scabarrus/k8s.webhook/configs/authn-config.yaml
```
You can check inside of kube-apiserver.yaml file that kube-apiserver will have some extra parameters in command section and hostpath volume mounted with the webhook configuration.
```
# grep authn /etc/kubernetes/manifests/kube-apiserver.yaml
    - --authentication-token-webhook-config-file=/etc/authn-config.yaml
    - mountPath: /etc/authn-config.yaml
      path: /root/authn-config.yaml

```
NB: when the kube-apiserver.yaml is modified, automatically api-server is redeployed.

## Create role and binding for user 
In this example, we will usee a user with login marvel and password.
```
kubectl create role admin --verb="*" --resource="*"
kubectl create rolebinding  marvel --role="admin" --user="marvel"
```

## Test authentication 
```
kubectl --kubeconfig scabarrus/k8s.webhook/configs/admin.conf --user marvel get pod
NAME                                      READY   STATUS    RESTARTS   AGE
authn-webhook-98d8b6fdf-7cw7v             1/1     Running   0          27m
pgadmin-8496757f7b-svg7q                  1/1     Running   1          12d
pod-webhook-pg                            1/1     Running   1          12d
user-management-deploy-57dcd4b4dc-ps9mm   1/1     Running   0          47h
```
