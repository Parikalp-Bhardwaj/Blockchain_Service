package ansibleconfig

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	model "github.com/deployment_engine/model"
)

type Inventory struct {
	Hosts []model.Host `json:"inventory"`
}

func SaveInventory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var inventory Inventory

	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	inventoryData := formatAnsibleInventory(inventory.Hosts)
	fmt.Println("Generated inventory data:\n", inventoryData)

	if inventoryData == "" {
		http.Error(w, "Generated inventory data is empty", http.StatusInternalServerError)
		return
	}

	workingDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Failed to get current working directory", http.StatusInternalServerError)
		return
	}

	parentDir := filepath.Dir(workingDir)

	filePath := filepath.Join(parentDir, "hosts.ini")

	err = writeToFile(filePath, inventoryData)
	if err != nil {
		http.Error(w, "Failed to write inventory to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Inventory saved successfully in parent directory."))
}

func formatAnsibleInventory(hosts []model.Host) string {
	groups := map[string][]string{}
	for _, host := range hosts {
		fmt.Printf("Processing host: %+v\n", host)

		hostEntry := fmt.Sprintf(
			"%s ansible_host=%s ansible_port=%s ansible_user=%s ansible_ssh_private_key_file=%s ansible_become_password='%s' ansible_ssh_pass='%s'",
			host.Name, host.IP, host.Port, host.AnsibleUser, host.AnsibleSSHKey, host.Password, host.Password)

		groups[host.Type] = append(groups[host.Type], hostEntry)
	}

	var inventoryBuilder strings.Builder
	for group, hosts := range groups {
		inventoryBuilder.WriteString(fmt.Sprintf("[%s]\n", group))
		for _, host := range hosts {
			inventoryBuilder.WriteString(fmt.Sprintf("%s\n", host))
		}
		inventoryBuilder.WriteString("\n")
	}

	inventoryBuilder.WriteString("[all:vars]\n")
	inventoryBuilder.WriteString("ansible_ssh_common_args='-o StrictHostKeyChecking=no'\n")
	inventoryBuilder.WriteString("ansible_become=yes\n")

	inventoryData := inventoryBuilder.String()

	return inventoryData
}

func writeToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("Successfully wrote data to file:", filename)
	return nil
}

func Runscript(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	scriptPath := "../setup-cluster.sh"

	cmd := exec.Command("/bin/bash", scriptPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute script: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
