package generatestate

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

func GenerateState(config *rest.Config, namespace string, newTimestamp int64, chainId int) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobName := "genesis-state"
	// commandArgument := fmt.Sprintf("--min-genesis-time=%d", newTimestamp)
	// eth1_id := fmt.Sprintf("--eth1-id=%d", chainId)

	// filePath := "/mnt/shared-files/network/deposit_contract_block_hash.txt"

	// content, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// }

	// blockHash := strings.TrimSpace(string(content))

	// eth1BLockHash := fmt.Sprintf("--eth1-block-hash=%s", string(blockHash))
	// fmt.Println("eth1BlockHash ", eth1BLockHash)

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
						"app": "genesis-state",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						// {
						// 	Name:  "genesis-state",
						// 	Image: "sigp/lcli:v4.5.0",
						// 	Args: []string{
						// 		"lcli",
						// 		"new-testnet",
						// 		"--spec=minimal",
						// 		"--deposit-contract-address=4242424242424242424242424242424242424242",
						// 		"--testnet-dir=/shared-files/network/",
						// 		"--min-genesis-active-validator-count=2",
						// 		commandArgument,
						// 		"--genesis-delay=0",
						// 		"--genesis-fork-version=0x42424242",
						// 		"--altair-fork-epoch=0",
						// 		"--bellatrix-fork-epoch=0",
						// 		"--capella-fork-epoch=1",
						// 		"--ttd=0",
						// 		eth1BLockHash,
						// 		eth1_id,
						// 		"--eth1-follow-distance=6",
						// 		"--seconds-per-slot=3",
						// 		"--seconds-per-eth1-block=3",
						// 		"--proposer-score-boost=40",
						// 		"--validator-count=2",
						// 		"--interop-genesis-state",
						// 		"--force",
						// 	},

						////////////////////////

						{
							Name:    "genesis-state",
							Image:   "parikalp456/lfi-validator:1.0",
							Command: []string{"/bin/bash", "-c"},
							Args: []string{
								`#!/bin/bash
									 /shared-files/createConfig.sh`,
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
