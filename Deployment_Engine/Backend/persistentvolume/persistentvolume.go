package persistentvolume

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func PersistentVolume1(config *rest.Config) error {
	// Create a Kubernetes clientset.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Define the PersistentVolume specification.
	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: "local-pv-worker1",
		},
		Spec: v1.PersistentVolumeSpec{
			Capacity: v1.ResourceList{
				v1.ResourceStorage: resource.MustParse("30Gi"),
			},
			AccessModes:                   []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
			StorageClassName:              "local-storage-class-worker1",
			// VolumeMode:                    "Filesystem",
			NodeAffinity: &v1.VolumeNodeAffinity{
				Required: &v1.NodeSelector{
					NodeSelectorTerms: []v1.NodeSelectorTerm{
						{
							MatchExpressions: []v1.NodeSelectorRequirement{
								{
									Key:      "node-type",
									Operator: v1.NodeSelectorOpIn,
									Values:   []string{"worker1"},
								},
							},
						},
					},
				},
			},
		},
	}

	pv.Spec.PersistentVolumeSource = v1.PersistentVolumeSource{
		Local: &v1.LocalVolumeSource{
			Path: "/home/devadmin/storage1",
		},
	}

	createdPV, err := clientset.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created PersistentVolume: %s\n", createdPV.Name)
	return nil
}

// Define Persistent Volume for 2nd Server
func PersistentVolume2(config *rest.Config) error {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: "local-pv-worker2",
		},
		Spec: v1.PersistentVolumeSpec{
			Capacity: v1.ResourceList{
				v1.ResourceStorage: resource.MustParse("30Gi"),
			},
			AccessModes:                   []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
			StorageClassName:              "local-storage-class-worker2",
			// VolumeMode:                    "Filesystem",
			NodeAffinity: &v1.VolumeNodeAffinity{
				Required: &v1.NodeSelector{
					NodeSelectorTerms: []v1.NodeSelectorTerm{
						{
							MatchExpressions: []v1.NodeSelectorRequirement{
								{
									Key:      "node-type",
									Operator: v1.NodeSelectorOpIn,
									Values:   []string{"worker2"},
								},
							},
						},
					},
				},
			},
		},
	}

	pv.Spec.PersistentVolumeSource = v1.PersistentVolumeSource{
		Local: &v1.LocalVolumeSource{
			Path: "/home/devadmin/storage2",
		},
	}

	createdPV, err := clientset.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created PersistentVolume: %s\n", createdPV.Name)
	return nil
}

// Define Persistent Volume for 3rd Server
func PersistentVolume3(config *rest.Config) error {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Define the PersistentVolume specification.
	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: "local-pv-worker3",
		},
		Spec: v1.PersistentVolumeSpec{
			Capacity: v1.ResourceList{
				v1.ResourceStorage: resource.MustParse("30Gi"),
			},
			AccessModes:                   []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
			StorageClassName:              "local-storage-class-worker3",
			// VolumeMode:                    "Filesystem",
			NodeAffinity: &v1.VolumeNodeAffinity{
				Required: &v1.NodeSelector{
					NodeSelectorTerms: []v1.NodeSelectorTerm{
						{
							MatchExpressions: []v1.NodeSelectorRequirement{
								{
									Key:      "node-type",
									Operator: v1.NodeSelectorOpIn,
									Values:   []string{"worker3"},
								},
							},
						},
					},
				},
			},
		},
	}

	pv.Spec.PersistentVolumeSource = v1.PersistentVolumeSource{
		Local: &v1.LocalVolumeSource{
			Path: "/home/devadmin/storage3",
		},
	}

	createdPV, err := clientset.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created PersistentVolume: %s\n", createdPV.Name)
	return nil
}

// Define Persistent Volume for 4th Server
func PersistentVolume4(config *rest.Config) error {
	// Create a Kubernetes clientset.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Define the PersistentVolume specification.
	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: "local-pv-worker4",
		},
		Spec: v1.PersistentVolumeSpec{
			Capacity: v1.ResourceList{
				v1.ResourceStorage: resource.MustParse("30Gi"),
			},
			AccessModes:                   []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
			StorageClassName:              "local-storage-class-worker4",
			// VolumeMode:                    "Filesystem",
			NodeAffinity: &v1.VolumeNodeAffinity{
				Required: &v1.NodeSelector{
					NodeSelectorTerms: []v1.NodeSelectorTerm{
						{
							MatchExpressions: []v1.NodeSelectorRequirement{
								{
									Key:      "node-type",
									Operator: v1.NodeSelectorOpIn,
									Values:   []string{"worker4"},
								},
							},
						},
					},
				},
			},
		},
	}

	// Define the Local volume source using `LocalPath` and the path on the node.
	pv.Spec.PersistentVolumeSource = v1.PersistentVolumeSource{
		Local: &v1.LocalVolumeSource{
			Path: "/home/devadmin/storage4",
		},
	}

	createdPV, err := clientset.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created PersistentVolume: %s\n", createdPV.Name)
	return nil
}
