package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	controller "github.com/deployment_engine/controller"
)

func main() {
	keysString := os.Getenv("privKeys1")
	PrivateKeys := strings.Split(keysString, ",")

	kubeconfig := flag.String("kubeconfig", "/home/devadmin/.kube/config", "location to yout kubeconfig file")
	flag.Parse()
	kubeconfigPath := *kubeconfig
	r := controller.Router(PrivateKeys, kubeconfigPath)
	fmt.Println("Port is Running on=> 8000")
	r.Run(":8000")
}
