package utils

import "golang.org/x/xerrors"

var (
	ErrAuthBodyInvalid       = xerrors.New("ERROR.AUTH.ROLE.INVALID_BODY")
	ErrAuthActionTypeInvalid = xerrors.New("ERROR.AUTH.ROLE.INVALID_ACTION_TYPE")
)
