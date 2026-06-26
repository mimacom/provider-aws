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

// Package apis contains Kubernetes API groups for AWS cloud provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	iamv1alpha1 "github.com/crossplane-contrib/provider-aws/apis/iam/v1alpha1"
	iamv1beta1 "github.com/crossplane-contrib/provider-aws/apis/iam/v1beta1"
	kmsv1alpha1 "github.com/crossplane-contrib/provider-aws/apis/kms/v1alpha1"
	s3v1alpha2 "github.com/crossplane-contrib/provider-aws/apis/s3/v1alpha3"
	s3v1beta1 "github.com/crossplane-contrib/provider-aws/apis/s3/v1beta1"
	snsv1beta1 "github.com/crossplane-contrib/provider-aws/apis/sns/v1beta1"
	sqsv1beta1 "github.com/crossplane-contrib/provider-aws/apis/sqs/v1beta1"
	awsv1alpha1 "github.com/crossplane-contrib/provider-aws/apis/v1alpha1"
	awsv1beta1 "github.com/crossplane-contrib/provider-aws/apis/v1beta1"
)

// NOTE(mimacom): This provider has been trimmed to the S3 and IAM service
// surfaces only, as a base for the Crossplane v2 namespaced spike. The
// kms/sns/sqs apis are retained purely as type dependencies of the S3
// reference resolvers, but are still registered so the scheme can resolve
// cross-resource references.

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		iamv1alpha1.SchemeBuilder.AddToScheme,
		iamv1beta1.SchemeBuilder.AddToScheme,
		kmsv1alpha1.SchemeBuilder.AddToScheme,
		s3v1alpha2.SchemeBuilder.AddToScheme,
		s3v1beta1.SchemeBuilder.AddToScheme,
		snsv1beta1.SchemeBuilder.AddToScheme,
		sqsv1beta1.SchemeBuilder.AddToScheme,
		awsv1alpha1.SchemeBuilder.AddToScheme,
		awsv1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
