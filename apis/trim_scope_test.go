/*
Copyright 2026 The Crossplane Authors.

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

package apis

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// allowedAPIPackages is the exact set of top-level directories permitted
// under apis/ for the mimacom DATEV S3+IAM trim (DCS-8/17/18). kms, sns and
// sqs are kept only as type dependencies of S3 reference resolvers (see the
// NOTE(mimacom) comment in apis/aws.go), not as active service surface.
//
// If this test fails because upstream introduced a new top-level apis/
// package, it means a new upstream service was NOT caught by the trim
// commit and has silently entered the tree. Do not just add it to this
// list — first check whether it needs deleting from the release patch, the
// same way f615da3a deleted the others.
var allowedAPIPackages = map[string]bool{
	"common":   true, // apis/common — shared s3/bucketpolicy types
	"iam":      true,
	"kms":      true, // type dependency only, see NOTE(mimacom) in aws.go
	"s3":       true,
	"sns":      true, // type dependency only, see NOTE(mimacom) in aws.go
	"sqs":      true, // type dependency only, see NOTE(mimacom) in aws.go
	"v1beta1":  true, // top-level ProviderConfig (legacy, cluster-scoped)
	"v1beta1m": true, // top-level ProviderConfig (namespaced)
}

// TestTrimScopeNoNewServicePackages fails if a directory appears under
// apis/ that isn't part of the DATEV S3+IAM trim. This is the automated
// tripwire for "upstream added a whole new service and our permanent patch
// never deleted it" — see runbook-provider-aws-personalizaciones-permanentes.md.
func TestTrimScopeNoNewServicePackages(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("cannot read apis/ directory: %v", err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		if !allowedAPIPackages[name] {
			t.Errorf("unexpected package apis/%s: upstream likely added a new "+
				"service that the DCS-8/17/18 trim patch does not delete. "+
				"Extend the permanent patch to remove it (or, if it must be "+
				"kept, update allowedAPIPackages deliberately) before releasing.", name)
		}
	}
}

// crossplaneRuntimeV1Import matches an import of crossplane-runtime v1 (no
// /v2/ in the path) — the exact combination that caused the protobuf
// registry panic this trim exists to avoid (two copies of the changelog
// proto registered in the same process). See DCS-8/17/18 and the
// provider-aws-datev runbook.
var crossplaneRuntimeV1Import = regexp.MustCompile(`"github\.com/crossplane/crossplane-runtime/(apis|pkg)/`)

// TestNoLegacyRuntimeImports walks every .go file under the module and
// fails if any of them import crossplane-runtime v1. A new file added
// upstream inside an already-kept package (apis/iam, apis/s3, ...) could in
// principle still import v1 if upstream itself hasn't migrated to v2 yet —
// this catches that case even though the file itself isn't new enough to
// show up as a whole new "apis/" package under TestTrimScopeNoNewServicePackages.
func TestNoLegacyRuntimeImports(t *testing.T) {
	root, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("cannot resolve module root: %v", err)
	}

	var offenders []string
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			base := info.Name()
			if base == "tools" || base == ".git" || strings.HasPrefix(base, ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if crossplaneRuntimeV1Import.Match(data) {
			rel, _ := filepath.Rel(root, path)
			offenders = append(offenders, rel)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("cannot walk module tree: %v", err)
	}

	for _, f := range offenders {
		t.Errorf("%s imports crossplane-runtime v1 — this provider must be "+
			"exclusively on crossplane-runtime/v2 (mixing v1 and v2 in the same "+
			"process re-triggers the protobuf changelog registry panic that "+
			"motivated the DCS-8/17/18 permanent patch)", f)
	}
}
