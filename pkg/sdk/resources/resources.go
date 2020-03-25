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

// Formatter represents a resource that provides print formatting information.
type Formatter interface {
	// Format returns column headers and values for the resource.
	Format() ([]string, []interface{})
}

// ResourceType represents the resource type.
type ResourceType int

// Resource types.
const (
	CollectionType ResourceType = iota
	EnvironmentType
	MockType
	MonitorType
	APIType
	APIVersionType
	APIRelationsType
	SchemaType
	WorkspaceType
	UserType
)

// String returns a string version of the ResourceType.
func (r ResourceType) String() string {
	switch r {
	case CollectionType:
		return "Collection"
	case EnvironmentType:
		return "Environment"
	case MockType:
		return "Mock"
	case MonitorType:
		return "Monitor"
	case APIType:
		return "API"
	case APIVersionType:
		return "APIVersion"
	case APIRelationsType:
		return "APIRelations"
	case SchemaType:
		return "Schema"
	case WorkspaceType:
		return "Workspace"
	case UserType:
		return "User"
	}

	return ""
}
