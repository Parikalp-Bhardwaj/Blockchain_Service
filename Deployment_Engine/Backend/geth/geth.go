package geth

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	// "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/rest"
)

func CreateGethStatefulSet(config *rest.Config, namespace string, chainId int) error {
	statefulSetName := "geth-statefulset" // Replace with your desired StatefulSet name

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	filePath := "/mnt/shared-files/Node-1-address.txt"

	// Read the file
	mydata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Convert the file data to a string
	address := strings.TrimSpace(string(mydata))
	fmt.Println("address ", address)
	fmt.Println("ChainId ", chainId)

	// Define the StatefulSet
	statefulSetClient := clientset.AppsV1().StatefulSets(namespace)

	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: statefulSetName,
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "geth-execution-service", // Use the headless service you've created
			Replicas:    int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "crystal-mev",
					"comp": "geth-node-pod",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "crystal-mev",
						"comp": "geth-node-pod",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "geth-node-pod",
							Image: "digifigroup/geth:500m",
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/Node-1/",
									SubPath:   "Node-1",
								},
							},
							Ports: []v1.ContainerPort{
								{ContainerPort: 8551},
								{ContainerPort: 8545},
								{ContainerPort: 8546},
								{ContainerPort: 30303},
								{ContainerPort: 30304},
							},

							Args: []string{
								"--syncmode=full",
								"--datadir=/Node-1/execution/geth",
								"--authrpc.addr=0.0.0.0",
								"--authrpc.port=8551",
								"--authrpc.vhosts=*",
								"--authrpc.jwtsecret=/Node-1/execution/secrets/jwtsecret",
								"--ws",
								"--ws.addr=0.0.0.0",
								"--ws.api=txpool,eth,net,web3,personal",
								"--ws.origins=*",
								"--http",
								"--http.addr=0.0.0.0",
								"--http.api=eth,net,web3,admin,personal,admin,debug",
								"--authrpc.vhosts=*",
								"--http.corsdomain=*",
								"--password=/Node-1/execution/geth_password.txt",
								// "--nodiscover",
								"--datadir=/Node-1/execution",
								"--allow-insecure-unlock",
								"--ws.port=8546",
								"--unlock=0x" + address,
								"--networkid=" + strconv.Itoa(chainId),

								//Change bypass default cap?
								//  "--rpc.txfeecap=0",

							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker1",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfiles",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker1", // Replace with your PVC name
								},
							},
						},
					},
					// ... Other pod spec configuration ...
				},
			},
		},
	}

	// Create the StatefulSet
	createdStatefulSet, err := statefulSetClient.Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("StatefulSet '%s' created in namespace '%s'\n", createdStatefulSet.Name, namespace)
	return nil

}

// Server 2

func CreateGethStatefulSet2(config *rest.Config, namespace string, chainId int) error {
	statefulSetName := "geth-statefulset2" // Replace with your desired StatefulSet name

	filePath := "/mnt/shared-files/Node-2-address.txt"

	// Read the file
	mydata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Fetch Enode

	cmd := exec.Command("bash", "./scripts/enode_script.sh")
	node, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %v\n", err)
	}

	enode := strings.TrimSpace(string(node))
	fmt.Println(enode)

	enodeOutput2 := strings.Replace(enode, "127.0.0.1:30303", "192.168.253.108:30303", 1)
	enodeOutputClear := strings.TrimSpace(string(enodeOutput2))

	fmt.Println("geth enode 2 \n", enodeOutputClear)

	// Convert the file data to a string
	address2 := strings.TrimSpace(string(mydata))
	fmt.Println("address2 ", address2)

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Define the StatefulSet
	statefulSetClient2 := clientset.AppsV1().StatefulSets(namespace)

	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: statefulSetName,
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "geth-execution-service2", // Use the headless service you've created
			Replicas:    int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "crystal-mev2",
					"comp": "geth-node-pod2",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "crystal-mev2",
						"comp": "geth-node-pod2",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "geth-node-pod2",
							Image: "digifigroup/geth:500m",
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/Node-2/",
									SubPath:   "Node-2",
								},
							},
							Ports: []v1.ContainerPort{
								{ContainerPort: 8551},
								{ContainerPort: 8545},
								{ContainerPort: 8546},
								{ContainerPort: 30303},
								{ContainerPort: 30304},
							},
							Args: []string{
								"--syncmode=full",
								"--datadir=/Node-2/execution/geth",
								// "--ipcdisable",
								"--authrpc.addr=0.0.0.0",
								"--authrpc.port=8551",
								"--authrpc.vhosts=*",
								"--authrpc.jwtsecret=/Node-2/execution/secrets/jwtsecret",
								"--ws",
								"--ws.addr=0.0.0.0",
								"--ws.api=txpool,eth,net,web3,personal",
								"--ws.origins=*",
								"--http",
								"--http.addr=0.0.0.0",
								"--http.api=txpool,eth,net,web3,personal,admin,debug",
								"--http.corsdomain=*",
								"--password=/Node-2/execution/geth_password.txt",
								// "--nodiscover",
								"--http.port=8545",
								"--http.corsdomain=*",
								"--ws.port=8546",
								"--datadir=/Node-2/execution",
								"--allow-insecure-unlock",
								"--networkid=" + strconv.Itoa(chainId),
								"--unlock=0x" + address2,
								"--bootnodes=" + enodeOutputClear,
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker2",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfiles",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker2", // Replace with your PVC name
								},
							},
						},
					},
					// ... Other pod spec configuration ...
				},
			},
		},
	}

	// Create the StatefulSet
	createdStatefulSet, err := statefulSetClient2.Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("StatefulSet '%s' created in namespace '%s'\n", createdStatefulSet.Name, namespace)
	return nil

}

// Server 3

func CreateGethStatefulSet3(config *rest.Config, namespace string, chainId int) error {
	statefulSetName := "geth-statefulset3" // Replace with your desired StatefulSet name

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	filePath := "/mnt/shared-files/Node-3-address.txt"

	// Read the file
	mydata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Fetch Enode

	cmd := exec.Command("bash", "./scripts/enode_script.sh")
	node, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %v\n", err)
	}

	enode := strings.TrimSpace(string(node))
	fmt.Println(enode)

	enodeOutput2 := strings.Replace(enode, "127.0.0.1:30303", "192.168.253.108:30303", 1)
	enodeOutputClear := strings.TrimSpace(string(enodeOutput2))

	fmt.Println("geth enode 2 \n", enodeOutputClear)

	// Convert the file data to a string
	address3 := strings.TrimSpace(string(mydata))
	fmt.Println("address3 ", address3)

	// Define the StatefulSet
	statefulSetClient3 := clientset.AppsV1().StatefulSets(namespace)

	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: statefulSetName,
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "geth-execution-service3", // Use the headless service you've created
			Replicas:    int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "crystal-mev3",
					"comp": "geth-node-pod3",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "crystal-mev3",
						"comp": "geth-node-pod3",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "geth-node-pod3",
							Image: "digifigroup/geth:500m",
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/Node-3/",
									SubPath:   "Node-3",
								},
							},
							Args: []string{
								"--syncmode=full",
								"--datadir=/Node-3/execution/geth",
								// "--ipcdisable",
								"--authrpc.addr=0.0.0.0",
								"--authrpc.port=8551",
								"--authrpc.vhosts=*",
								"--authrpc.jwtsecret=/Node-3/execution/secrets/jwtsecret",
								"--ws",
								"--ws.addr=0.0.0.0",
								"--ws.api=txpool,eth,net,web3,personal",
								"--ws.origins=*",
								"--http",
								"--http.addr=0.0.0.0",
								"--http.api=txpool,eth,net,web3,personal,admin,debug",
								"--http.corsdomain=*",
								"--password=/Node-3/execution/geth_password.txt",
								// "--nodiscover",
								"--http.port=8545",
								"--http.corsdomain=*",
								"--ws.port=8546",
								"--datadir=/Node-3/execution",
								"--allow-insecure-unlock",
								"--networkid=" + strconv.Itoa(chainId),
								"--unlock=0x" + address3,
								"--bootnodes=" + enodeOutputClear,
								//Change bypass default cap?
								// "--rpc.txfeecap=0",
							},
							// ... Other container configuration ...
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker3",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfiles",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker3", // Replace with your PVC name
								},
							},
						},
					},
					// ... Other pod spec configuration ...
				},
			},
		},
	}

	// Create the StatefulSet
	createdStatefulSet, err := statefulSetClient3.Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("StatefulSet '%s' created in namespace '%s'\n", createdStatefulSet.Name, namespace)
	return nil

}

// Server 4

func CreateGethStatefulSet4(config *rest.Config, namespace string) error {
	statefulSetName := "geth-statefulset4" // Replace with your desired StatefulSet name

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	filePath := "/mnt/shared-files/Node-4-address.txt"

	// Read the file
	mydata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	scriptPath := "./beaconnode/enode_script.sh"

	// Check if the script file exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Printf("Script file not found: %s\n", scriptPath)
		return err
	}

	// Run the script
	cmd := exec.Command(scriptPath)

	// Capture the script's output
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running script with sudo: %v\n", err.Error())
		return err
	}

	// Prepend "enode://" to the output
	enodeOutput := "enode://" + string(output) + "@192.168.253.108:30303"
	enodeOutput = strings.Replace(enodeOutput, "\n", "", 1)

	fmt.Println("geth enode 2", enodeOutput)

	// Convert the file data to a string
	address := strings.TrimSpace(string(mydata))
	fmt.Println("Address Node-4", address)

	// Define the StatefulSet
	statefulSetClient := clientset.AppsV1().StatefulSets(namespace)

	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: statefulSetName,
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "geth-execution-service4", // Use the headless service you've created
			Replicas:    int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "crystal-mev4",
					"comp": "geth-node-pod4",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "crystal-mev4",
						"comp": "geth-node-pod4",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "geth-node-pod4",
							Image: "ethereum/client-go:latest",
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/Node-4/",
									SubPath:   "Node-4/",
								},
							},
							Args: []string{
								"--syncmode=full",
								"--datadir=/Node-4/execution/geth",
								"--authrpc.addr=0.0.0.0",
								"--authrpc.port=8551",
								"--bootnodes=" + enodeOutput,
								"--authrpc.vhosts=*",
								"--authrpc.jwtsecret=/Node-1/execution/secrets/jwtsecret",
								"--ws",
								"--ws.addr=0.0.0.0",
								"--ws.api=txpool,eth,net,web3,personal",
								"--ws.origins=*",
								"--http",
								"--http.addr=0.0.0.0",
								"--http.api=txpool,eth,net,web3,personal,admin",
								"--http.corsdomain=*",
								"--http.vhosts=*",
								"--datadir=/Node-4/execution",
								"--allow-insecure-unlock",
								"--unlock=0x" + address,
								"--password=/Node-4/execution/geth_password.txt",

								//Change bypass default cap?
								// "--rpc.txfeecap=0",
							},
							// ... Other container configuration ...
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker4",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfiles",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker4", // Replace with your PVC name
								},
							},
						},
					},
					// ... Other pod spec configuration ...
				},
			},
		},
	}

	// Create the StatefulSet
	createdStatefulSet, err := statefulSetClient.Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("StatefulSet '%s' created in namespace '%s'\n", createdStatefulSet.Name, namespace)
	return nil

}

func int32Ptr(i int32) *int32 { return &i }
