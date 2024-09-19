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


# Delete Storage Classes
delete_k8s_resources "sc"
delete_k8s_resources "job"

kubectl delete job delete-files

kubectl delete pvc data-worker1
kubectl delete pv local-pv-worker1

kubectl delete pvc data-worker2
kubectl delete pv local-pv-worker2

kubectl delete pvc data-worker3
kubectl delete pv local-pv-worker3


kubectl delete pvc data-worker4
kubectl delete pv local-pv-worker4