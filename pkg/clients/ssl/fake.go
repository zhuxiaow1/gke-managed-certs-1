/*
Copyright 2020 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ssl

import (
	compute "google.golang.org/api/compute/v0.beta"
)

func NewFakeQuotaExceededError() *Error {
	return newFakeError(codeQuotaExceeded)
}

func newFakeError(code string) *Error {
	return &Error{
		operation: &compute.Operation{
			Error: &compute.OperationError{
				Errors: []*compute.OperationErrorErrors{{Code: code}},
			},
		},
	}
}