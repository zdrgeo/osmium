package command

import "errors"

type builder string

const (
	filepath builder = "filepath"
	pattern  builder = "pattern"
)

func (b *builder) String() string {
	return string(*b)
}

func (b *builder) Set(value string) error {
	switch value {
	case "filepath":
		*b = builder(value)
		return nil
	case "pattern":
		*b = builder(value)
		return nil
	default:
		return errors.New(`must be one of "filepath" or "pattern"`)
	}
}

func (b *builder) Type() string {
	return "builder"
}
