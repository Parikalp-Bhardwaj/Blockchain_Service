#!/bin/bash

# Function to delete Kubernetes resources
delete_k8s_resources() {
  resource_type="$1"
  resources=$(kubectl get "$resource_type" -o custom-columns=NAME:.metadata.name --no-headers)

  if [ -n "$resources" ]; then
    for resource in $resources; do
      kubectl delete "$resource_type" "$resource"
      echo "Deleted $resource_type: $resource"
    done
    echo "${resource_type^} deletion completed."
  else
    echo "No $resource_type found."
  fi
}


# Delete pods
delete_k8s_resources "pods"
sleep 2

# Delete services
delete_k8s_resources "svc"
sleep 2

# Delete StatefulSets
delete_k8s_resources "statefulset"
sleep 2

# Delete Jobs
delete_k8s_resources "job"
sleep 2





# Change directory and run commands
# cd /mnt/nfs_clientshare/
# PASSWORD="Sd{zM30i*\01U)aJg;;(#{6="

# Run a command with sudo and provide the password
# echo "$PASSWORD" | sudo ./restart_node.sh

# Remove a directory
# rm -rf ./shared-files/nodekey
# rm -rf ./shared-files/nodekey
# rm -rf ./shared-files/Node-1-address.txt
# rm -rf ./shared-files/Node-2-address.txt
# rm -rf ./shared-files/Node-3-address.txt
sleep 2

# Apply Kubernetes resources
kubectl apply -f /home/devadmin/K8s-ETH2-Prysm-Validator/nfs/.

echo "Script execution completed."