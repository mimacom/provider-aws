/*
Copyright 2019 The Crossplane Authors.

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

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

// RoleARN returns the status.atProvider.ARN of a Role.
func RoleARN() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		r, ok := mg.(*Role)
		if !ok {
			return ""
		}
		return r.Status.AtProvider.ARN

	}
}

// NOTE(mimacom): PolicyARN and UserARN moved to the namespaced
// iam.aws.m.crossplane.io package (apis/iam/v1beta1m) together with the
// Policy and User types. The cluster-scoped group/role policy attachments and
// group-user memberships no longer auto-resolve their Policy/User references
// (they take explicit policyArn/userName values), because cross-scope
// (cluster -> namespaced) reference resolution is not supported.
