package adapter

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/kubectl/pkg/scheme"
)

type KubernetesClient struct {
}

func (k KubernetesClient) GetApp(appName string, namespace string) string {
	//TODO
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/fhc/kube/config")
	if err != nil {
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	data, err := client.RESTClient().Get().AbsPath("/apis/core.oam.dev/v1beta1/namespaces/" + namespace + "/applications/" + appName).DoRaw(context.TODO())
	yamlString, err := bytesToYAML(data)
	if err != nil {
		panic(err)
	}
	return yamlString
}
func (k KubernetesClient) CreateApp(appYaml string, namespace string) string {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentRes := schema.GroupVersionResource{Group: "core.oam.dev", Version: "v1beta1", Resource: "applications"}

	kk := KubernetesClient{}
	deployment, err := kk.YamlToUnstructured(appYaml)

	// Create Deployment
	fmt.Println("Creating app...")
	result, err := client.Resource(deploymentRes).Namespace("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created app %q.\n", result.GetName())
	return result.GetName()
}

func bytesToYAML(b []byte) (string, error) {
	var data map[interface{}]interface{}
	err := yaml.Unmarshal(b, &data)
	if err != nil {
		return "", err
	}
	res, err := yaml.Marshal(&data)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (k KubernetesClient) YamlToUnstructured(yamlStr string) (*unstructured.Unstructured, error) {
	dec := json.NewYAMLSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme)
	obj := &unstructured.Unstructured{}
	_, gvk, err := dec.Decode([]byte(yamlStr), nil, obj)
	if err != nil {
		return nil, err
	}
	obj.SetGroupVersionKind(*gvk)
	return obj, nil
}

func (k KubernetesClient) ListPods(namespace string) {
	// 集群内
	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// 配置认证信息和 API 服务器地址
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/fhc/kube/config")
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 获取 Pod 对象列表
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 输出 Pod 对象列表
	fmt.Printf("There are %d pods in the cluster:\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("  %s\n", pod.ObjectMeta.Name)
	}
}
