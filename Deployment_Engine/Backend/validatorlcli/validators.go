package validatorlcli

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

func ValidatorK8s1(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset

	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// jobName := "validator-lcli"

	// jobsClient := clientset.BatchV1().Jobs(namespace)

	// // Define your Job configuration
	// job := &batchv1.Job{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: jobName,
	// 	},
	// 	Spec: batchv1.JobSpec{
	// 		Parallelism: int32Ptr(1),
	// 		Completions: int32Ptr(1),
	// 		Template: v1.PodTemplateSpec{
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Labels: map[string]string{
	// 					"app": "validator-lcli",
	// 				},
	// 			},
	// 			Spec: v1.PodSpec{
	// 				Containers: []v1.Container{
	// 					{
	// 						Name:  "validator-lcli",
	// 						Image: "sigp/lcli:v4.5.0",
	// 						Args: []string{
	// 							"lcli",
	// 							"insecure-validators",
	// 							"--count=1",
	// 							"--base-dir=/Node-1/consensus/lighthouse-datadir/",
	// 							"--node-count=1",
	// 						},
	// 						VolumeMounts: []v1.VolumeMount{
	// 							{
	// 								Name:      "geth",
	// 								MountPath: "/Node-1/",
	// 								SubPath:   "Node-1",
	// 							},
	// 							{
	// 								Name:      "gethfile",
	// 								MountPath: "/shared-files",
	// 								SubPath:   "shared-files",
	// 							},
	// 						},
	// 					},
	// 				},
	// 				NodeSelector: map[string]string{
	// 					"node-type": "worker1",
	// 				},
	// 				Volumes: []v1.Volume{
	// 					{
	// 						Name: "gethfile",
	// 						VolumeSource: v1.VolumeSource{
	// 							PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
	// 								ClaimName: "mypvc",
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Name: "geth",
	// 						VolumeSource: v1.VolumeSource{
	// 							PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
	// 								ClaimName: "data-worker1",
	// 							},
	// 						},
	// 					},
	// 				},
	// 				RestartPolicy: "Never",
	// 			},
	// 		},
	// 	},
	// }

	// createdGethJob, err := jobsClient.Create(context.TODO(), job, metav1.CreateOptions{})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }
	// fmt.Printf("Job '%s' created in namespace '%s'\n", createdGethJob.Name, namespace)
	// waitForJobCompletion(clientset, jobName, namespace)
	// return nil

	//***********************************************************************//

	// Create a Kubernetes clientset .....................

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobName := "validator-lcli-1"

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
						"app": "validator-lcli-1",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "validator-lcli-1",
							Image: "sigp/lighthouse:v4.5.0-amd64-dev",
							Args: []string{
								"lighthouse",
								"account",
								"validator",
								"import",
								"--directory=/Node-1/lighthouse_validator/validator_keys/",
								"--datadir=/Node-1/consensus/lighthouse-datadir/",
								"--testnet-dir=/shared-files/network/",
								"--spec=minimal",
								"--password-file=/shared-files/password.txt",
								"--reuse-password",
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

// Server 2
func ValidatorK8s2(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobName := "validator-lcli-2"

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
						"app": "validator-lcli-2",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "validator-lcli-2",
							Image: "sigp/lighthouse:v4.5.0-amd64-dev",
							Args: []string{
								"lighthouse",
								"account",
								"validator",
								"import",
								"--directory=/Node-2/lighthouse_validator/validator_keys/",
								"--datadir=/Node-2/consensus/lighthouse-datadir/",
								"--testnet-dir=/shared-files/network/",
								"--spec=minimal",
								"--password-file=/shared-files/password.txt",
								"--reuse-password",
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

// Server 3
func ValidatorK8s3(config *rest.Config, namespace string) error {
	// Create a Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobName := "validator-lcli-3"

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
						"app": "validator-lcli-3",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "validator-lcli-3",
							Image: "sigp/lighthouse:v4.5.0-amd64-dev",
							Args: []string{
								"lighthouse",
								"account",
								"validator",
								"import",
								"--directory=/Node-3/lighthouse_validator/validator_keys/",
								"--datadir=/Node-3/consensus/lighthouse-datadir/",
								"--testnet-dir=/shared-files/network/",
								"--spec=minimal",
								"--password-file=/shared-files/password.txt",
								"--reuse-password",
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
