package database

import "github.com/casbin/casbin"

// CasbinDataSource is a interface
type CasbinDataSource interface {
	ClearPolicy()
	LoadPolicy() error
	SavePolicy() error
	Enforce(rvals ...interface{}) bool
	GetAllSubjects() []string
	GetAllNamedSubjects(ptype string) []string
	GetAllObjects() []string
	GetAllNamedObjects(ptype string) []string
	GetAllActions() []string
	GetAllNamedActions(ptype string) []string
	GetAllRoles() []string
	GetAllNamedRoles(ptype string) []string
	GetPolicy() [][]string
	GetFilteredPolicy(fieldIndex int, fieldValues ...string) [][]string
	GetNamedPolicy(ptype string) [][]string
	GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) [][]string
	GetGroupingPolicy() [][]string
	GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) [][]string
	GetNamedGroupingPolicy(ptype string) [][]string
	GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) [][]string
	HasPolicy(params ...interface{}) bool
	HasNamedPolicy(ptype string, params ...interface{}) bool
	AddPolicy(params ...interface{}) bool
	AddNamedPolicy(ptype string, params ...interface{}) bool
	RemovePolicy(params ...interface{}) bool
	RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) bool
	RemoveNamedPolicy(ptype string, params ...interface{}) bool
	RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) bool
	HasGroupingPolicy(params ...interface{}) bool
	HasNamedGroupingPolicy(ptype string, params ...interface{}) bool
	AddGroupingPolicy(params ...interface{}) bool
	AddNamedGroupingPolicy(ptype string, params ...interface{}) bool
	RemoveGroupingPolicy(params ...interface{}) bool
	RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) bool
	RemoveNamedGroupingPolicy(ptype string, params ...interface{}) bool
	RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) bool
	AddFunction(name string, function func(args ...interface{}) (interface{}, error))
}

type casbinDataSource struct {
	Enforcer *casbin.Enforcer
}

// NewCasbinDataSource is a instance
func NewCasbinDataSource() CasbinDataSource {
	return &casbinDataSource{}
}

func (c *casbinDataSource) ClearPolicy() {
	c.Enforcer.ClearPolicy()
}

func (c *casbinDataSource) LoadPolicy() error {
	return c.Enforcer.LoadPolicy()
}

func (c *casbinDataSource) SavePolicy() error {
	return c.Enforcer.SavePolicy()
}

func (c *casbinDataSource) Enforce(rvals ...interface{}) bool {
	return c.Enforcer.Enforce(rvals...)
}

func (c *casbinDataSource) GetAllSubjects() []string {
	return c.Enforcer.GetAllSubjects()
}

func (c *casbinDataSource) GetAllNamedSubjects(ptype string) []string {
	return c.Enforcer.GetAllNamedSubjects(ptype)
}

func (c *casbinDataSource) GetAllObjects() []string {
	return c.Enforcer.GetAllObjects()
}

func (c *casbinDataSource) GetAllNamedObjects(ptype string) []string {
	return c.Enforcer.GetAllNamedObjects(ptype)
}

func (c *casbinDataSource) GetAllActions() []string {
	return c.Enforcer.GetAllActions()
}

func (c *casbinDataSource) GetAllNamedActions(ptype string) []string {
	return c.Enforcer.GetAllNamedActions(ptype)
}

func (c *casbinDataSource) GetAllRoles() []string {
	return c.Enforcer.GetAllRoles()
}

func (c *casbinDataSource) GetAllNamedRoles(ptype string) []string {
	return c.Enforcer.GetAllNamedRoles(ptype)
}

func (c *casbinDataSource) GetPolicy() [][]string {
	return c.Enforcer.GetPolicy()
}

func (c *casbinDataSource) GetFilteredPolicy(fieldIndex int, fieldValues ...string) [][]string {
	return c.Enforcer.GetFilteredPolicy(fieldIndex, fieldValues...)
}

func (c *casbinDataSource) GetNamedPolicy(ptype string) [][]string {
	return c.Enforcer.GetNamedPolicy(ptype)
}

func (c *casbinDataSource) GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) [][]string {
	return c.Enforcer.GetFilteredNamedPolicy(ptype, fieldIndex, fieldValues...)
}

func (c *casbinDataSource) GetGroupingPolicy() [][]string {
	return c.Enforcer.GetGroupingPolicy()
}

func (c *casbinDataSource) GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) [][]string {
	return c.Enforcer.GetFilteredGroupingPolicy(fieldIndex, fieldValues...)
}

func (c *casbinDataSource) GetNamedGroupingPolicy(ptype string) [][]string {
	return c.Enforcer.GetNamedGroupingPolicy(ptype)
}

func (c *casbinDataSource) GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) [][]string {
	return c.Enforcer.GetFilteredNamedPolicy(ptype, fieldIndex, fieldValues...)
}

func (c *casbinDataSource) HasPolicy(params ...interface{}) bool {
	return c.Enforcer.HasPolicy(params...)
}

func (c *casbinDataSource) HasNamedPolicy(ptype string, params ...interface{}) bool {
	return c.Enforcer.HasNamedPolicy(ptype, params)
}

func (c *casbinDataSource) AddPolicy(params ...interface{}) bool {
	return c.Enforcer.AddPolicy(params...)
}

func (c *casbinDataSource) AddNamedPolicy(ptype string, params ...interface{}) bool {
	return c.Enforcer.AddNamedPolicy(ptype, params...)
}

func (c *casbinDataSource) RemovePolicy(params ...interface{}) bool {
	return c.Enforcer.RemovePolicy(params...)
}

func (c *casbinDataSource) RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) bool {
	return c.Enforcer.RemoveFilteredPolicy(fieldIndex, fieldValues...)
}

func (c *casbinDataSource) RemoveNamedPolicy(ptype string, params ...interface{}) bool {
	return c.Enforcer.RemoveNamedPolicy(ptype, params...)
}

func (c *casbinDataSource) RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) bool {
	return c.Enforcer.RemoveFilteredNamedPolicy(ptype, fieldIndex, fieldValues...)
}

func (c *casbinDataSource) HasGroupingPolicy(params ...interface{}) bool {
	return c.Enforcer.HasGroupingPolicy(params...)
}

func (c *casbinDataSource) HasNamedGroupingPolicy(ptype string, params ...interface{}) bool {
	return c.Enforcer.HasGroupingPolicy(params...)
}

func (c *casbinDataSource) AddGroupingPolicy(params ...interface{}) bool {
	return c.Enforcer.AddGroupingPolicy(params...)
}

func (c *casbinDataSource) AddNamedGroupingPolicy(ptype string, params ...interface{}) bool {
	return c.Enforcer.AddNamedGroupingPolicy(ptype, params...)
}

func (c *casbinDataSource) RemoveGroupingPolicy(params ...interface{}) bool {
	return c.Enforcer.RemoveGroupingPolicy(params...)
}

func (c *casbinDataSource) RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) bool {
	return c.Enforcer.RemoveFilteredGroupingPolicy(fieldIndex, fieldValues...)
}

func (c *casbinDataSource) RemoveNamedGroupingPolicy(ptype string, params ...interface{}) bool {
	return c.Enforcer.RemoveNamedGroupingPolicy(ptype, params...)
}

func (c *casbinDataSource) RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) bool {
	return c.Enforcer.RemoveFilteredNamedGroupingPolicy(ptype, fieldIndex, fieldValues...)
}

func (c *casbinDataSource) AddFunction(name string, function func(args ...interface{}) (interface{}, error)) {
	c.Enforcer.AddFunction(name, function)
}
