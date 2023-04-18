package adapter_test

import (
	"mini-paas/app/infrastructure/adapter"
	"testing"
)

func TestListPodList(t *testing.T) {
	var kubernetesClient = adapter.KubernetesClient{}
	kubernetesClient.ListPods("")
}

func TestApplyYaml(t *testing.T) {
	// var file = "/Users/fhc/work/MiniPaaS/conf/demo-app.yaml"
	// src, err := ioutil.ReadFile(file)
	// if err != nil {
	// }
	// var yamlStr = string(src)
	var kubernetesClient = adapter.KubernetesClient{}
	kubernetesClient.GetApp("first-vela-app", "default")
}
