package beaconnodeservices

import (
	"context"
	// "flag"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	// "k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/rest"
)

func ServicesGethNode(config *rest.Config, namespace string) error {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// namespace := "default" // Replace with the desired namespace

	gethNPTCPService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-np-tcp",
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "geth-p2p-tcp-np",
					Protocol:   v1.ProtocolTCP,
					Port:       30303,
					TargetPort: intstr.FromInt(30303),
					NodePort:   30303,
				},
				{
					Name:       "geth-p2p-tcp-np2",
					Protocol:   v1.ProtocolUDP,
					Port:       30303,
					TargetPort: intstr.FromInt(30303),
					NodePort:   30303,
				},
			},

			ExternalIPs: []string{"192.168.253.108"},

			Selector: map[string]string{
				"app":  "crystal-mev",
				"comp": "geth-node-pod",
			},
		},
	}

	createdGethNPTCPService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethNPTCPService, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethNPTCPService.Name, namespace)

	// Define and create the geth-np-udp Service
	gethNPUdpService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-np-udp",
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "geth-p2p-udp-np",
					Protocol:   v1.ProtocolUDP,
					Port:       30304,
					TargetPort: intstr.FromInt(30304),
					NodePort:   30304,
				},
				{
					Name:       "geth-p2p-udp-np2",
					Protocol:   v1.ProtocolTCP,
					Port:       30304,
					TargetPort: intstr.FromInt(30304),
					NodePort:   30304,
				},
			},
			Selector: map[string]string{
				"app":  "crystal-mev",
				"comp": "geth-node-pod",
			},
		},
	}

	createdGethNPUdpService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethNPUdpService, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethNPUdpService.Name, namespace)

	gethExecutionService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-execution-service",
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None", // Headless service
			Selector: map[string]string{
				"app":  "crystal-mev",
				"comp": "geth-node-pod",
			},
			Ports: []v1.ServicePort{
				{
					Port:       8551,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8551),
				},
			},
		},
	}

	createdGethExecutionService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethExecutionService, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethExecutionService.Name, namespace)

	// Define and create the geth-rpc-service Service
	gethRPCService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-rpc-service",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "crystal-mev",
				"comp": "geth-node-pod",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "rpc-endpoint",
					Port:       8545,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8545),
				},
			},
			ExternalIPs: []string{"192.168.253.108"},
		},
	}

	createdGethRPCService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethRPCService, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethRPCService.Name, namespace)

	// Define and create the geth-ws-service Service
	gethWSService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-ws-service",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "crystal-mev",
				"comp": "geth-node-pod",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "ws-endpoint",
					Port:       8546,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8546),
				},
			},
			ExternalIPs: []string{"192.168.253.108"},
		},
	}

	createdGethWSService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethWSService, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethWSService.Name, namespace)
	return nil
}

// Server 2
func ServicesGethNode2(config *rest.Config, namespace string) error {
	// // Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	fmt.Println("going to create 2")

	// Define and create the geth-np-tcp Service
	gethNPTCPService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-np-tcp2",
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "geth-p2p-tcp-np2",
					Protocol:   v1.ProtocolTCP,
					Port:       30303,
					TargetPort: intstr.FromInt(30303),
					// NodePort:   30303,
				},
				{
					Name:       "geth-p2p-tcp1-np2",
					Protocol:   v1.ProtocolUDP,
					Port:       30303,
					TargetPort: intstr.FromInt(30303),
					// NodePort:   30303,
				},
			},
			ExternalIPs: []string{"192.168.253.106"},
			Selector: map[string]string{
				"app":  "crystal-mev2",
				"comp": "geth-node-pod2",
			},
		},
	}

	createdGethNPTCPService, _ := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethNPTCPService, metav1.CreateOptions{})

	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethNPTCPService.Name, namespace)

	// Define and create the geth-np-udp Service
	gethNPUdpService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-np-udp2",
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "geth-p2p-udp-np2",
					Protocol:   v1.ProtocolUDP,
					Port:       30304,
					TargetPort: intstr.FromInt(30304),
				},
				{
					Name:       "geth-p2p-tcp2-np2",
					Protocol:   v1.ProtocolTCP,
					Port:       30304,
					TargetPort: intstr.FromInt(30304),
				},
			},
			// ExternalIPs: []string{"192.168.253.106"},
			Selector: map[string]string{
				"app":  "crystal-mev2",
				"comp": "geth-node-pod2",
			},
		},
	}

	createdGethNPUdpService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethNPUdpService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethNPUdpService.Name, namespace)

	// Define and create the geth-execution-service Service
	gethExecutionService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-execution-service2",
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None", // Headless service
			Selector: map[string]string{
				"app":  "crystal-mev2",
				"comp": "geth-node-pod2",
			},
			Ports: []v1.ServicePort{
				{
					Port:       8551,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8551),
				},
			},
		},
	}

	createdGethExecutionService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethExecutionService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethExecutionService.Name, namespace)

	// Define and create the geth-rpc-service Service
	gethRPCService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-rpc-service2",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "crystal-mev2",
				"comp": "geth-node-pod2",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "rpc-endpoint2",
					Port:       8545,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8545),
					// NodePort:   30318,
				},
			},
			ExternalIPs: []string{"192.168.253.106"},
		},
	}

	createdGethRPCService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethRPCService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethRPCService.Name, namespace)

	// Define and create the geth-ws-service Service
	gethWSService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-ws-service2",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "crystal-mev2",
				"comp": "geth-node-pod2",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "ws-endpoint2",
					Port:       8546,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8546),
					// NodePort:   30328,
				},
			},
			ExternalIPs: []string{"192.168.253.106"},
		},
	}

	createdGethWSService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethWSService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethWSService.Name, namespace)
	return nil
}

// Server 3

func ServicesGethNode3(config *rest.Config, namespace string) error {
	// // Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Define and create the geth-np-tcp Service
	gethNPTCPService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-np-tcp3",
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "geth-p2p-tcp-np3",
					Protocol:   v1.ProtocolTCP,
					Port:       30303,
					TargetPort: intstr.FromInt(30303),
					// NodePort:   30303,
				},
			},
			ExternalIPs: []string{"192.168.253.107"},
			Selector: map[string]string{
				"app":  "crystal-mev3",
				"comp": "geth-node-pod3",
			},
		},
	}

	createdGethNPTCPService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethNPTCPService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethNPTCPService.Name, namespace)

	// Define and create the geth-np-udp Service
	gethNPUdpService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-np-udp3",
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "geth-p2p-udp-np3",
					Protocol:   v1.ProtocolUDP,
					Port:       30304,
					TargetPort: intstr.FromInt(30304),
					// NodePort:   30363,
				},
			},
			Selector: map[string]string{
				"app":  "crystal-mev3",
				"comp": "geth-node-pod3",
			},
		},
	}

	createdGethNPUdpService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethNPUdpService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethNPUdpService.Name, namespace)

	// Define and create the geth-execution-service Service
	gethExecutionService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-execution-service3",
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None", // Headless service
			Selector: map[string]string{
				"app":  "crystal-mev3",
				"comp": "geth-node-pod3",
			},
			Ports: []v1.ServicePort{
				{
					Port:       8551,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8551),
				},
			},
		},
	}

	createdGethExecutionService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethExecutionService, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethExecutionService.Name, namespace)

	// Define and create the geth-rpc-service Service
	gethRPCService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-rpc-service3",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "crystal-mev3",
				"comp": "geth-node-pod3",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "rpc-endpoint3",
					Port:       8545,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8545),
					// NodePort:   30319,
				},
			},
			ExternalIPs: []string{"192.168.253.107"},
		},
	}

	createdGethRPCService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethRPCService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethRPCService.Name, namespace)

	// Define and create the geth-ws-service Service
	gethWSService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-ws-service3",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "crystal-mev3",
				"comp": "geth-node-pod3",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "ws-endpoint3",
					Port:       8546,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8546),
					// NodePort:   30329,
				},
			},
			ExternalIPs: []string{"192.168.253.107"},
		},
	}

	createdGethWSService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethWSService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethWSService.Name, namespace)
	return nil
}

// Server 4

func ServicesGethNode4(config *rest.Config, namespace string) error {
	// // Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Define and create the geth-np-tcp Service
	gethNPTCPService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-np-tcp4",
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "geth-p2p-tcp-np4",
					Protocol:   v1.ProtocolTCP,
					Port:       30303,
					TargetPort: intstr.FromInt(30303),
					// NodePort:   30303,
				},
			},
			ExternalIPs: []string{"192.168.253.111"},
			Selector: map[string]string{
				"app":  "crystal-mev4",
				"comp": "geth-node-pod4",
			},
		},
	}

	createdGethNPTCPService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethNPTCPService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethNPTCPService.Name, namespace)

	// Define and create the geth-np-udp Service
	gethNPUdpService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-np-udp4",
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "geth-p2p-udp-np4",
					Protocol:   v1.ProtocolUDP,
					Port:       30304,
					TargetPort: intstr.FromInt(30304),
					// NodePort:   30366,
				},
			},
			Selector: map[string]string{
				"app":  "crystal-mev4",
				"comp": "geth-node-pod4",
			},
		},
	}

	createdGethNPUdpService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethNPUdpService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethNPUdpService.Name, namespace)

	// Define and create the geth-execution-service Service
	gethExecutionService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-execution-service4",
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None", // Headless service
			Selector: map[string]string{
				"app":  "crystal-mev4",
				"comp": "geth-node-pod4",
			},
			Ports: []v1.ServicePort{
				{
					Port:       8551,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8551),
				},
			},
		},
	}

	createdGethExecutionService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethExecutionService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethExecutionService.Name, namespace)

	// Define and create the geth-rpc-service Service
	gethRPCService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-rpc-service4",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "crystal-mev4",
				"comp": "geth-node-pod4",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "rpc-endpoint4",
					Port:       8545,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8545),
					// NodePort:   30355,
				},
			},
			ExternalIPs: []string{"192.168.253.111"},
		},
	}

	createdGethRPCService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethRPCService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethRPCService.Name, namespace)

	// Define and create the geth-ws-service Service
	gethWSService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "geth-ws-service4",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "crystal-mev4",
				"comp": "geth-node-pod4",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "ws-endpoint4",
					Port:       8546,
					Protocol:   v1.ProtocolTCP,
					TargetPort: intstr.FromInt(8546),
					// NodePort:   30330,
				},
			},
			ExternalIPs: []string{"192.168.253.111"},
		},
	}

	createdGethWSService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethWSService, metav1.CreateOptions{})
	if err != nil {
		log.Printf("Error creating service: %v", err)
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethWSService.Name, namespace)
	return nil
}
