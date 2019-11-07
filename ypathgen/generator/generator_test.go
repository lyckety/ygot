// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ygot/ypathgen"
)

func TestWritePathCode(t *testing.T) {
	tests := []struct {
		name string
		in   *ypathgen.GeneratedPathCode
		want string
	}{{
		name: "simple",
		in: &ypathgen.GeneratedPathCode{
			CommonHeader: "path common header\n",
			OneOffHeader: "\npath one-off header\n",
			Structs: []ypathgen.GoPathStructCodeSnippet{{
				PathStructName:    "PathStructName",
				StructBase:        "\nStructDef\n",
				ChildConstructors: "\nChildConstructor\n",
			}},
		},
		want: `path common header

path one-off header

StructDef

ChildConstructor
`,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			if err := writeGoCodeSingleFile(&b, tt.in); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, b.String()); diff != "" {
				t.Errorf("diff (-want,+got):\n%s", diff)
			}
		})
	}
}