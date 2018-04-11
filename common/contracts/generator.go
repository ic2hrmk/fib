package contracts

type INumberGenerator interface {
	Next() uint64
	Current() uint64
}