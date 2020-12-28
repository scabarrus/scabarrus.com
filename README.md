# scabarrus.com
This repository contains some projects that help to understand webhook plugin that cant be embeded in a K8S cluster to manage:
- authentication
- authorization
- validation
- mutating

## Lab environment
My lab environment is a minimal K8S cluster with one Master and one Worker (poor cluster :-)).

## Module webhook k8s

This module contains three 4 main project:
- user-management 
- auhtn webhook (not yet)
- authz webhook (not yet)
- validate webhook (not yet)
- mutating webhook (not yet)

## User-managemnt
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
kubectl apply -f configmap.yaml
kubectl apply -f deployment.yaml
kubectl expose deploy user-management-deploy --name user-management-svc  --type=NodePort
```

## Authn webhook
This webhook just query user-management microservice with bearer token sent when a user wants to interact with api-server.

### Build the image from source
```
docker build -t authn-webhook:1.0 -f scabarrus.com/k8s.webhook/deployment/authn/Dockerfile .
```


