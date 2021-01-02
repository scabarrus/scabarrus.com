# User management

The following schema provides an overview of the microservice:
![Project link](https://github.com/scabarrus/scabarrus.com/blob/master/k8s.webhook/docs/user-management.PNG)

Details:
* A swagger is exposed on following URI: /swagger/index.html
* No authenticaiton is required for now (it's not a production use case)
* 3 resources are managed (user, group and role)
* 1 Endpoint provide health check which just check the connectivity with it's database on URI /api/v1/healtz
* 1 Endpoint provide database initialization /api/admin/db-migrate

## Deployment
### Build image 
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

### Deploy resource in k8s
```
# kubectl apply -f scabarrus.com/k8s.webhook/deployment/usr-mgt/configmap.yaml
# kubectl apply -f scabarrus.com/k8s.webhook/deployment/usr-mgt/deployment.yaml
# kubectl expose deploy user-management-deploy --name user-management-svc  --type=NodePort
```
NB: in my example, I don't have ingress then it's why I use node port.

You can access to the swagger by following url:
http://{worker ip address}:{node port}/swagger/index.html

```

You can check log of the pod and see that at least each request are trace.
# kubectl logs -f user-management-deploy-{pod indentifier}
2021/01/02 00:28:03 GET   192.168.169.129:31958   /swagger/index.html   192.168.169.129:38063
2021/01/02 00:28:03 GET   192.168.169.129:31958   /swagger/swagger-ui.css   192.168.169.129:38063
2021/01/02 00:28:03 GET   192.168.169.129:31958   /swagger/swagger-ui-bundle.js   192.168.169.129:60197
2021/01/02 00:28:03 GET   192.168.169.129:31958   /swagger/swagger-ui-standalone-preset.js   192.168.169.129:28613
2021/01/02 00:28:03 GET   192.168.169.129:31958   /swagger/doc.json   192.168.169.129:60197
2021/01/02 00:28:03 GET   192.168.169.129:31958   /swagger/favicon-32x32.png   192.168.169.129:60197
```
