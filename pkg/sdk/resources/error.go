/*
Copyright © 2020 Kevin Swiber <kswiber@gmail.com>

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

package resources

import (
	"fmt"
)

// ErrorResponse is the struct representation of a Postman API error.
type ErrorResponse struct {
	Error Error `json:"error"`
}

// Error is a struct representation of error details from a Postman API error.
type Error struct {
	Name    string                 `json:"name"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

func (e *Error) String() string {
	return fmt.Sprintf("name: %s, message: %s", e.Name, e.Message)
}
