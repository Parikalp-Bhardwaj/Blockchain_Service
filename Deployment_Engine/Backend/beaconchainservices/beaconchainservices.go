package beaconchainservices

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/rest"
)

func ServicesBeaconNode(config *rest.Config, namespace string) error {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	// Create the beacon-service
	err = createBeaconService(clientset, namespace)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}

	// Create the lighthousebeacon-service
	err = createlighthouseBeaconService(clientset, namespace)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	return nil
}

func createBeaconService(clientset *kubernetes.Clientset, namespace string) error {
	serviceClient := clientset.CoreV1().Services(namespace)

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "beacon-service",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": "lighthousebeacon",
			},
			// ************************** We Changed it ClusterIP: "None" to Type: v1.ServiceTypeNodePort,
			// and NodePort 30000
			// ClusterIP: "None",
			// Type: v1.ServiceTypeNodePort,
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "grpc",
					Port:       5053,
					TargetPort: intstr.FromInt(5053),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "metrics",
					Port:       8080,
					TargetPort: intstr.FromInt(8080),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "debug",
					Port:       5052,
					TargetPort: intstr.FromInt(5052),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof",
					Port:       9000,
					TargetPort: intstr.FromInt(9000),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof2",
					Port:       9000,
					TargetPort: intstr.FromInt(9000),
					Protocol:   v1.ProtocolUDP,
				},
				{
					Name:       "pprof3",
					Port:       13000,
					TargetPort: intstr.FromInt(13000),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof4",
					Port:       12000,
					TargetPort: intstr.FromInt(12000),
					Protocol:   v1.ProtocolUDP,
				},
				// {
				// 	Name:       "execution",
				// 	Port:       30303,
				// 	TargetPort: intstr.FromInt(30303),
				// 	Protocol:   v1.ProtocolTCP,
				// },
				// {
				// 	Name:       "execution2",
				// 	Port:       30303,
				// 	TargetPort: intstr.FromInt(30303),
				// 	Protocol:   v1.ProtocolUDP,
				// },
			},
			ExternalIPs: []string{"192.168.253.108"},
		},
	}

	createServices, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	fmt.Printf("Services ''%s' created in namespace '%s' \n", createServices.Name, namespace)
	return err
}

func createlighthouseBeaconService(clientset *kubernetes.Clientset, namespace string) error {
	serviceClient := clientset.CoreV1().Services(namespace)

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-service",
			Labels: map[string]string{
				"app": "crystal-mev",
			},
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None",
			Selector: map[string]string{
				"app": "lighthousebeacon",
			},
			Ports: []v1.ServicePort{
				{
					Port:       5054,
					TargetPort: intstr.FromInt(5054),
					Protocol:   v1.ProtocolTCP,
				},
			},
			// ExternalIPs: []string{"192.168.253.108"},
		},
	}

	createServices, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	fmt.Printf("Services ''%s' created in namespace '%s' \n", createServices.Name, namespace)
	return err
}

// Server 2

func ServicesBeaconNode2(config *rest.Config, namespace string) error {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	// Create the beacon-service
	err = createBeaconService2(clientset, namespace)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}

	// Create the lighthousebeacon-service
	err = createLighthouseBeaconService2(clientset, namespace)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	return nil
}

// Server 2

func createBeaconService2(clientset *kubernetes.Clientset, namespace string) error {
	serviceClient := clientset.CoreV1().Services(namespace)

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "beacon-service2",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": "lighthousebeacon2",
			},
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "grpc",
					Port:       5053,
					TargetPort: intstr.FromInt(5053),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "metrics",
					Port:       8080,
					TargetPort: intstr.FromInt(8080),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "debug",
					Port:       5052,
					TargetPort: intstr.FromInt(5052),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof",
					Port:       9000,
					TargetPort: intstr.FromInt(9000),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof2",
					Port:       9000,
					TargetPort: intstr.FromInt(9000),
					Protocol:   v1.ProtocolUDP,
				},
				{
					Name:       "pprof3",
					Port:       13000,
					TargetPort: intstr.FromInt(13000),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof4",
					Port:       12000,
					TargetPort: intstr.FromInt(12000),
					Protocol:   v1.ProtocolUDP,
				},
				
			},
			ExternalIPs: []string{"192.168.253.106"},
		},
	}

	createServices, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	fmt.Printf("Service '%s' created in namespace '%s' \n", createServices.Name, namespace)
	return err
}

func createLighthouseBeaconService2(clientset *kubernetes.Clientset, namespace string) error {
	serviceClient := clientset.CoreV1().Services(namespace)

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-service2",
			Labels: map[string]string{
				"app": "crystal-mev2",
			},
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None",
			Selector: map[string]string{
				"app":  "lighthousebeacon2",
				"comp": "lighthousebeacon-svc2",
			},
			Ports: []v1.ServicePort{
				{
					Port:       5054,
					TargetPort: intstr.FromInt(5054),
					Protocol:   v1.ProtocolTCP,
				},
				// Add other ports as needed
			},
		},
	}

	createServices, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	fmt.Printf("Service '%s' created in namespace '%s' \n", createServices.Name, namespace)
	return err
}

// Server 3

func ServicesBeaconNode3(config *rest.Config, namespace string) error {

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	// Create the beacon-service
	err = createBeaconService3(clientset, namespace)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}

	// Create the lighthousebeacon-service
	err = createlighthouseBeaconService3(clientset, namespace)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	return nil
}

func createBeaconService3(clientset *kubernetes.Clientset, namespace string) error {
	serviceClient := clientset.CoreV1().Services(namespace)

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "beacon-service3",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app":  "lighthousebeacon3",
				"comp": "lighthousebeacon-svc3",
			},
			// ************************** We Changed it ClusterIP: "None" to Type: v1.ServiceTypeNodePort,
			// and NodePort 30000
			// ClusterIP: "None",
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Name:       "grpc",
					Port:       5053,
					TargetPort: intstr.FromInt(5053),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "metrics",
					Port:       8080,
					TargetPort: intstr.FromInt(8080),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "debug",
					Port:       5052,
					TargetPort: intstr.FromInt(5052),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof",
					Port:       9000,
					TargetPort: intstr.FromInt(9000),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof2",
					Port:       9000,
					TargetPort: intstr.FromInt(9000),
					Protocol:   v1.ProtocolUDP,
				},
				{
					Name:       "pprof3",
					Port:       13000,
					TargetPort: intstr.FromInt(13000),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "pprof4",
					Port:       12000,
					TargetPort: intstr.FromInt(12000),
					Protocol:   v1.ProtocolUDP,
				},
				// {
				// 	Name:       "execution",
				// 	Port:       30303,
				// 	TargetPort: intstr.FromInt(30303),
				// 	Protocol:   v1.ProtocolTCP,
				// },
				// {
				// 	Name:       "execution2",
				// 	Port:       30303,
				// 	TargetPort: intstr.FromInt(30303),
				// 	Protocol:   v1.ProtocolUDP,
				// },
			},
			ExternalIPs: []string{"192.168.253.107"},
		},
	}

	createServices, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	fmt.Printf("Services ''%s' created in namespace '%s' \n", createServices.Name, namespace)
	return err
}

func createlighthouseBeaconService3(clientset *kubernetes.Clientset, namespace string) error {
	serviceClient := clientset.CoreV1().Services(namespace)

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "lighthousebeacon-service3",
			Labels: map[string]string{
				"app": "crystal-mev3",
			},
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None",
			Selector: map[string]string{
				"app":  "lighthousebeacon3",
				"comp": "lighthousebeacon-svc3",
			},
			Ports: []v1.ServicePort{
				{
					Port:       5054,
					TargetPort: intstr.FromInt(5054),
					Protocol:   v1.ProtocolTCP,
				},
			},
		},
	}

	createServices, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	fmt.Printf("Services ''%s' created in namespace '%s' \n", createServices.Name, namespace)
	return err
}
