!/bin/bash

INVENTORY_FILE="$(pwd)/hosts.ini"


if [ ! -f "$INVENTORY_FILE" ]; then
    echo "Inventory file not found: $INVENTORY_FILE"
    exit 1
fi

if [ -z "$1" ]; then
    echo "Usage: $0 <ssh_key_path>"
    exit 1
fi

SSH_KEY_PATH=$1

run_playbook() {
    PLAYBOOK=$1
    EXTRA_VARS=$2
    echo "Running playbook: $PLAYBOOK"
    ansible-playbook -i "$INVENTORY_FILE" "$PLAYBOOK" $EXTRA_VARS
    if [ $? -ne 0 ]; then
        echo "Playbook $PLAYBOOK failed. Exiting."
        exit 1
    fi
}


run_playbook "./ansible-playbook/deploy_ssh_key.yaml" "-e ssh_key_path=${SSH_KEY_PATH}"


run_playbook "./ansible-playbook/create-new-user.yaml"
run_playbook "./ansible-playbook/setup-all-nodes.yaml"
run_playbook "./ansible-playbook/join-workers.yaml"
run_playbook "./ansible-playbook/nfs-service.yaml"
run_playbook "./ansible-playbook/workers-db.yaml"

echo "All playbooks have been executed successfully."



# echo "Creating new users..."
# ansible-playbook -i $INVENTORY_FILE ./copy_files.yaml 

# echo "All playbooks have been executed successfully."

# ./deploy_and_setup.sh /home/devadmin/.ssh/id_rsa.pub