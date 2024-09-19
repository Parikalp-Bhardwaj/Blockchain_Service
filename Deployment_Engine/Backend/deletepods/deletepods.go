package deletepods

import (
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	// "k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/rest"
)

func DeletePods1(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())

	}

	// Define the Job
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "delete-files-1",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "delete-pods-1",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`
								#!/bin/bash
								# Server 1
								rm -rf /Node-1/execution/geth/*
								rm -rf /Node-1/execution/geth
								rm -rf /Node-1/consensus/beacondata
								rm -rf /Node-1/consensus/validator
								rm -rf /Node-2/execution/keystore/*
								rm -rf /Node-1/execution/keystore
								rm -rf /Node-1/execution/geth.ipc
								rm -rf /Node-1/consensus/lighthouse-datadir/


								rm -rf /shared-files/Node-1-address.txt
								rm -rf /shared-files/Node-2-address.txt
								rm -rf /shared-files/Node-3-address.txt
								rm -rf /shared-files/Node-4-address.txt
								rm -rf /shared-files/nodekey
								rm -rf /shared-files/network/genesis.ssz
								rm -rf /shared-files/network/tranches
								

								`,
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfile",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "data-1",
									MountPath: "/Node-1/",
									SubPath:   "Node-1/",
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []v1.Volume{
						{
							Name: "data-1",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker1",
								},
							},
						},
						{
							Name: "gethfile",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc",
								},
							},
						},
					},
				},
			},
			BackoffLimit: int32Ptr(4),
		},
	}

	// Create the Job
	_, err = clientset.BatchV1().Jobs(namespace).Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating Job: %v\n", err)
		return err
	}

	fmt.Println("Job created successfully.")

	return nil

}

// Server 2
func DeletePods2(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())

	}

	// Define the Job
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "delete-files-2",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "delete-pods-2",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`
								#!/bin/bash
								# Server 2
								rm -rf /Node-2/execution/geth/*
								rm -rf /Node-2/execution/geth
								rm -rf /Node-2/consensus/beacondata
								rm -rf /Node-2/consensus/validatordata
								rm -rf /Node-2/execution/keystore/*
								rm -rf /Node-2/execution/keystore
								rm -rf /Node-2/execution/geth.ipc
								rm -rf /Node-2/consensus/lighthouse-datadir/

								`,
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "datanode",
									MountPath: "/Node-2/",
									SubPath:   "Node-2/",
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []v1.Volume{
						{
							Name: "datanode",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker2",
								},
							},
						},
					},
				},
			},
			BackoffLimit: int32Ptr(4),
		},
	}

	// Create the Job
	_, err = clientset.BatchV1().Jobs(namespace).Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating Job: %v\n", err)
		return err
	}

	fmt.Println("Job created successfully.")

	return nil

}

// Server 3
func DeletePods3(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())

	}

	// Define the Job
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "delete-files-3",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "delete-pods-3",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`
								#!/bin/bash

								# Server 3
								rm -rf /Node-3/execution/geth/*
								rm -rf /Node-3/execution/geth
								rm -rf /Node-3/consensus/beacondata
								rm -rf /Node-3/consensus/validatordata
								rm -rf /Node-3/execution/keystore/*
								rm -rf /Node-3/execution/keystore
								rm -rf /Node-3/execution/geth.ipc
								rm -rf /Node-3/consensus/lighthouse-datadir/
								`,
							},
							VolumeMounts: []v1.VolumeMount{

								{
									Name:      "data-3",
									MountPath: "/Node-3/",
									SubPath:   "Node-3/",
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []v1.Volume{

						{
							Name: "data-3",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker3",
								},
							},
						},
					},
				},
			},
			BackoffLimit: int32Ptr(4),
		},
	}

	// Create the Job
	_, err = clientset.BatchV1().Jobs(namespace).Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating Job: %v\n", err)
		return err
	}

	fmt.Println("Job created successfully.")

	return nil

}

// Server 4
func DeletePods4(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())

	}

	// Define the Job
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "delete-file-4",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "delete-pods-4",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`
								#!/bin/bash


								# Server 4
								rm -rf /Node-4/execution/geth
								rm -rf /Node-4/consensus/beacondata
								rm -rf /Node-4/consensus/validatordata
								rm -rf /Node-4/execution/keystore
								rm -rf /Node-4/execution/geth.ipc


								`,
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "data-4",
									MountPath: "/Node-4/",
									SubPath:   "Node-4/",
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []v1.Volume{
						{
							Name: "data-4",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker4",
								},
							},
						},
					},
				},
			},
			BackoffLimit: int32Ptr(4),
		},
	}

	// Create the Job
	_, err = clientset.BatchV1().Jobs(namespace).Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating Job: %v\n", err)
		return err
	}

	fmt.Println("Job created successfully.")

	return nil

}

func int32Ptr(i int32) *int32 {
	return &i
}
