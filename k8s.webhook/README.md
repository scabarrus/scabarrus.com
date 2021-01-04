# scabarrus.com/k8s.webhook
This repository contains some projects that help to understand webhook plugin that cant be embeded in a K8S cluster to manage:

![Project link](https://github.com/scabarrus/scabarrus.com/blob/master/k8s_webhook.PNG)

4 Webhooks will be developed to provide understanding of how your K8S cluster can be customized:
* Authn webhook for user authentication
* Authz webhook to control permission
* mutating webhook to apply modify resource on fly
* validate webhook to check compliance rules of resources managed

An additional microservice called user-management provide Endpoint to register in postgres user, groups and roles.



## Lab environment
My lab environment is a minimal K8S cluster with one Master and one Worker (poor cluster :-)).

## Module webhook k8s

This module contains three 4 main project:
- user-management 
- auhtn webhook 
- authz webhook (not yet)
- validate webhook (not yet)
- mutating webhook (not yet)


## Project layout

```
# tree -d
.
└── k8s.webhook
    ├── cmd
    │   ├── admission
    │   ├── authn
    │   ├── authz
    │   └── user-management
    │       └── docs
    ├── configs
    ├── deployment
    │   ├── authn
    │   └── user-mgt
    ├── docs
    ├── internal
    │   ├── domain
    │   ├── dto
    │   ├── error
    │   ├── repository
    │   ├── service
    │   └── utils
    └── middleware
```

### First layout level description
| Folder        | Description                                               |
| ------------- | ----------------------------------------------------------|
| cmd           | contain folder of each main application                   |
| configs       | contain all global config on K8S                          |
| deployment    | contain all yaml files required to deploy the application |
| internal      | contain all internal package that no need to be exported outside of the project|
| middleware    | contain middleware that act as decorator method for each http request (example: logging)|


## User-management
This microservices provides endpoint to manage users and groups.
There is for now no authentication required because it's just as an example.

A swagger is provided to play with api at following uri: /swagger/index.html

This microservice should be used only as example, because:
* No authentication is required to add user,group and role (it's very bad)
* The microservice is not in https 
* It was written for example and not for production
* It will be better to use oidc capabilities to manage authentication (perhaps in another repo)

For more details go to:

* ![User management documentation](https://github.com/scabarrus/scabarrus.com/blob/master/k8s.webhook/docs/user-management.md)

## Authn webhook
This webhook just query user-management microservice with bearer token sent when a user wants to interact with api-server.

* ![Authn webhook documentation](https://github.com/scabarrus/scabarrus.com/blob/master/k8s.webhook/docs/authn-webhook.md)
