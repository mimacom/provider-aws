/*
Copyright 2023 The Crossplane Authors.

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

package managed

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// TestUpdateCriticalAnnotationsScopes verifies the updater persists annotations
// for BOTH namespaced and cluster-scoped objects. The namespaced case is a
// regression test: the lookup key must carry the object's namespace, otherwise
// the Get resolves against the empty namespace and returns NotFound — which
// wedges namespaced (Crossplane v2) managed resources on
// "cannot determine creation result".
func TestUpdateCriticalAnnotationsScopes(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		t.Fatalf("add corev1 to scheme: %v", err)
	}

	cases := map[string]struct {
		existing client.Object
		update   client.Object
		getKey   client.ObjectKey
		into     client.Object
	}{
		"Namespaced": {
			// A namespaced object (stands in for a v2 managed resource).
			existing: &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{Namespace: "mabu", Name: "mybucket"},
			},
			update: &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "mabu",
					Name:        "mybucket",
					Annotations: map[string]string{"crossplane.io/external-create-succeeded": "now"},
				},
			},
			getKey: client.ObjectKey{Namespace: "mabu", Name: "mybucket"},
			into:   &corev1.ConfigMap{},
		},
		"ClusterScoped": {
			// A cluster-scoped object (stands in for a legacy cluster MR).
			existing: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{Name: "myrole"},
			},
			update: &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "myrole",
					Annotations: map[string]string{"crossplane.io/external-create-succeeded": "now"},
				},
			},
			getKey: client.ObjectKey{Name: "myrole"},
			into:   &corev1.Namespace{},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(tc.existing).Build()
			u := NewRetryingCriticalAnnotationUpdater(c)

			if err := u.UpdateCriticalAnnotations(context.Background(), tc.update); err != nil {
				t.Fatalf("UpdateCriticalAnnotations() returned error: %v", err)
			}

			if err := c.Get(context.Background(), tc.getKey, tc.into); err != nil {
				t.Fatalf("Get after update: %v", err)
			}
			if got := tc.into.GetAnnotations()["crossplane.io/external-create-succeeded"]; got != "now" {
				t.Errorf("annotation not persisted: got %q, want %q", got, "now")
			}
		})
	}
}
