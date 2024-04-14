// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2024 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

package aa

type Dbus struct {
	Rule
	Qualifier
	Access    string
	Bus       string
	Name      string
	Path      string
	Interface string
	Member    string
	PeerName  string
	PeerLabel string
}

func newDbusFromLog(log map[string]string) *Dbus {
	name := ""
	peerName := ""
	if log["mask"] == "bind" {
		name = log["name"]
	} else {
		peerName = log["name"]
	}
	return &Dbus{
		Rule:      newRuleFromLog(log),
		Qualifier: newQualifierFromLog(log),
		Access:    log["mask"],
		Bus:       log["bus"],
		Name:      name,
		Path:      log["path"],
		Interface: log["interface"],
		Member:    log["member"],
		PeerName:  peerName,
		PeerLabel: log["peer_label"],
	}
}

func (r *Dbus) Less(other any) bool {
	o, _ := other.(*Dbus)
	if r.Access != o.Access {
		return r.Access < o.Access
	}
	if r.Bus != o.Bus {
		return r.Bus < o.Bus
	}
	if r.Name != o.Name {
		return r.Name < o.Name
	}
	if r.Path != o.Path {
		return r.Path < o.Path
	}
	if r.Interface != o.Interface {
		return r.Interface < o.Interface
	}
	if r.Member != o.Member {
		return r.Member < o.Member
	}
	if r.PeerName != o.PeerName {
		return r.PeerName < o.PeerName
	}
	if r.PeerLabel != o.PeerLabel {
		return r.PeerLabel < o.PeerLabel
	}
	return r.Qualifier.Less(o.Qualifier)
}

func (r *Dbus) Equals(other any) bool {
	o, _ := other.(*Dbus)
	return r.Access == o.Access && r.Bus == o.Bus && r.Name == o.Name &&
		r.Path == o.Path && r.Interface == o.Interface &&
		r.Member == o.Member && r.PeerName == o.PeerName &&
		r.PeerLabel == o.PeerLabel && r.Qualifier.Equals(o.Qualifier)
}
