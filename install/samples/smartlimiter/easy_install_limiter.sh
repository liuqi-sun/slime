#!/bin/bash
kubectl create ns mesh-operator
kubectl apply -f https://raw.githubusercontent.com/slime-io/slime/v0.2.0-alpha/install/init/crds.yaml
kubectl apply -f https://raw.githubusercontent.com/slime-io/slime/v0.2.0-alpha/install/init/deployment_slime-boot.yaml
kubectl apply -f https://raw.githubusercontent.com/slime-io/slime/v0.2.0-alpha/install/samples/smartlimiter/slimeboot_smartlimiter.yaml