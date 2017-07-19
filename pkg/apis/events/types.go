/*
Copyright 2017 The Kubernetes Authors.

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

package events

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// a list of values for a given metric for some set of objects
type EventValueList struct {
	
	Items []EventValue `json:"items"`
}

// a metric value for some object
type EventValue struct {
	
	// the name of the event
	EventName string `json:"name"`

	// kind of the event
	EventKind string `json:"kind"`

	// reason of the event
	Reason string `json:"reason"`

	// source of the event
	Source string `json:"source"`

}

// allObjects is a wildcard used to select metrics
// for all objects matching the given label selector
const AllObjects = "*"

// NOTE: ObjectReference is copied from k8s.io/kubernetes/pkg/api/types.go. We
// cannot depend on k8s.io/kubernetes/pkg/api because that creates cyclic
// dependency between k8s.io/metrics and k8s.io/kubernetes. We cannot depend on
// k8s.io/client-go/pkg/api because the package is going to be deprecated soon.
// There is no need to keep it an exact copy. Each repo can define its own
// internal objects.

// ObjectReference contains enough information to let you inspect or modify the referred object.
type ObjectReference struct {
	Kind            string
	Namespace       string
	Name            string
	UID             types.UID
	APIVersion      string
	ResourceVersion string
	FieldPath       string
}