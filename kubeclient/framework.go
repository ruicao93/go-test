package kubeclient

import (
	"context"
	"flag"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"testing"
)

var (
	homedir, _     = os.UserHomeDir()
	kubeconfigPath = flag.String("kubeconfig", `D:\bugs\loadbalancer\ovs-np.kubeconf`, "Path of kubeconfig file")
)

type TestData struct {
	kubeconfigPath string
	kubeconfig     *restclient.Config
	clientset      kubernetes.Interface
}

func (data *TestData) createClient() error {
	data.kubeconfigPath = *kubeconfigPath

	config, err := clientcmd.BuildConfigFromFlags("", data.kubeconfigPath)
	if err != nil {
		return err
	}
	data.kubeconfig = config
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	data.clientset = clientSet
	return nil
}

func (data *TestData) createClien2t() error {
	data.kubeconfigPath = *kubeconfigPath

	loadingrules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingrules.ExplicitPath = data.kubeconfigPath
	configOverrides := &clientcmd.ConfigOverrides{}

	kubeconfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingrules, configOverrides).ClientConfig()
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return err
	}
	data.kubeconfig = kubeconfig
	data.clientset = clientset
	return nil
}

func (data *TestData) createService(serviceName, namespace string, port, targetPort int, serviceType v1.ServiceType, selector map[string]string,
	ingressIPs []string) (*v1.Service, error) {
	ingress := make([]v1.LoadBalancerIngress, len(ingressIPs))
	for idx, ingressIP := range ingressIPs {
		ingress[idx].IP = ingressIP
	}
	service := v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: namespace,
			Labels: map[string]string{
				"antrea-e2e": serviceName,
				"app":        serviceName,
			},
		},
		Spec: v1.ServiceSpec{
			SessionAffinity: v1.ServiceAffinityNone,
			Ports: []v1.ServicePort{{
				Port:       int32(port),
				TargetPort: intstr.FromInt(targetPort),
			}},
			Type:     serviceType,
			Selector: selector,
		},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{Ingress: ingress},
		},
	}
	return data.clientset.CoreV1().Services(namespace).Create(context.TODO(), &service, metav1.CreateOptions{})
}

func setupTest(tb testing.TB) (*TestData, error) {
	flag.Parse()
	data := &TestData{}
	if err := data.createClient(); err != nil {
		return nil, err
	}
	return data, nil
}
