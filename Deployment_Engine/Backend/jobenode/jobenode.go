package jobenode

import (
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetEnode(config *rest.Config, namespace string) error {

	// // Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// Define the Job
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "data-retrieval-job",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "data-container",
							Image:   "parikalp456/lfi-validator:2.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								"/shared-files/nodekey_script.sh",
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfile",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "data",
									MountPath: "/Node-1/",
									SubPath:   "Node-1/",
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []v1.Volume{
						{
							Name: "data",
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
			// StartingDeadlineSeconds: int64Ptr(20),
			BackoffLimit:          int32Ptr(4),
			ActiveDeadlineSeconds: int64Ptr(30),
		},
	}

	// Create the Job
	_, err = clientset.BatchV1().Jobs(namespace).Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating Job: %v\n", err)
		return nil
	}

	fmt.Println("Job created successfully.")
	return nil
}

func int32Ptr(i int32) *int32 {
	return &i
}
func int64Ptr(i int64) *int64 {
	return &i
}
