// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Marc Berhault (marc@cockroachlabs.com)

package sql

import (
	"fmt"
	"sort"

	"github.com/cockroachdb/cockroach/security"
	"github.com/cockroachdb/cockroach/sql/privilege"
)

// Sort methods for the PrivilegeDescriptor.Users list.
type userPrivilegeList []*UserPrivileges

func (upl userPrivilegeList) Len() int {
	return len(upl)
}

func (upl userPrivilegeList) Swap(i, j int) {
	upl[i], upl[j] = upl[j], upl[i]
}

func (upl userPrivilegeList) Less(i, j int) bool {
	return upl[i].User < upl[j].User
}

// findUserIndex looks for a given user and returns its
// index in the User array if found. Returns -1 otherwise.
func (p *PrivilegeDescriptor) findUserIndex(user string) int {
	idx := sort.Search(len(p.Users), func(i int) bool {
		return p.Users[i].User >= user
	})
	if idx < len(p.Users) && p.Users[idx].User == user {
		return idx
	}
	return -1
}

// findUser looks for a specific user in the list.
// Returns (nil, false) if not found, or (obj, true) if found.
func (p *PrivilegeDescriptor) findUser(user string) (*UserPrivileges, bool) {
	idx := p.findUserIndex(user)
	if idx == -1 {
		return nil, false
	}
	return p.Users[idx], true
}

// findOrCreateUser looks for a specific user in the list, creating it if needed.
func (p *PrivilegeDescriptor) findOrCreateUser(user string) *UserPrivileges {
	idx := sort.Search(len(p.Users), func(i int) bool {
		return p.Users[i].User >= user
	})
	if idx == len(p.Users) {
		// Not found but should be inserted at the end.
		p.Users = append(p.Users, &UserPrivileges{User: user})
	} else if p.Users[idx].User == user {
		// Found.
	} else {
		// New element to be inserted at i.
		p.Users = append(p.Users, nil)
		copy(p.Users[idx+1:], p.Users[idx:])
		p.Users[idx] = &UserPrivileges{User: user}
	}
	return p.Users[idx]
}

// removeUser looks for a given user in the list and removes it if present.
func (p *PrivilegeDescriptor) removeUser(user string) {
	idx := p.findUserIndex(user)
	if idx == -1 {
		// Not found.
		return
	}

	copy(p.Users[idx:], p.Users[idx+1:])
	p.Users[len(p.Users)-1] = nil
	p.Users = p.Users[:len(p.Users)-1]
}

// NewDefaultDatabasePrivilegeDescriptor returns a privilege descriptor
// with ALL privileges for the root user.
func NewDefaultDatabasePrivilegeDescriptor() *PrivilegeDescriptor {
	return &PrivilegeDescriptor{
		Users: []*UserPrivileges{
			{
				User:       security.RootUser,
				Privileges: (1 << privilege.ALL),
			},
		},
	}
}

// Grant adds new privileges to this descriptor for a given list of users.
// TODO(marc): if all privileges other than ALL are set, should we collapse
// them into ALL?
func (p *PrivilegeDescriptor) Grant(user string, privList privilege.List) {
	userPriv := p.findOrCreateUser(user)
	if userPriv.Privileges&(1<<privilege.ALL) != 0 {
		// User already has 'ALL' privilege: no-op.
		return
	}

	bits := privList.ToBitField()
	if bits&(1<<privilege.ALL) != 0 {
		// Granting 'ALL' privilege: overwrite.
		// TODO(marc): the grammar does not allow it, but we should
		// check if other privileges are being specified and error out.
		userPriv.Privileges = (1 << privilege.ALL)
		return
	}
	userPriv.Privileges |= bits
}

// Revoke removes privileges from this descriptor for a given list of users.
func (p *PrivilegeDescriptor) Revoke(user string, privList privilege.List) {
	userPriv, ok := p.findUser(user)
	if !ok || userPriv.Privileges == 0 {
		// Removing privileges from a user without privileges is a no-op.
		return
	}

	bits := privList.ToBitField()
	if bits&(1<<privilege.ALL) != 0 {
		// Revoking 'ALL' privilege: remove user.
		// TODO(marc): the grammar does not allow it, but we should
		// check if other privileges are being specified and error out.
		p.removeUser(user)
		return
	}

	if userPriv.Privileges&(1<<privilege.ALL) != 0 {
		// User has 'ALL' privilege. Remove it and set
		// all other privileges one.
		userPriv.Privileges = 0
		for _, v := range privilege.ByValue {
			if v != privilege.ALL {
				userPriv.Privileges |= (1 << v)
			}
		}
	}

	// One doesn't see "AND NOT" very often.
	userPriv.Privileges &^= bits

	if userPriv.Privileges == 0 {
		p.removeUser(user)
	}
}

// Validate is called when writing a descriptor.
func (p *PrivilegeDescriptor) Validate() error {
	if p == nil {
		return fmt.Errorf("missing privilege descriptor")
	}
	userPriv, ok := p.findUser(security.RootUser)
	if !ok {
		return fmt.Errorf("%s user does not have privileges", security.RootUser)
	}
	if userPriv.Privileges&(1<<privilege.ALL) == 0 {
		return fmt.Errorf("%s user does not have ALL privileges", security.RootUser)
	}
	return nil
}

// UserPrivilegeString is a pair of strings describing the
// privileges for a given user.
type UserPrivilegeString struct {
	User       string
	Privileges string
}

// Show returns the list of {username, privileges} sorted by username.
// 'privileges' is a string of comma-separated sorted privilege names.
func (p *PrivilegeDescriptor) Show() ([]UserPrivilegeString, error) {
	ret := []UserPrivilegeString{}
	for _, userPriv := range p.Users {
		ret = append(ret, UserPrivilegeString{
			User:       userPriv.User,
			Privileges: privilege.ListFromBitField(userPriv.Privileges).SortedString(),
		})
	}
	return ret, nil
}

// CheckPrivilege returns true if 'user' has 'privilege' on this descriptor.
func (p *PrivilegeDescriptor) CheckPrivilege(user string, priv privilege.Kind) bool {
	userPriv, ok := p.findUser(user)
	if !ok {
		return false
	}
	// ALL is always good.
	if userPriv.Privileges&(1<<privilege.ALL) != 0 {
		return true
	}
	return userPriv.Privileges&(1<<priv) != 0
}