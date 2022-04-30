/*
Copyright © 2021 OpenSLO Team

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
package fmt

import (
	"bytes"
	"testing"
)

func Test_fmtFile(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			name: "Invalid file",
			args: args{
				source: "../../../test/invalid-file.yaml",
			},
			wantErr: true,
			wantOut: "",
		},
		{
			name: "Invalid content",
			args: args{
				source: "../../../test/invalid-service.yaml",
			},
			wantErr: true,
			wantOut: "",
		},
		{
			name: "Passes",
			args: args{
				source: "../../../test/valid-service.yaml",
			},
			wantErr: false,
			wantOut: `apiVersion: openslo/v1alpha
kind: Service
metadata:
  name: my-rad-service
  displayName: My Rad Service
spec:
  description: This is a great description of an even better service.
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := fmtFile(out, tt.args.source); (err != nil) != tt.wantErr {
				t.Errorf("fmtFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("fmtFile() = %v, want %v", gotOut, tt.wantOut)
			}

		})
	}
}
