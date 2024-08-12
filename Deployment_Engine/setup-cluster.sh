#!/bin/bash


INVENTORY_FILE="$(pwd)"


if [ ! -f "$INVENTORY_FILE" ]; then
    echo "Inventory file not found: $INVENTORY_FILE"
    exit 1
fi


echo "Creating new users..."
ansible-playbook -i $INVENTORY_FILE ./ansible-playbook/create-new-user.yaml 

echo "Setting up all nodes..."
ansible-playbook -i $INVENTORY_FILE ./ansible-playbook/setup-all-nodes.yaml 

echo "Joining worker nodes to the cluster..."
ansible-playbook -i $INVENTORY_FILE ./ansible-playbook/join-workers.yaml 

echo "Setting up NFS services..."
ansible-playbook -i $INVENTORY_FILE ./ansible-playbook/nfs-service.yaml 

echo "Transferring folders to worker nodes..."
ansible-playbook -i $INVENTORY_FILE ./ansible-playbook/workers-db.yaml 

echo "All playbooks have been executed."
