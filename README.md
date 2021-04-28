# Kubernetes Log Shipping

This demonstration has the purpose of presenting how you can easyly integrate EKS logs with Amazon Elasticsearch service and secure the access to the logs using Amazon Cognito.

## Architecture


<p align="center"> 
<img src="images/EKS-Log-Architecture.png">
</p>

**This demonstration was tested in us-east-1 region**

## Pre Reqs

- eksctl
- Helm 3+
- awscli
- kubectl

## Provision EKS cluster

We are going to use eksctl to provision our EKS cluster for this demonstration, eksctl will provision the Kubernetes master, nodes and **also the VPC**.

```shell
eksctl create cluster -f kubernetes-cluster-us-east-1.yaml
```

Now we have to configure our `~/.kube/config` file in order to access our cluster using kubectl.

```shell
aws eks --region us-east-1 update-kubeconfig --name kubelogs-cluster
```

This will update your kube config file, now let's test if we can access our cluster.

```shell
kubectl get nodes
```

### Get Public IP of EC2 EKS Instances

We will use this instances IP's to allow fluentbit to publish logs to our Elasticsearch cluster.

```shell
aws ec2 describe-instances --filter 'Name=tag:Name,Values=kubelogs-cluster*' --query 'Reservations[*].Instances[*].PublicIpAddress' --region us-east-1 | jq
```

The output will look like the following.

```json
[
  [
    "18.220.19.250"
  ],
  [
    "3.19.221.120"
  ],
  [
    "3.129.10.94"
  ]
]
```

## Creating our Amazon Elasticsearch cluster with Cognito integration

In this repository you will find a CloudFormation template that will create all the components that we need to integrate Amazon Elasticsearch with Amazon Cognito.

```shell
aws cloudformation create-stack \
    --stack-name kubelogs-es-stack \
    --template-body file://cludformation/stack.yaml \
    --capabilities CAPABILITY_IAM \
    --region us-east-1 \
    --parameters ParameterKey=EksInstancesIps,ParameterValue={IPS_SEPARETED_BY_COMMA}
```

## Creating Cognito User to Access Kibana Dashboard

## References

https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html
https://github.com/aws-samples/amazon-elasticsearch-service-with-cognito/