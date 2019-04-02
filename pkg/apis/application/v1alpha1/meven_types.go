package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MevenSpec defines the desired state of Meven
// +k8s:openapi-gen=true
type MevenSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Replicas           int  `json:"replicas"`
	ApplicationVersion string `json:"applicationVersion"`
}

// MevenStatus defines the observed state of Meven
// +k8s:openapi-gen=true
type MevenStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Replicas           int  `json:"replicas"`
	ApplicationVersion string `json:"applicationVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Meven is the Schema for the mevens API
// +k8s:openapi-gen=true
type Meven struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MevenSpec   `json:"spec,omitempty"`
	Status MevenStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MevenList contains a list of Meven
type MevenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Meven `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Meven{}, &MevenList{})
}
