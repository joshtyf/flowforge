package util

type ContextKey int

const (
	ConcurrencyIdKey ContextKey = iota
	UrlKey
	NextStepKey
	ServiceRequestKey
	StepKey
)

type OrgContextKey struct{}