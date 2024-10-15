package assets

import (
	"embed"
	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	//go:embed manifests
	manifests embed.FS

	assetsScheme = runtime.NewScheme()
	assetsCodecs = serializer.NewCodecFactory(assetsScheme)
)

func init() {
	utilruntime.Must(corev1.AddToScheme(assetsScheme))
	utilruntime.Must(appsv1.AddToScheme(assetsScheme))
}

func GetPersistentVolumeClaimFromFile(name string) (*corev1.PersistentVolumeClaim, error) {
	pvcDataBytes, err := manifests.ReadFile(name)
	if err != nil {
		return nil, err
	}
	pvcDataObject, err := runtime.Decode(assetsCodecs.UniversalDecoder(corev1.SchemeGroupVersion), pvcDataBytes)
	if err != nil {
		return nil, err
	}
	pvcData := pvcDataObject.(*corev1.PersistentVolumeClaim)
	return pvcData, nil
}

func GetDeploymentmFromFile(name string) (*appsv1.Deployment, error) {
	deploymentBytes, err := manifests.ReadFile(name)
	if err != nil {
		return nil, err
	}
	deploymentObject, err := runtime.Decode(assetsCodecs.UniversalDecoder(appsv1.SchemeGroupVersion), deploymentBytes)
	if err != nil {
		return nil, err
	}
	deploy := deploymentObject.(*appsv1.Deployment)
	return deploy, nil
}
