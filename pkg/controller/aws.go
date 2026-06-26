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

package controller

import (
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane-contrib/provider-aws/pkg/controller/config"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/s3"
	"github.com/crossplane-contrib/provider-aws/pkg/utils/setup"
)

// NOTE(mimacom): This provider has been trimmed to the S3 and IAM service
// surfaces only, as a base for the Crossplane v2 namespaced spike. The
// kms/sns/sqs apis are retained purely as type dependencies of the S3
// reference resolvers; their controllers are intentionally not registered.

// Setup creates all AWS controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	return setup.SetupControllers(
		mgr, o,
		config.Setup,
		iam.Setup,
		s3.Setup,
	)
}
