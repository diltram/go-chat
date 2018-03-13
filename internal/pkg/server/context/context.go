package context

// Context describes all methods which server context should have.
//
// It will allow to store and retrieve attributes of all types which will be
// provided across whole server execution flow.
// For example you via context you can provide access to Chat structure with
// access to all clients and channels created.
type Context interface {
	SetAttribute(name string, attr interface{}) error
	Attribute(name string) (interface{}, error)
}
