package events

import "main/pkg/types"

type ValidatorActive struct {
	Validator *types.Validator
}
