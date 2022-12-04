/*
Copyright 2022.

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

// GitHttpServerSpec defines the desired state of GitHttpServer
type GitHttpServerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of GitHttpServer. Edit githttpserver_types.go to remove/update
	Image        string `json:"image,omitempty"`
	PathClone    string `json:"path-clone,omitempty"`
	PathPull     string `json:"path-pull,omitempty"`
	PathVersion  string `json:"path-version,omitempty"`
	PathWebHook  string `json:"path-web-hook,omitempty"`
	PathHealth   string `json:"path-health,omitempty"`
	RepoURL      string `json:"repo-url"`
	RepoBranch   string `json:"repo-branch,omitempty"`
	RepoTarget   string `json:"repo-target,omitempty"`
	RepoUsername string `json:"repo-username,omitempty"`
	RepoPassword string `json:"repo-password,omitempty"`
	HttpPort     int32  `json:"http-port,omitempty"`
}

// GitHttpServerStatus defines the observed state of GitHttpServer
type GitHttpServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GitHttpServer is the Schema for the githttpservers API
type GitHttpServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitHttpServerSpec   `json:"spec,omitempty"`
	Status GitHttpServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GitHttpServerList contains a list of GitHttpServer
type GitHttpServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitHttpServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitHttpServer{}, &GitHttpServerList{})
}
