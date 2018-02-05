package structs

import (
	"errors"
	"strings"
)

const (
	errNoLeader         = "No cluster leader"
	errNoRegionPath     = "No path to region"
	errTokenNotFound    = "ACL token not found"
	errPermissionDenied = "Permission denied"
	errNoNodeConn       = "No path to node"
	errUnknownMethod    = "unknown rpc method"
)

var (
	ErrNoLeader         = errors.New(errNoLeader)
	ErrNoRegionPath     = errors.New(errNoRegionPath)
	ErrTokenNotFound    = errors.New(errTokenNotFound)
	ErrPermissionDenied = errors.New(errPermissionDenied)
	ErrNoNodeConn       = errors.New(errNoNodeConn)
	ErrUnknownMethod    = errors.New(errUnknownMethod)
)

// IsErrNoLeader returns whether the error is due to there being no leader.
func IsErrNoLeader(err error) bool {
	return err != nil && strings.Contains(err.Error(), errNoLeader)
}

// IsErrNoRegionPath returns whether the error is due to there being no path to
// the given region.
func IsErrNoRegionPath(err error) bool {
	return err != nil && strings.Contains(err.Error(), errNoRegionPath)
}

// IsErrTokenNotFound returns whether the error is due to the passed token not
// being resolvable.
func IsErrTokenNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), errTokenNotFound)
}

// IsErrPermissionDenied returns whether the error is due to the operation not
// being allowed due to lack of permissions.
func IsErrPermissionDenied(err error) bool {
	return err != nil && strings.Contains(err.Error(), errPermissionDenied)
}

// IsErrNoNodeConn returns whether the error is due to there being no path to
// the given node.
func IsErrNoNodeConn(err error) bool {
	return err != nil && strings.Contains(err.Error(), errNoNodeConn)
}

// IsErrUnknownMethod returns whether the error is due to the operation not
// being allowed due to lack of permissions.
func IsErrUnknownMethod(err error) bool {
	return err != nil && strings.Contains(err.Error(), errUnknownMethod)
}