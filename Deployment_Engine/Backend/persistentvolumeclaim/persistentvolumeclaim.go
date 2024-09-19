package persistentvolumeclaim

import (
	"context"
	"fmt"


	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func PersistentVolumeClaim1(config *rest.Config, namespace string) error {
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Define the PVC configuration
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: "data-worker1", // Name of the PVC
		},
		Spec: v1.PersistentVolumeClaimSpec{
			StorageClassName: ptr("local-storage-class-worker1"), // StorageClass name
			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: resource.MustParse("30Gi"), // Requested storage size
				},
			},
		},
	}

	// Create the PVC
	_, err = clientset.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating PVC: %v\n", err)
		return err
	}

	fmt.Println("Persistent Volume Claim created successfully.")
	return nil

}

// Define Persistent Volume Claim for server 2
func PersistentVolumeClaim2(config *rest.Config, namespace string) error {
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Define the PVC configuration
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: "data-worker2", // Name of the PVC
		},
		Spec: v1.PersistentVolumeClaimSpec{
			StorageClassName: ptr("local-storage-class-worker2"), // StorageClass name
			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: resource.MustParse("30Gi"), // Requested storage size
				},
			},
		},
	}

	// Create the PVC
	_, err = clientset.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating PVC: %v\n", err)
		return err
	}

	fmt.Println("Persistent Volume Claim created successfully.")
	return nil

}

// Define Persistent Volume Claim for server 3
func PersistentVolumeClaim3(config *rest.Config, namespace string) error {
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Define the PVC configuration
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: "data-worker3", // Name of the PVC
		},
		Spec: v1.PersistentVolumeClaimSpec{
			StorageClassName: ptr("local-storage-class-worker3"), // StorageClass name
			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: resource.MustParse("30Gi"), // Requested storage size
				},
			},
		},
	}

	// Create the PVC
	_, err = clientset.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating PVC: %v\n", err)
		return err
	}

	fmt.Println("Persistent Volume Claim created successfully.")
	return nil

}

// Define Persistent Volume Claim for server 4
func PersistentVolumeClaim4(config *rest.Config, namespace string) error {
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Define the PVC configuration
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: "data-worker4", // Name of the PVC
		},
		Spec: v1.PersistentVolumeClaimSpec{
			StorageClassName: ptr("local-storage-class-worker4"), // StorageClass name
			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: resource.MustParse("30Gi"), // Requested storage size
				},
			},
		},
	}

	// Create the PVC
	_, err = clientset.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating PVC: %v\n", err)
		return err
	}

	fmt.Println("Persistent Volume Claim created successfully.")
	return nil

}

func ptr(s string) *string {
	return &s
}
