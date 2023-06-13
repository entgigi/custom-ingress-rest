package controllers

import (
	"fmt"
	"net/http"

	"github.com/entgigi/custom-ingress-rest/clients"
	"github.com/entgigi/custom-ingress-rest/utilities"
	"github.com/entgigi/custom-ingress/api/v1alpha1"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type CustomIngressDto struct {
	Name        string `json:"name"`
	IngressName string `json:"ingressName"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Path        string `json:"path"`
	Service     string `json:"service"`
}

type IngressCtrl struct {
	apiClient *clients.CustomIngressV1Alpha1Client
}

func NewIngressCtrl(config *rest.Config) (*IngressCtrl, error) {
	s, err := clients.NewForConfig(config)
	return &IngressCtrl{s}, err
}

func (bc *IngressCtrl) ListIngresses(ctx *gin.Context) {

	ns, _ := utilities.GetWatchNamespace()
	bundles, err := bc.apiClient.CustomIngress(ns).List(metav1.ListOptions{})
	fmt.Println("ListBundles")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		//log.Infof(context.TODO(), "found number of bundles: %n", len(bundles.Items))
		ctx.JSON(http.StatusOK, gin.H{"data": converter(bundles.Items)})
	}
}

func converter(inCustomIngresses []v1alpha1.CustomIngress) []CustomIngressDto {
	outIngresses := make([]CustomIngressDto, 0)
	for _, cIngress := range inCustomIngresses {
		b := CustomIngressDto{
			Name:    cIngress.GetName(),
			Host:    cIngress.Spec.Name,
			Path:    cIngress.Spec.Path,
			Port:    cIngress.Spec.Port,
			Service: cIngress.Spec.Service,
		}
		outIngresses = append(outIngresses, b)
	}
	return outIngresses
}
