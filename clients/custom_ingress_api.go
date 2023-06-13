package clients

import (
	"github.com/entgigi/custom-ingress/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type CustomIngressV1Alpha1Api interface {
	CustomIngress(ns string) CustomIngressInterface
}

type CustomIngressV1Alpha1Client struct {
	client rest.Interface
}

func NewForConfig(c *rest.Config) (*CustomIngressV1Alpha1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupVersion.Group, Version: v1alpha1.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CustomIngressV1Alpha1Client{client}, nil
}

func (api *CustomIngressV1Alpha1Client) CustomIngress(ns string) CustomIngressInterface {
	return &customIngressClient{
		restClient: api.client,
		namespace:  ns,
	}
}
