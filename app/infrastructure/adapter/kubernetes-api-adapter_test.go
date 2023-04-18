package adapter_test

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"mini-paas/app/infrastructure/adapter"
	"path/filepath"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func TestListPodList(t *testing.T) {
	var kubernetesClient = adapter.KubernetesClient{}
	kubernetesClient.ListPods("")
}

func TestGetApp(t *testing.T) {
	var kubernetesClient = adapter.KubernetesClient{}
	kubernetesClient.GetApp("first-vela-app", "default")
}

func TestCreateApp(t *testing.T) {
	var file = "/Users/fhc/work/MiniPaaS/conf/demo-app.yaml"
	src, err := ioutil.ReadFile(file)
	if err != nil {
	}
	var yamlStr = string(src)
	var kubernetesClient = adapter.KubernetesClient{}
	kubernetesClient.CreateApp(yamlStr, "default")
}

func TestDynamicClient(t *testing.T) {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentRes := schema.GroupVersionResource{Group: "core.oam.dev", Version: "v1beta1", Resource: "applications"}

	var file = "/Users/fhc/work/MiniPaaS/conf/demo-app.yaml"
	src, err := ioutil.ReadFile(file)
	if err != nil {
	}
	var yamlStr = string(src)

	var kubernetesClient = adapter.KubernetesClient{}
	deployment, err := kubernetesClient.YamlToUnstructured(yamlStr)

	// Create Deployment
	fmt.Println("Creating app...")
	result, err := client.Resource(deploymentRes).Namespace("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created app %q.\n", result.GetName())
}
