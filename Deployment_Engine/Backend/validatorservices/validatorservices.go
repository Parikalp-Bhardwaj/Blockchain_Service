package validatorservices

import (
	"context"
	// "flag"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	// "k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/rest"
)

func ValidatorServices1(config *rest.Config, namespace string) error {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Define and create the geth-ws-service Service
	gethWSService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "validator1-service",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "lighthousebeacon-validator1",
				"comp": "lighthousebeacon-validator1",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:     "vc-port1",
					Port:     5062,
					Protocol: v1.ProtocolTCP,
				},
				{
					Name:     "vc-metrics1",
					Port:     5064,
					Protocol: v1.ProtocolTCP,
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
func ValidatorServices2(config *rest.Config, namespace string) error {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Define and create the geth-ws-service Service
	gethWSService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "validator2-service",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "lighthousebeacon-validator2",
				"comp": "lighthousebeacon-validator2",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:     "vc-port2",
					Port:     5062,
					Protocol: v1.ProtocolTCP,
				},
				{
					Name:     "vc-metrics2",
					Port:     5064,
					Protocol: v1.ProtocolTCP,
				},
			},
			ExternalIPs: []string{"192.168.253.106"},
		},
	}

	createdGethWSService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethWSService, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethWSService.Name, namespace)
	return nil
}

// Server 3
func ValidatorServices3(config *rest.Config, namespace string) error {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Define and create the geth-ws-service Service
	gethWSService := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "validator3-service",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "lighthousebeacon-validator3",
				"comp": "lighthousebeacon-validator3",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:     "vc-port3",
					Port:     5062,
					Protocol: v1.ProtocolTCP,
				},
				{
					Name:     "vc-metrics3",
					Port:     5064,
					Protocol: v1.ProtocolTCP,
				},
			},
			ExternalIPs: []string{"192.168.253.107"},
		},
	}

	createdGethWSService, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), gethWSService, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Service '%s' created in namespace '%s'\n", createdGethWSService.Name, namespace)
	return nil
}
