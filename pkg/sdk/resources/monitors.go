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
	"encoding/json"
	"time"
)

// MonitorListResponse represents the top-level monitors response from the
// Postman API.
type MonitorListResponse struct {
	Monitors MonitorListItems `json:"monitors"`
}

// MonitorListItems is a slice of MonitorListItem.
type MonitorListItems []MonitorListItem

// Format returns column headers and values for the resource.
func (r MonitorListItems) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r))
	for i, v := range r {
		s[i] = v
	}

	return []string{"ID", "Name"}, s
}

// MonitorListItem represents a single item in an MonitorListResponse.
type MonitorListItem struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	UID   string      `json:"uid"`
	Owner json.Number `json:"owner"`
}

// MonitorResponse is the top-level monitor response from the
// Postman API.
type MonitorResponse struct {
	Monitor Monitor `json:"monitor"`
}

// Monitor represents the single monitor response from the
// Postman API
type Monitor struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	UID            string         `json:"uid"`
	Owner          int            `json:"owner"`
	CollectionUID  string         `json:"collectionUid"`
	EnvironmentUID string         `json:"environmentUid"`
	Options        MonitorOptions `json:"options"`
	Notifications  Notifications  `json:"notifications"`
	Distribution   []interface{}  `json:"distribution"`
	Schedule       Schedule       `json:"schedule"`
}

// Format returns column headers and values for the resource.
func (r Monitor) Format() ([]string, []interface{}) {
	s := make([]interface{}, 1)
	s[0] = r

	return []string{"ID", "Name"}, s
}

// MonitorSlice is a slice of Monitor.
type MonitorSlice []*Monitor

// Format returns column headers and values for the resource.
func (r MonitorSlice) Format() ([]string, []interface{}) {
	s := make([]interface{}, len(r))
	for i, v := range r {
		s[i] = v
	}

	return []string{"ID", "Name"}, s
}

// MonitorOptions list options for a monitor.
type MonitorOptions struct {
	StrictSSL       bool        `json:"strictSSL"`
	FollowRedirects bool        `json:"followRedirects"`
	RequestTimeout  interface{} `json:"requestTimeout"`
	RequestDelay    int         `json:"requestDelay"`
}

// OnError represents a communication mechanism for errors.
type OnError struct {
	Email string `json:"email"`
}

// OnFailure represents a communication mechanism for failures.
type OnFailure struct {
	Email string `json:"email"`
}

// Notifications represents a communication structure for notifications.
type Notifications struct {
	OnError   []OnError   `json:"onError"`
	OnFailure []OnFailure `json:"onFailure"`
}

// Schedule represents when the monitor is scheduled to run.
type Schedule struct {
	Cron     string    `json:"cron"`
	Timezone string    `json:"timezone"`
	NextRun  time.Time `json:"nextRun"`
}
