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

package sdk

import (
	"context"
	"encoding/json"
)

// RunMonitor runs a Postman monitor.
func (s *Service) RunMonitor(ctx context.Context, id string) (json.RawMessage, error) {
	var responseBody json.RawMessage
	if _, err := s.post(ctx, nil, &responseBody, nil, "monitors", id, "run"); err != nil {
		return nil, err
	}

	return responseBody, nil
}
