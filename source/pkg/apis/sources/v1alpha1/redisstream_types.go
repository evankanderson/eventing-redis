/*
Copyright 2020 The Knative Authors

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
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:defaulter-gen=true

// RedisStreamSource is the Schema for the RedisStream API.
type RedisStreamSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisStreamSourceSpec   `json:"spec,omitempty"`
	Status RedisStreamSourceStatus `json:"status,omitempty"`
}

// Check the interfaces that PingSource should be implementing.
var (
	_ runtime.Object     = (*RedisStreamSource)(nil)
	_ kmeta.OwnerRefable = (*RedisStreamSource)(nil)
	//_ apis.Validatable   = (*RedisStreamSource)(nil)
	//_ apis.Defaultable   = (*RedisStreamSource)(nil)
	_ apis.HasSpec    = (*RedisStreamSource)(nil)
	_ duckv1.KRShaped = (*RedisStreamSource)(nil)
)

// RedisStreamSourceSpec defines the desired state of the RedisStreamSource.
type RedisStreamSourceSpec struct {
	// inherits duck/v1 SourceSpec, which currently provides:
	// * Sink - a reference to an object that will resolve to a domain name or
	//   a URI directly to use as the sink.
	// * CloudEventOverrides - defines overrides to control the output format
	//   and modifications of the event sent to the sink.
	duckv1.SourceSpec `json:",inline"`

	// Stream is the name of the stream.
	Stream string `json:"stream"`
}

// RedisStreamSourceStatus defines the observed state of RedisStreamSource.
type RedisStreamSourceStatus struct {
	// inherits duck/v1 SourceStatus, which currently provides:
	// * ObservedGeneration - the 'Generation' of the Service that was last
	//   processed by the controller.
	// * Conditions - the latest available observations of a resource's current
	//   state.
	// * SinkURI - the current active sink URI that has been configured for the
	//   Source.
	duckv1.SourceStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisStreamSourceList contains a list of RedisStreamSources.
type RedisStreamSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisStreamSource `json:"items"`
}

// GetStatus retrieves the status of the PingSource. Implements the KRShaped interface.
func (p *RedisStreamSource) GetStatus() *duckv1.Status {
	return &p.Status.Status
}
