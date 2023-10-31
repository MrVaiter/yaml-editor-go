package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func deploymentToYAML(deployment *v1.Deployment) ([]byte, error) {
    deploymentUnstructured, err := runtime.DefaultUnstructuredConverter.ToUnstructured(deployment)
    if err != nil {
        return nil, err
    }

    return yaml.Marshal(deploymentUnstructured)
}

func main() {

	// Read yaml file
	yamlFile, err := os.ReadFile("./test.yaml")
	checkerr(err)

	// Decode yaml file
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode(yamlFile, nil, nil)
	checkerr(err)

	// Assert Deployment type
	deployment := obj.(*v1.Deployment)

	// Add new label
	deployment.Labels["cloud"] = "reef"
	deployment.Spec.Template.Labels["cloud"] = "reef"

	// Convert into yaml
	yamlDeploy, err := deploymentToYAML(deployment)
	checkerr(err)

	err = os.WriteFile("changed.yaml", yamlDeploy, 1001)
	checkerr(err)
}
