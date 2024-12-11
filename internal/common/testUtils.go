package common

import "github.com/stretchr/testify/mock"

func HandleMockCall[T any](args mock.Arguments) (T, error) {
	var zero T
	if args.Error(1) != nil {
		return zero, args.Error(1)
	}
	return args.Get(0).(T), nil
}