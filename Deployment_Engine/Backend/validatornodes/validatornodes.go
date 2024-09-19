package validatornodes

import (
	"context"
	"io/ioutil"
	"log"
	"strings"

	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	// "k8s.io/apimachinery/pkg/api/resource"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/rest"
)

func ValidatorNodes1(config *rest.Config, namespace string) error {

	filePath := "/mnt/shared-files/Node-1-address.txt"

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
	// Define the StatefulSet
	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-statefulset-validator1", // Update with your desired name
			Labels: map[string]string{
				"app": "lighthousebeacon",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "lighthousebeacon-service-validator1", // Update with your desired service name
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "lighthousebeacon-validator1",
					"comp": "lighthousebeacon-validator1",
				},
			},
			Replicas: int32Ptr(1), // Set the number of replicas as needed
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "lighthousebeacon-validator1",
						"comp": "lighthousebeacon-validator1",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "lighthousebeacon-validator1",
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
								{ContainerPort: 5062},
								{ContainerPort: 5064},
							},

							Args: []string{
								// "lighthouse",
								// "vc",
								// "--datadir=/Node-1/consensus/lighthouse-datadir/node_1/",
								// "--init-slashing-protection",
								// "--testnet-dir=/shared-files/network/",
								// "--beacon-nodes=http://192.168.253.108:5052",
								// "--suggested-fee-recipient=" + string(address),
								// "--spec=minimal",
								// "--http",
								// "--http-port=5062",
								// "--metrics",
								// "--metrics-address=0.0.0.0",
								// "--metrics-port=5064",
								// "--debug-level=info",

								"lighthouse",
								"vc",
								"--datadir=/Node-1/consensus/lighthouse-datadir/",
								"--testnet-dir=/shared-files/network/",
								"--init-slashing-protection",
								"--beacon-nodes=http://192.168.253.108:5052",
								"--spec=minimal",
								"--suggested-fee-recipient=" + string(address),
								"--http",
								"--http-port=5062",
								"--metrics",
								"--metrics-address=0.0.0.0",
								"--metrics-port=5064",
								"--debug-level=info",
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

// Server 2

func ValidatorNodes2(config *rest.Config, namespace string) error {

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
	// Define the StatefulSet
	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-statefulset-validator2", // Update with your desired name
			Labels: map[string]string{
				"app": "lighthousebeacon2",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "lighthousebeacon-service-validator2", // Update with your desired service name
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "lighthousebeacon-validator2",
					"comp": "lighthousebeacon-validator2",
				},
			},
			Replicas: int32Ptr(1), // Set the number of replicas as needed
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "lighthousebeacon-validator2",
						"comp": "lighthousebeacon-validator2",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "lighthousebeacon-validator2",
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
								{ContainerPort: 5062},
								{ContainerPort: 5064},
							},

							Args: []string{
								"lighthouse",
								"vc",
								"--datadir=/Node-2/consensus/lighthouse-datadir/",
								"--testnet-dir=/shared-files/network/",
								"--init-slashing-protection",
								"--beacon-nodes=http://192.168.253.106:5052",
								"--spec=minimal",
								"--suggested-fee-recipient=" + string(address),
								"--http",
								"--http-port=5062",
								"--metrics",
								"--metrics-address=0.0.0.0",
								"--metrics-port=5064",
								"--debug-level=info",

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

// Server 3

func ValidatorNodes3(config *rest.Config, namespace string) error {

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
	// Define the StatefulSet
	statefulSet := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-statefulset-validator3", // Update with your desired name
			Labels: map[string]string{
				"app": "lighthousebeacon3",
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: "lighthousebeacon-service-validator3", // Update with your desired service name
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "lighthousebeacon-validator3",
					"comp": "lighthousebeacon-validator3",
				},
			},
			Replicas: int32Ptr(1), // Set the number of replicas as needed
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "lighthousebeacon-validator3",
						"comp": "lighthousebeacon-validator3",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "lighthousebeacon-validator3",
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
								{ContainerPort: 5062},
								{ContainerPort: 5064},
							},

							Args: []string{
								"lighthouse",
								"vc",
								"--datadir=/Node-3/consensus/lighthouse-datadir/",
								"--testnet-dir=/shared-files/network/",
								"--init-slashing-protection",
								"--beacon-nodes=http://192.168.253.107:5052",
								"--spec=minimal",
								"--suggested-fee-recipient=" + string(address),
								"--http",
								"--http-port=5062",
								"--metrics",
								"--metrics-address=0.0.0.0",
								"--metrics-port=5064",
								"--debug-level=info",

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
