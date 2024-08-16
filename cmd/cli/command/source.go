package command

import "errors"

type source string

const (
	git_commit         source = "git:commit"
	gitHub_pullRequest source = "github:pullrequest"
)

func (s *source) String() string {
	return string(*s)
}

func (s *source) Set(value string) error {
	switch value {
	case "git:commit":
		*s = source(value)
		return nil
	case "github:pullrequest":
		*s = source(value)
		return nil
	default:
		return errors.New(`must be one of "git:commit" or "github:pullrequest"`)
	}
}

func (s *source) Type() string {
	return "source"
}
