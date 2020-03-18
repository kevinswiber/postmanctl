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

import "encoding/json"

// UserResponse represents the top-level struct of a user response in the
// Postman API.
type UserResponse struct {
	User User `json:"user"`
}

// Format returns column headers and values for the resource.
func (r UserResponse) Format() ([]string, []interface{}) {
	s := make([]interface{}, 1)
	s[0] = r.User

	return []string{"ID"}, s
}

// User represents the user info associated with a user request in the
// Postman API.
type User struct {
	ID json.Number `json:"id"`
}
