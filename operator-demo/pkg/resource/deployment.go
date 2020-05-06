package resource

import (
	"fmt"
	v1 "github.com/CrossingTheRiverPeole/operator-demo/pkg/apis/app/v1"
	v12 "k8s.io/api/apps/v1"
	v14 "k8s.io/api/core/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewDeploy(app *v1.AppService) *v12.Deployment {
	labels := map[string]string{"app": app.Name}
	selectors := &v13.LabelSelector{MatchLabels: labels}
	return &v12.Deployment{
		TypeMeta: v13.TypeMeta{
			APIVersion: "apps.v1",
			Kind:       "Deployment",
		},
		ObjectMeta: v13.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			OwnerReferences: []v13.OwnerReference{
				*v13.NewControllerRef(
					app, schema.GroupVersionKind{
						Group:   v1.SchemeGroupVersion.Group,
						Version: v1.SchemeGroupVersion.Version,
						Kind:    "AppService",
					}),
			},
		},
		Spec: v12.DeploymentSpec{
			Replicas: &app.Spec.Size,
			Template: v14.PodTemplateSpec{
				ObjectMeta: v13.ObjectMeta{
					Labels: labels,
				},
				Spec: v14.PodSpec{
					Containers: newContainers(app),

					//Containers: nil,
				},
			},
			Selector: selectors,
		},
		//Status: v12.DeploymentStatus{},
	}

}

func newContainers(app *v1.AppService) []v14.Container {
	containerPorts := []v14.ContainerPort{}
	for _, svcPort := range app.Spec.Ports {
		cport := v14.ContainerPort{}
		cport.ContainerPort = svcPort.TargetPort.IntVal
		containerPorts = append(containerPorts, cport)
	}

	fmt.Print("------------", "command is ", app.Spec.Commands)

	return []v14.Container{
		{
			Name:            app.Name,
			Image:           app.Spec.Image,
			Resources:       app.Spec.Resources,
			Ports:           containerPorts,
			ImagePullPolicy: v14.PullIfNotPresent,
			Env:             app.Spec.Envs,
			Command:         app.Spec.Commands,
			Args:            app.Spec.Args,
		},
	}
}
