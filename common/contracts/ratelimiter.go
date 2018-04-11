package contracts

type IRateLimiter interface {
	IsAllowed() bool
	AddRequest()
}