# Kubernetes Log Shipping

This demonstration has the purpose of presenting how you can easyly integrate EKS logs with Amazon Elasticsearch service and secure the access to the logs using Amazon Cognito.

## Architecture


<p align="center"> 
<img src="images/EKS-Log-Architecture.png">
</p>

## Pre Reqs

- eksctl
- Helm 3+

## Provision EKS cluster

We are going to use eksctl to provision our EKS cluster for this demonstration, eksctl will provision the Kubernetes master, nodes and also the VPC.


## References

https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html
https://github.com/aws-samples/amazon-elasticsearch-service-with-cognito/