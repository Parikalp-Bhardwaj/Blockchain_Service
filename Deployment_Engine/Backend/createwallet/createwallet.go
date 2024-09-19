package createwallet

import (
	"context"
	"fmt"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	// "k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/rest"
)

func CreateWallet_Server1(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobName := "geth-job-wallet-1"

	jobsClient := clientset.BatchV1().Jobs(namespace)

	// Define your Job configuration
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Parallelism: int32Ptr(1),
			Completions: int32Ptr(1),
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "geth-job-wallet-1",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "geth-genesis",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`#!/bin/bash
								 /shared-files/create_wallet_Node_1.sh`,
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "geth",
									MountPath: "/Node-1/",
									SubPath:   "Node-1",
								},
								{
									Name:      "gethfile",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker1",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfile",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc",
								},
							},
						},
						{
							Name: "geth",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker1",
								},
							},
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	createdGethJob, err := jobsClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Job '%s' created in namespace '%s'\n", createdGethJob.Name, namespace)
	waitForJobCompletion(clientset, jobName, namespace)
	return nil
}

// For Server 2

func CreateWallet_Server2(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobName := "geth-job-wallet-2"

	jobsClient := clientset.BatchV1().Jobs(namespace)

	// Define your Job configuration
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Parallelism: int32Ptr(1),
			Completions: int32Ptr(1),
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "geth-job-wallet-2",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "geth-genesis",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`#!/bin/bash
								 /shared-files/create_wallet_Node_2.sh`,
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "geth",
									MountPath: "/Node-2/",
									SubPath:   "Node-2",
								},
								{
									Name:      "gethfile",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker2",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfile",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc",
								},
							},
						},
						{
							Name: "geth",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker2",
								},
							},
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	createdGethJob, err := jobsClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Job '%s' created in namespace '%s'\n", createdGethJob.Name, namespace)
	waitForJobCompletion(clientset, jobName, namespace)
	return nil
}

// For Server 3

func CreateWallet_Server3(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobName := "geth-job-wallet-3"

	jobsClient := clientset.BatchV1().Jobs(namespace)

	// Define your Job configuration
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Parallelism: int32Ptr(1),
			Completions: int32Ptr(1),
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "geth-job-wallet-3",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "geth-genesis",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`#!/bin/bash
								 /shared-files/create_wallet_Node_3.sh`,
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "geth",
									MountPath: "/Node-3/",
									SubPath:   "Node-3",
								},
								{
									Name:      "gethfile",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker3",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfile",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc",
								},
							},
						},
						{
							Name: "geth",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker3",
								},
							},
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	createdGethJob, err := jobsClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Job '%s' created in namespace '%s'\n", createdGethJob.Name, namespace)
	waitForJobCompletion(clientset, jobName, namespace)
	return nil
}

// For Server 4
func CreateWallet_Server4(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobName := "geth-job-wallet-4"

	jobsClient := clientset.BatchV1().Jobs(namespace)

	// Define your Job configuration
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Parallelism: int32Ptr(1),
			Completions: int32Ptr(1),
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "geth-job-wallet-4",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "geth-genesis",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`#!/bin/bash
								 /shared-files/create_wallet_Node_4.sh`,
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "geth",
									MountPath: "/Node-4/",
									SubPath:   "Node-4",
								},
								{
									Name:      "gethfile",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker4",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfile",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc",
								},
							},
						},
						{
							Name: "geth",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker4",
								},
							},
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	createdGethJob, err := jobsClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Job '%s' created in namespace '%s'\n", createdGethJob.Name, namespace)
	waitForJobCompletion(clientset, jobName, namespace)
	return nil
}

func int32Ptr(i int32) *int32 {
	return &i
}

func waitForJobCompletion(clientset *kubernetes.Clientset, jobName, namespace string) {
	// Wait for the Job to complete
	jobClient := clientset.BatchV1().Jobs(namespace)

	for {
		job, err := jobClient.Get(context.TODO(), jobName, metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}

		for _, condition := range job.Status.Conditions {
			if condition.Type == batchv1.JobComplete {
				fmt.Printf("Job '%s' in namespace '%s' completed successfully\n", jobName, namespace)
				return
			} else if condition.Type == batchv1.JobFailed {
				fmt.Printf("Job '%s' in namespace '%s' failed\n", jobName, namespace)
				return
			}
		}

		// Sleep for a while before checking again
		time.Sleep(5 * time.Second)
	}
}
