package gorbacx

import (
	"sync"

	"github.com/deckarep/golang-set"
)

type Role struct {
	mutex             sync.Mutex
	ID                string
	unitPermissions   map[string]*Permission
	personPermissions map[string]*Permission
}

func NewRole(id string, ups, pps []*Permission) *Role {

	upMap := make(map[string]*Permission, len(ups))
	for _, p := range ups {
		upMap[p.ID] = p
	}

	ppMap := make(map[string]*Permission, len(pps))
	for _, p := range pps {
		ppMap[p.ID] = p
	}

	return &Role{ID: id,
		unitPermissions: upMap, personPermissions: ppMap}
}

func (r *Role) Add(ps []*Permission, isUnit bool) {
	if isUnit {
		for _, p := range ps {
			r.unitPermissions[p.ID] = p
		}
	} else {
		for _, p := range ps {
			r.personPermissions[p.ID] = p
		}
	}
}

func (r *Role) Remove(ps []*Permission, isUnit bool) {
	if isUnit {
		for _, p := range ps {
			delete(r.unitPermissions, p.ID)
		}
	} else {
		for _, p := range ps {
			delete(r.personPermissions, p.ID)
		}
	}
}

func (r *Role) Replace(ps []*Permission, isUnit bool) {
	if isUnit {
		r.unitPermissions = make(map[string]*Permission, 0)
	} else {
		r.personPermissions = make(map[string]*Permission, 0)
	}
	r.Add(ps, isUnit)
}

func (r *Role) matchedTypes(isUnit bool) mapset.Set {
	set := mapset.NewSet()
	if isUnit {
		for _, v := range r.unitPermissions {
			set = set.Union(v.types)
		}
	} else {
		for _, v := range r.personPermissions {
			set = set.Union(v.types)
		}
	}
	return set
}
