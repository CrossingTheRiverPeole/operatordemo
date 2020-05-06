package resource

import (
	v1 "github.com/CrossingTheRiverPeole/operator-demo/pkg/apis/app/v1"
	v14 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewService(app *v1.AppService) *v14.Service {
	return &v14.Service{
		TypeMeta: v12.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: v12.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			OwnerReferences: []v12.OwnerReference{
				*v12.NewControllerRef(app, schema.GroupVersionKind{
					Group:   v1.SchemeGroupVersion.Group,
					Version: v1.SchemeGroupVersion.Version,
					Kind:    "AppService",
				}),
			},
		},
		Spec: v14.ServiceSpec{
			Type:  v14.ServiceTypeNodePort,
			Ports: app.Spec.Ports,
			Selector: map[string]string{
				"app": app.Name,
			},
		},
	}

}
