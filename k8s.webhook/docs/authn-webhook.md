# Authn webhook
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

Create certificate if you want to use an autosigned:

```

```
```
# kubectl create configmap authn-confmap --from-file=certificate.crt --from-file=privateKey.key
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
kubectl --kubeconfig scabarrus.com/k8s.webhook/configs/admin.conf --user marvel get pod
NAME                                      READY   STATUS    RESTARTS   AGE
authn-webhook-98d8b6fdf-7cw7v             1/1     Running   0          27m
pgadmin-8496757f7b-svg7q                  1/1     Running   1          12d
pod-webhook-pg                            1/1     Running   1          12d
user-management-deploy-57dcd4b4dc-ps9mm   1/1     Running   0          47h
```
