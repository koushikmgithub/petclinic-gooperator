/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PetclinicGoSpec defines the desired state of PetclinicGo
type PetclinicGoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size defines the number of PetclinicGo instances
	Size int32 `json:"size,omitempty"`

	// Image defines the docker image of PetclinicGo instances
	Image string `json:"image,omitempty"`

	// Foo is an example field of PetclinicGo. Edit petclinicgo_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// PetclinicGoStatus defines the observed state of PetclinicGo
type PetclinicGoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Nodes store the name of the pods which are running PetclinicGo instances
	Nodes []string `json:"nodes,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PetclinicGo is the Schema for the petclinicgoes API
type PetclinicGo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PetclinicGoSpec   `json:"spec,omitempty"`
	Status PetclinicGoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PetclinicGoList contains a list of PetclinicGo
type PetclinicGoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PetclinicGo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PetclinicGo{}, &PetclinicGoList{})
}
