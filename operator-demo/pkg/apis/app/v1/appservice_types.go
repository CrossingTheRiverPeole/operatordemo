package v1

import (
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppServiceSpec defines the desired state of AppService
type AppServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Size      int32                   `json:"size"`
	Image     string                  `json:"image"`
	Resources v1.ResourceRequirements `json:"resource,omitempty"`
	Envs      []v1.EnvVar             `json:"envs,omitempty"`
	Ports     []v1.ServicePort        `json:"ports,omitempty"`
	Commands  []string                `json:"commands,omitempty"`
	Args      []string                `json:"args,omitempty"`
}

// AppServiceStatus defines the observed state of AppService
type AppServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	v12.DeploymentStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppService is the Schema for the appservices API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=appservices,scope=Namespaced
type AppService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppServiceSpec   `json:"spec,omitempty"`
	Status AppServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppServiceList contains a list of AppService
type AppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppService{}, &AppServiceList{})
}
