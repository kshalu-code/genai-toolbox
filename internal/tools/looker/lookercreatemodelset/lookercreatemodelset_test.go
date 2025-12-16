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

package lookercreatemodelset_test

import (
	"testing"

	yaml "github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
	"github.com/googleapis/genai-toolbox/internal/server"
	"github.com/googleapis/genai-toolbox/internal/testutils"
	"github.com/googleapis/genai-toolbox/internal/tools/looker/lookercreatemodelset"
)

func TestParseFromYamlLookerCreateModelSet(t *testing.T) {
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
  create_model_set:
    kind: looker-create-model-set
    source: looker-source
    description: Create a Looker model set
`,
			want: server.ToolConfigs{
				"create_model_set": lookercreatemodelset.Config{
					Name:         "create_model_set",
					Kind:         "looker-create-model-set",
					Source:       "looker-source",
					Description:  "Create a Looker model set",
					AuthRequired: []string{},
				},
			},
		},
		{
			desc: "with auth",
			in: `
tools:
  create_model_set:
    kind: looker-create-model-set
    source: looker-source
    description: Create a Looker model set
    authRequired:
      - google-auth-service
`,
			want: server.ToolConfigs{
				"create_model_set": lookercreatemodelset.Config{
					Name:         "create_model_set",
					Kind:         "looker-create-model-set",
					Source:       "looker-source",
					Description:  "Create a Looker model set",
					AuthRequired: []string{"google-auth-service"},
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

func TestFailParseFromYamlLookerCreateModelSet(t *testing.T) {
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
  create_model_set:
    source: looker-source
    description: Create a Looker model set
`,
		},
		{
			desc: "missing source",
			in: `
tools:
  create_model_set:
    kind: looker-create-model-set
    description: Create a Looker model set
`,
		},
		{
			desc: "missing description",
			in: `
tools:
  create_model_set:
    kind: looker-create-model-set
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
				// Validation happens at a higher level usually, but UnmarshalContext with validate tag should catch it if configured.
				// However, Genai toolbox uses a validator after unmarshaling.
				// For the purpose of these unit tests, we are checking if the Config struct itself can be unmarshaled.
				// If validation is strictly required in unit tests here, we'd need to call a validator.
				if _, ok := got.Tools["create_model_set"].(lookercreatemodelset.Config); !ok {
					t.Fatal("expected lookercreatemodelset.Config")
				}

				// Standard toolbox pattern: Config uses `validate:"required"` tags.
				// These are validated during UnmarshalYAML if using NewStrictDecoder,
				// which server.ToolConfigs does.
				// If we reached here, it means it didn't fail when it should have,
				// or we didn't use the strict unmarshaler correctly in the test.
				// Re-testing with strict unmarshaler if needed, but for now just fail if no error.
				t.Fatalf("expected error for %s", tc.desc)
			}
		})
	}
}
