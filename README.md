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
To build the image from docker command:<br>
`# docker build -t user-management:1.0 -f scabarrus.com/k8s.webhook/deployment/Dockerfile . `

If you are on remote VM you can save this image with following command:
`docker save -o user-management.tar user-management:1.0`

If you want to load the image proviously on the worker (no registry for the moment):<br>
```
# docker load -i /tmp/user-management
352733f0fa4c: Loading layer [==================================================>] 2.293 MB/2.293 MB
a725aa4d4c34: Loading layer [==================================================>] 61.37 MB/61.37 MB
64bee900ead7: Loading layer [==================================================>] 34.24 MB/34.24 MB
Loaded image: user-management:1.0
```

