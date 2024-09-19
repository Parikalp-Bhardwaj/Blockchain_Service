package beaconchain

import (
	"context"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	// "k8s.io/apimachinery/pkg/api/resource"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/rest"
)

func BeaconNode(config *rest.Config, namespace string, chainId int) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Define the StatefulSet
	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-statefulset", // Update with your desired name
			Labels: map[string]string{
				"app": "lighthousebeacon",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "lighthousebeacon-service", // Update with your desired service name
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "lighthousebeacon",
					"comp": "lighthousebeacon-svc",
				},
			},
			Replicas: int32Ptr(1), // Set the number of replicas as needed
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "lighthousebeacon",
						"comp": "lighthousebeacon-svc",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "lighthousebeacon",
							Image: "sigp/lighthouse:v4.5.0-amd64-dev",
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "lighthousebeacon",
									MountPath: "/Node-1/",
									SubPath:   "Node-1",
								},
							},
							Ports: []v1.ContainerPort{
								{ContainerPort: 5052},
								{ContainerPort: 5053},
								{ContainerPort: 8080},
								{ContainerPort: 5054},
								{ContainerPort: 9000},
								{ContainerPort: 13000},
								{ContainerPort: 12000},
							},

							Args: []string{
								"lighthouse",
								"bn",
								"--datadir=/Node-1/consensus/lighthouse-datadir/",
								"--http",
								"--execution-jwt=/Node-1/execution/secrets/jwtsecret",
								"--metrics",
								"--http-address=0.0.0.0",
								"--execution-endpoint=http://geth-execution-service:8551",
								"--metrics-address=0.0.0.0",
								"--enr-udp-port=9000",
								"--enr-tcp-port=9000",
								"--discovery-port=9000",
								"--purge-db",
								"--metrics-port=8080",
								"--testnet-dir=/shared-files/network/",
								"--enable-private-discovery",
								"--subscribe-all-subnets",
								"--debug-level=trace",
								"--http-allow-origin=*",
								"--eth1",
								"--import-all-attestations",
								"--validator-monitor-auto",
								"--disable-upnp",
								"--disable-backfill-rate-limiting",
								"--http-allow-sync-stalled",
								"--spec=minimal",
								"--gui",
								"--staking",
								"--genesis-backfill",
								"--target-peers=2",
								"--port=5054",

								// "--accept-terms-of-use",
								// "--jwt-secret=/Node-1/execution/jwtsecret",
								// "--suggested-fee-recipient=0x" + address,
								// "--force-clear-db",
								// "--minimum-peers-per-subnet=0",
								// "--enable-debug-rpc-endpoints",

								// "--p2p-static-id",
								// "--p2p-host-ip=192.168.253.108",
								// "--p2p-tcp-port=13000",
								// "--p2p-udp-port=12000",
								// "--bootstrap-node=enr:-Ku4QGQJf2bcDAwVGvbvtq3AB4KKwAvStTenY-i_QnW2ABNRRBncIU_5qR_e_um-9t3s9g-Y5ZfFATj1nhtzq6lvgc4Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpDEqCQNAAAAAv__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQNoed9JnQh7ltcAacHEGOjwocL1BhMQbYTgaPX0kFuXtIN1ZHCCE4g --clear-db",
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
									ClaimName: "mypvc", // Replace with your PVC name
								},
							},
						},
						{
							Name: "lighthousebeacon",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker1",
								},
							},
						},
					},
				},
			},
		},
	}

	// Create the StatefulSet
	createdStatefulSet, err := clientset.AppsV1().StatefulSets(namespace).Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("StatefulSet '%s' created in namespace '%s'\n", createdStatefulSet.Name, namespace)
	return nil
}

func BeaconNode2(config *rest.Config, namespace string, chainId int) error {

	filePath := "/mnt/shared-files/Node-2-address.txt"

	// Read the file
	mydata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Convert the file data to a string
	address := strings.TrimSpace(string(mydata))

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// boot nodes
	cmd := exec.Command("bash", "./scripts/boot_node.sh")
	enr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %v\n", err)
	}

	boot_nodes := strings.TrimSpace(string(enr))
	fmt.Println(boot_nodes)

	cmd_enr := exec.Command("bash", "./scripts/peerId_script.sh")
	peerId, err := cmd_enr.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %v\n", err)
	}

	peer_id := strings.TrimSpace(string(peerId))
	fmt.Println(peer_id)

	trusted_peers := "/ip4/192.168.253.108/tcp/9000/p2p/" + string(peer_id)
	fmt.Println(trusted_peers)

	enr_boot_nodes := fmt.Sprintf("--boot-nodes=%s", boot_nodes)
	libp2p_address := fmt.Sprintf("--libp2p-addresses=%s", trusted_peers)
	trustedPeerId := fmt.Sprintf("--trusted-peers=%s", peer_id)

	// fmt.Println("newString address \n", newString)

	// Define the StatefulSet
	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-statefulset2", // Update with your desired name
			Labels: map[string]string{
				"app": "lighthousebeacon2",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "lighthousebeacon-service2", // Update with your desired service name
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "lighthousebeacon2",
					"comp": "lighthousebeacon-svc2",
				},
			},
			Replicas: int32Ptr(1), // Set the number of replicas as needed
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "lighthousebeacon2",
						"comp": "lighthousebeacon-svc2",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "lighthousebeacon2",
							Image: "sigp/lighthouse:v4.5.0-amd64-dev",
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "lighthousebeacon",
									MountPath: "/Node-2/",
									SubPath:   "Node-2",
								},
							},
							Ports: []v1.ContainerPort{
								{ContainerPort: 5052},
								{ContainerPort: 5053},
								{ContainerPort: 8080},
								{ContainerPort: 5054},
								{ContainerPort: 9000},
								{ContainerPort: 13000},
								{ContainerPort: 12000},
							},

							Args: []string{
								"lighthouse",
								"bn",
								"--datadir=/Node-2/consensus/lighthouse-datadir/",
								"--http",
								"--execution-jwt=/Node-2/execution/secrets/jwtsecret",
								"--metrics",
								"--http-address=0.0.0.0",
								"--execution-endpoint=http://geth-execution-service2:8551",
								"--metrics-address=0.0.0.0",
								"--enr-udp-port=9000",
								"--enr-tcp-port=9000",
								"--discovery-port=9000",
								"--purge-db",
								"--metrics-port=8080",
								"--testnet-dir=/shared-files/network/",
								"--enable-private-discovery",
								"--subscribe-all-subnets",
								"--debug-level=trace",
								"--http-allow-origin=*",
								"--eth1",
								"--import-all-attestations",
								"--slots-per-restore-point=32",
								"--validator-monitor-auto",
								"--disable-upnp",
								"--disable-backfill-rate-limiting",
								"--http-allow-sync-stalled",
								"--spec=minimal",
								"--gui",
								"--port=5054",
								"--staking",
								"--genesis-backfill",
								"--suggested-fee-recipient=" + string(address),
								enr_boot_nodes,
								libp2p_address,
								trustedPeerId,
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
									ClaimName: "mypvc", // Replace with your PVC name
								},
							},
						},
						{
							Name: "lighthousebeacon",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker2",
								},
							},
						},
					},
				},
			},
		},
	}

	// Create the StatefulSet
	createdStatefulSet, err := clientset.AppsV1().StatefulSets(namespace).Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("StatefulSet '%s' created in namespace '%s'\n", createdStatefulSet.Name, namespace)
	return nil
}

func BeaconNode3(config *rest.Config, namespace string, chainId int) error {

	filePath := "/mnt/shared-files/Node-3-address.txt"

	// Read the file
	mydata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Convert the file data to a string
	address := strings.TrimSpace(string(mydata))

	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// boot nodes
	cmd := exec.Command("bash", "./scripts/boot_node.sh")
	enr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %v\n", err)
	}

	boot_nodes := strings.TrimSpace(string(enr))
	fmt.Println(boot_nodes)

	cmd_enr := exec.Command("bash", "./scripts/peerId_script.sh")
	peerId, err := cmd_enr.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running script: %v\n", err)
	}

	peer_id := strings.TrimSpace(string(peerId))
	fmt.Println(peer_id)

	trusted_peers := "/ip4/192.168.253.108/tcp/9000/p2p/" + string(peer_id)
	fmt.Println(trusted_peers)

	enr_boot_nodes := fmt.Sprintf("--boot-nodes=%s", boot_nodes)
	libp2p_address := fmt.Sprintf("--libp2p-addresses=%s", trusted_peers)
	trustedPeerId := fmt.Sprintf("--trusted-peers=%s", peer_id)

	// fmt.Println("newString address \n", newString)

	// Define the StatefulSet
	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-statefulset3", // Update with your desired name
			Labels: map[string]string{
				"app": "lighthousebeacon3",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "lighthousebeacon-service3", // Update with your desired service name
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "lighthousebeacon3",
					"comp": "lighthousebeacon-svc3",
				},
			},
			Replicas: int32Ptr(1), // Set the number of replicas as needed
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "lighthousebeacon3",
						"comp": "lighthousebeacon-svc3",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "lighthousebeacon3",
							Image: "sigp/lighthouse:v4.5.0-amd64-dev",
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "lighthousebeacon",
									MountPath: "/Node-3/",
									SubPath:   "Node-3",
								},
							},
							Ports: []v1.ContainerPort{
								{ContainerPort: 5052},
								{ContainerPort: 5053},
								{ContainerPort: 8080},
								{ContainerPort: 5054},
								{ContainerPort: 9000},
								{ContainerPort: 13000},
								{ContainerPort: 12000},
							},

							Args: []string{
								"lighthouse",
								"bn",
								"--datadir=/Node-3/consensus/lighthouse-datadir/",
								"--http",
								"--execution-jwt=/Node-3/execution/secrets/jwtsecret",
								"--metrics",
								"--http-address=0.0.0.0",
								"--execution-endpoint=http://geth-execution-service3:8551",
								"--metrics-address=0.0.0.0",
								"--enr-udp-port=9000",
								"--enr-tcp-port=9000",
								"--discovery-port=9000",
								"--purge-db",
								"--metrics-port=8080",
								"--testnet-dir=/shared-files/network/",
								"--enable-private-discovery",
								"--subscribe-all-subnets",
								"--debug-level=trace",
								"--http-allow-origin=*",
								"--eth1",
								"--import-all-attestations",
								"--slots-per-restore-point=32",
								"--validator-monitor-auto",
								"--disable-upnp",
								"--disable-backfill-rate-limiting",
								"--http-allow-sync-stalled",
								"--spec=minimal",
								"--gui",
								"--port=5054",
								"--staking",
								"--genesis-backfill",
								"--suggested-fee-recipient=" + string(address),
								enr_boot_nodes,
								libp2p_address,
								trustedPeerId,
							},
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
									ClaimName: "mypvc", // Replace with your PVC name
								},
							},
						},
						{
							Name: "lighthousebeacon",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker3",
								},
							},
						},
					},
				},
			},
		},
	}

	// Create the StatefulSet
	createdStatefulSet, err := clientset.AppsV1().StatefulSets(namespace).Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("StatefulSet '%s' created in namespace '%s'\n", createdStatefulSet.Name, namespace)
	return nil
}

func int32Ptr(i int32) *int32 { return &i }
