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

// ItemTree represents a folder/request structure in a Collection
type ItemTree struct {
	Root ItemTreeNode
}

// AddBranch adds a branch to a tree.
func (b *ItemTreeNode) AddBranch(br ItemTreeNode) ItemTreeNode {
	if b.Branches == nil {
		b.Branches = &[]ItemTreeNode{br}
	} else {
		branch := append(*b.Branches, br)
		b.Branches = &branch
	}

	return br
}

// AddItem adds a single item to a tree node.
func (b *ItemTreeNode) AddItem(item Item) Item {
	if b.Items == nil {
		b.Items = &[]Item{item}
	} else {
		items := append(*b.Items, item)
		b.Items = &items
	}

	return item
}

// MakeGroup makes the current node an item group.
func (b *ItemTreeNode) MakeGroup(group ItemGroup) ItemGroup {
	b.ItemGroup = &group
	return *b.ItemGroup
}

// NewItemTree creates a new tree for storing items.
func NewItemTree() *ItemTree {
	tree := &ItemTree{}
	return tree
}

// ItemTreeNode represents a group of items or a single item in a tree.
type ItemTreeNode struct {
	*ItemGroup
	Branches *[]ItemTreeNode
	Items    *[]Item
}
