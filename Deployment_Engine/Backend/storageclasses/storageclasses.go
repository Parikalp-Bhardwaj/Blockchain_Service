package storageclasses

import (
	"context"
	"fmt"

	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func StorageClasses1(config *rest.Config) error {

	storageClassName := "local-storage-class-worker1"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	storageClass := &v1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: storageClassName,
		},
		Provisioner: "kubernetes.io/no-provisioner",
	}
	volumeBindingMode := v1.VolumeBindingMode("Immediate")
	storageClass.VolumeBindingMode = &volumeBindingMode

	// Create the StorageClass
	_, err = clientset.StorageV1().StorageClasses().Create(context.TODO(), storageClass, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating StorageClass: %v\n", err)
		return err
	}

	fmt.Printf("StorageClass '%s' created successfully.\n", storageClassName)
	return nil

}

// Define the StorageClass for server 2

func StorageClasses2(config *rest.Config) error {

	storageClassName := "local-storage-class-worker2"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	storageClass := &v1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: storageClassName,
		},
		Provisioner: "kubernetes.io/no-provisioner",
	}
	volumeBindingMode := v1.VolumeBindingMode("Immediate")
	storageClass.VolumeBindingMode = &volumeBindingMode

	_, err = clientset.StorageV1().StorageClasses().Create(context.TODO(), storageClass, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating StorageClass: %v\n", err)
		return err
	}

	fmt.Printf("StorageClass '%s' created successfully.\n", storageClassName)
	return nil

}

// Define the StorageClass for server 3

func StorageClasses3(config *rest.Config) error {

	storageClassName := "local-storage-class-worker3"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	storageClass := &v1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: storageClassName,
		},
		Provisioner: "kubernetes.io/no-provisioner",
	}
	volumeBindingMode := v1.VolumeBindingMode("Immediate")
	storageClass.VolumeBindingMode = &volumeBindingMode

	_, err = clientset.StorageV1().StorageClasses().Create(context.TODO(), storageClass, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating StorageClass: %v\n", err)
		return err
	}

	fmt.Printf("StorageClass '%s' created successfully.\n", storageClassName)
	return nil

}

func StorageClasses4(config *rest.Config) error {

	storageClassName := "local-storage-class-worker4"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	storageClass := &v1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: storageClassName,
		},
		Provisioner: "kubernetes.io/no-provisioner",
	}
	volumeBindingMode := v1.VolumeBindingMode("Immediate")
	storageClass.VolumeBindingMode = &volumeBindingMode

	_, err = clientset.StorageV1().StorageClasses().Create(context.TODO(), storageClass, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating StorageClass: %v\n", err)
		return err
	}

	fmt.Printf("StorageClass '%s' created successfully.\n", storageClassName)
	return nil

}
