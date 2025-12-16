// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lookercreaterole_test

import (
	"testing"

	yaml "github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
	"github.com/googleapis/genai-toolbox/internal/server"
	"github.com/googleapis/genai-toolbox/internal/testutils"
	"github.com/googleapis/genai-toolbox/internal/tools/looker/lookercreaterole"
)

func TestParseFromYamlLookerCreateRole(t *testing.T) {
	ctx, err := testutils.ContextWithNewLogger()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	tcs := []struct {
		desc string
		in   string
		want server.ToolConfigs
	}{
		{
			desc: "basic example",
			in: `
tools:
  create_role:
    kind: looker-create-role
    source: looker-source
    description: Create a Looker role
`,
			want: server.ToolConfigs{
				"create_role": lookercreaterole.Config{
					Name:         "create_role",
					Kind:         "looker-create-role",
					Source:       "looker-source",
					Description:  "Create a Looker role",
					AuthRequired: []string{},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := struct {
				Tools server.ToolConfigs `yaml:"tools"`
			}{}

			err := yaml.UnmarshalContext(ctx, testutils.FormatYaml(tc.in), &got)
			if err != nil {
				t.Fatalf("unable to unmarshal: %s", err)
			}

			if diff := cmp.Diff(tc.want, got.Tools); diff != "" {
				t.Fatalf("incorrect parse: diff %v", diff)
			}
		})
	}
}

func TestFailParseFromYamlLookerCreateRole(t *testing.T) {
	ctx, err := testutils.ContextWithNewLogger()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	tcs := []struct {
		desc string
		in   string
	}{
		{
			desc: "missing kind",
			in: `
tools:
  create_role:
    source: looker-source
    description: Create a Looker role
`,
		},
		{
			desc: "missing source",
			in: `
tools:
  create_role:
    kind: looker-create-role
    description: Create a Looker role
`,
		},
		{
			desc: "missing description",
			in: `
tools:
  create_role:
    kind: looker-create-role
    source: looker-source
`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := struct {
				Tools server.ToolConfigs `yaml:"tools"`
			}{}

			err := yaml.UnmarshalContext(ctx, testutils.FormatYaml(tc.in), &got)
			if err == nil {
				if _, ok := got.Tools["create_role"].(lookercreaterole.Config); !ok {
					t.Fatal("expected lookercreaterole.Config")
				}
				t.Fatalf("expected error for %s", tc.desc)
			}
		})
	}
}
