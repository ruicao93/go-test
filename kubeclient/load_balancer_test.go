package kubeclient

import (
	"context"
	"encoding/json"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLoadBalancer(t *testing.T) {
	data, err := setupTest(t)
	if err != nil {
		t.Fatalf("%v", err)
	}
	service, err := data.clientset.CoreV1().Services("default").Get(context.TODO(), "win-webserver", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("Service: %v", service)

	if len(service.Status.LoadBalancer.Ingress) > 0 {
		t.Logf("Ingress IP alread allocated: %v, skip allocation.", service.Status.LoadBalancer.Ingress)
		return
	}

	var ingress v1.LoadBalancerIngress
	ingress.IP = "192.168.33.33"

	updated := service.DeepCopy()
	updated.Status.LoadBalancer.Ingress = []v1.LoadBalancerIngress{ingress}
	patchData, err := json.Marshal(updated)
	if err != nil {
		t.Fatalf("%v", err)
	}
	result, err := data.clientset.CoreV1().Services("default").Patch(context.TODO(), "win-webserver", types.MergePatchType, patchData, metav1.PatchOptions{}, "status")
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("Updated service: %v", result)
}

func TestCreateService(t *testing.T) {
	data, err := setupTest(t)
	if err != nil {
		t.Fatalf("%v", err)
	}
	svc, err := data.createService("test-create-service", "default", 80, 8080, v1.ServiceTypeLoadBalancer, make(map[string]string), []string{"192.168.22.22"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("Created service: %v", svc)
}

func TestListNodes(t *testing.T) {
	data, err := setupTest(t)
	if err != nil {
		t.Fatalf("%v", err)
	}
	nodes, err := data.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		t.Fatalf("%v", err)
	}
	for _, node := range nodes.Items {
		t.Logf("%v", node.Name)
	}
}
