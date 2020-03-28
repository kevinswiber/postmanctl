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
	Root *itemBranch
}

func (b *itemBranch) addBranch(name string) *itemBranch {
	br := itemBranch{
		Name: name,
	}

	branch := append(*b.Branch, &br)

	b.Branch = &branch
	b.Item = nil

	return (*b.Branch)[len(*b.Branch)-1]
}

func (b *itemBranch) addNode() *Item {
	item := Item{}

	if b.Item == nil {
		b.Item = &[]*Item{&item}
	}

	items := append(*b.Item, &item)

	b.Item = &items
	b.Branch = nil

	return (*b.Item)[len(*b.Item)-1]
}

func newItemTree() *ItemTree {
	tree := &ItemTree{
		Root: &itemBranch{},
	}
	return tree
}

type itemBranch struct {
	Name   string
	Branch *[]*itemBranch
	Item   *[]*Item
}
