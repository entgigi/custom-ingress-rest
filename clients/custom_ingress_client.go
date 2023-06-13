package clients

import (
	"context"

	"github.com/entgigi/custom-ingress/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// copied from
// https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html

type CustomIngressInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.CustomIngressList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.CustomIngress, error)
	//Create(*v1alpha1.EntandoBundleV2) (*v1alpha1.EntandoBundleV2, error)
}
type customIngressClient struct {
	restClient rest.Interface
	namespace  string
}

func (bs *customIngressClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.CustomIngress, error) {
	result := v1alpha1.CustomIngress{}
	err := bs.restClient.
		Get().
		Namespace(bs.namespace).
		Resource("entandobundlev2s").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}

func (bs *customIngressClient) List(opts metav1.ListOptions) (*v1alpha1.CustomIngressList, error) {
	result := v1alpha1.CustomIngressList{}
	err := bs.restClient.
		Get().
		Namespace(bs.namespace).
		Resource("entandobundlev2s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)

	return &result, err
}
