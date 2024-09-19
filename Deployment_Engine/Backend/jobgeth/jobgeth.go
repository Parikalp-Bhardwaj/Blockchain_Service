package jobgeth

import (
	"context"
	"fmt"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	// "k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/rest"
)

func JobGeth(config *rest.Config, namespace string) error {

	jobName := "geth-job"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Define the Job
	jobGenesis := clientset.BatchV1().Jobs(namespace)

	// Define and create the Geth Job
	gethJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Parallelism: int32Ptr(1), // Number of parallel pods to run (1 in this case)
			Completions: int32Ptr(1), // Number of successful completions required (1 in this case)
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "geth",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "geth-genesis",
							Image: "ethereum/client-go:latest",

							Command: []string{"geth"},
							Args: []string{
								"--datadir=/Node-1/execution/",
								"init",
								"/shared-files/network/genesis.json",
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "prysmbeacon",
									MountPath: "/Node-1//",
									SubPath:   "Node-1",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker1",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfiles",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc", // Replace with your PVC name
								},
							},
						},
						{
							Name: "prysmbeacon",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker1",
								},
							},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
		},
	}

	createdGethJob, err := jobGenesis.Create(context.TODO(), gethJob, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Printf("Job '%s' created in namespace '%s'\n", createdGethJob.Name, namespace)
	waitForJobCompletion(clientset, jobName, namespace)
	return nil
}

// Server 2

func JobGeth2(config *rest.Config, namespace string) error {

	jobName := "geth-job2"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Define the Job
	jobGenesis := clientset.BatchV1().Jobs(namespace)

	// Define and create the Geth Job
	gethJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Parallelism: int32Ptr(1), // Number of parallel pods to run (1 in this case)
			Completions: int32Ptr(1), // Number of successful completions required (1 in this case)
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "geth2",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "geth-genesis2",
							Image: "ethereum/client-go:latest",
							// Resources: v1.ResourceRequirements{
							// 	Requests: v1.ResourceList{
							// 		v1.ResourceMemory: resourceQuantity("300Mi"),
							// 		v1.ResourceCPU:    resourceQuantity("300m"),
							// 	},
							// 	Limits: v1.ResourceList{
							// 		v1.ResourceMemory: resourceQuantity("500Mi"),
							// 		v1.ResourceCPU:    resourceQuantity("350m"),
							// 	},
							// },
							Command: []string{"geth"},
							Args: []string{
								"--datadir=/Node-2/execution/",
								"init",
								"/shared-files/network/genesis.json",
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "prysmbeacon",
									MountPath: "/Node-2/",
									SubPath:   "Node-2",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker2",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfiles",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc", // Replace with your PVC name
								},
							},
						},
						{
							Name: "prysmbeacon",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker2",
								},
							},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
		},
	}

	createdGethJob, err := jobGenesis.Create(context.TODO(), gethJob, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Printf("Job '%s' created in namespace '%s'\n", createdGethJob.Name, namespace)
	waitForJobCompletion(clientset, jobName, namespace)
	return nil
}

// Server 3
func JobGeth3(config *rest.Config, namespace string) error {

	jobName := "geth-job3"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Define the Job
	jobGenesis := clientset.BatchV1().Jobs(namespace)

	// Define and create the Geth Job
	gethJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Parallelism: int32Ptr(1), // Number of parallel pods to run (1 in this case)
			Completions: int32Ptr(1), // Number of successful completions required (1 in this case)
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "geth3",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "geth-genesis3",
							Image: "ethereum/client-go:latest",
							// Resources: v1.ResourceRequirements{
							// 	Requests: v1.ResourceList{
							// 		v1.ResourceMemory: resourceQuantity("300Mi"),
							// 		v1.ResourceCPU:    resourceQuantity("300m"),
							// 	},
							// 	Limits: v1.ResourceList{
							// 		v1.ResourceMemory: resourceQuantity("500Mi"),
							// 		v1.ResourceCPU:    resourceQuantity("350m"),
							// 	},
							// },
							Command: []string{"geth"},
							Args: []string{
								"--datadir=/Node-3/execution/",
								"init",
								"/shared-files/network/genesis.json",
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "prysmbeacon",
									MountPath: "/Node-3/",
									SubPath:   "Node-3",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker3",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfiles",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc", // Replace with your PVC name
								},
							},
						},
						{
							Name: "prysmbeacon",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker3",
								},
							},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
		},
	}

	createdGethJob, err := jobGenesis.Create(context.TODO(), gethJob, metav1.CreateOptions{})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Printf("Job '%s' created in namespace '%s'\n", createdGethJob.Name, namespace)
	waitForJobCompletion(clientset, jobName, namespace)
	return nil
}

// Server 4
func JobGeth4(config *rest.Config, namespace string) error {

	jobName := "geth-job4"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// Define the Job
	jobGenesis := clientset.BatchV1().Jobs(namespace)

	// Define and create the Geth Job
	gethJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Parallelism: int32Ptr(1), // Number of parallel pods to run (1 in this case)
			Completions: int32Ptr(1), // Number of successful completions required (1 in this case)
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "geth4",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "geth-genesis4",
							Image: "ethereum/client-go:latest",
							// Resources: v1.ResourceRequirements{
							// 	Requests: v1.ResourceList{
							// 		v1.ResourceMemory: resourceQuantity("300Mi"),
							// 		v1.ResourceCPU:    resourceQuantity("300m"),
							// 	},
							// 	Limits: v1.ResourceList{
							// 		v1.ResourceMemory: resourceQuantity("500Mi"),
							// 		v1.ResourceCPU:    resourceQuantity("350m"),
							// 	},
							// },
							Command: []string{"geth"},
							Args: []string{
								"--datadir=/Node-4/execution",
								"init",
								"/shared-files/genesis.json",
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "gethfiles",
									MountPath: "/shared-files",
									SubPath:   "shared-files",
								},
								{
									Name:      "prysmbeacon",
									MountPath: "/Node-4/",
									SubPath:   "Node-4",
								},
							},
						},
					},
					NodeSelector: map[string]string{
						"node-type": "worker4",
					},
					Volumes: []v1.Volume{
						{
							Name: "gethfiles",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "mypvc", // Replace with your PVC name
								},
							},
						},
						{
							Name: "prysmbeacon",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: "data-worker4", // Replace with your PVC name
								},
							},
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
		},
	}

	createdGethJob, err := jobGenesis.Create(context.TODO(), gethJob, metav1.CreateOptions{})

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

func resourceQuantity(s string) resource.Quantity {
	q, _ := resource.ParseQuantity(s)
	return q
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
