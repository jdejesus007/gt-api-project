package models

import "strconv"

type BaseStatusEnum uint8

const (
	_                      BaseStatusEnum = iota
	BaseStatusEnumActive                  = 1
	BaseStatusEnumDisabled                = 2
	BaseStatusEnumDeleted                 = 3
)

func (s BaseStatusEnum) String() string {
	return strconv.Itoa(int(s))
}

func (s BaseStatusEnum) StringText() string {
	switch s {
	case BaseStatusEnumActive:
		return "active"
	case BaseStatusEnumDisabled:
		return "disabled"
	case BaseStatusEnumDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}
