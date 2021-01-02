# User management

The following schema provides an overview of the microservice:
![Project link](https://github.com/scabarrus/scabarrus.com/blob/master/k8s.webhook/cmd/user-management/user-management.PNG)

Details:
* A swagger is exposed on following URI: /swagger/index.html
* No authenticaiton is required for now (it's not a production use case)
* 3 resources are managed (user, group and role)
* 1 Endpoint provide health check which just check the connectivity with it's database on URI /api/v1/healtz
* 1 Endpoint provide database initialization /api/admin/db-migrate
