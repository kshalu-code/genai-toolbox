// Copyright 2024 Google LLC
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

package lookersearchroles_test

import (
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
	"github.com/googleapis/genai-toolbox/internal/server"
	"github.com/googleapis/genai-toolbox/internal/testutils"
	"github.com/googleapis/genai-toolbox/internal/tools/looker/lookersearchroles"
)

func TestParseFromYamlLookerSearchRoles(t *testing.T) {
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
  search_roles:
    kind: looker-search-roles
    source: my-instance
    description: Search for roles in Looker.
`,
			want: server.ToolConfigs{
				"search_roles": lookersearchroles.Config{
					Name:         "search_roles",
					Kind:         "looker-search-roles",
					Source:       "my-instance",
					Description:  "Search for roles in Looker.",
					AuthRequired: []string{},
				},
			},
		},
		{
			desc: "with auth",
			in: `
tools:
  search_roles:
    kind: looker-search-roles
    source: my-instance
    description: Search for roles in Looker.
    authRequired:
      - looker
`,
			want: server.ToolConfigs{
				"search_roles": lookersearchroles.Config{
					Name:         "search_roles",
					Kind:         "looker-search-roles",
					Source:       "my-instance",
					Description:  "Search for roles in Looker.",
					AuthRequired: []string{"looker"},
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

func TestFailParseFromYamlLookerSearchRoles(t *testing.T) {
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
  search_roles:
    source: my-instance
    description: Search for roles in Looker.
`,
		},
		{
			desc: "missing source",
			in: `
tools:
  search_roles:
    kind: looker-search-roles
    description: Search for roles in Looker.
`,
		},
		{
			desc: "missing description",
			in: `
tools:
  search_roles:
    kind: looker-search-roles
    source: my-instance
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
				t.Fatalf("expected error for %s", tc.desc)
			}
		})
	}
}
