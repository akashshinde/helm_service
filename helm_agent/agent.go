package helm_agent

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

func GetActionConfigurations() *action.Configuration {
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/akash/kubeconfig")
	// creates the clientset
	clientset, _ := kubernetes.NewForConfig(config)
	store := createStorage("akash-helm-server", clientset)
	conf := &action.Configuration{
		RESTClientGetter: nil,
		Releases:         store,
		KubeClient:       kube.New(nil),
		RegistryClient:   nil,
		Capabilities:     nil,
		Log:              klog.Infof,
	}
	return conf
}

func createStorage(namespace string, clientset *kubernetes.Clientset) *storage.Storage {
	var store *storage.Storage
	d := driver.NewSecrets(clientset.CoreV1().Secrets(namespace))
	d.Log = klog.Infof
	store = storage.Init(d)
	return store
}