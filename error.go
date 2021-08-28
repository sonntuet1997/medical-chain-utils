package utils

import "golang.org/x/xerrors"

var (
	ErrAuthBodyInvalid       = xerrors.New("ERROR.AUTH.INVALID_BODY")
	ErrAuthActionTypeInvalid = xerrors.New("ERROR.AUTH.INVALID_ACTION_TYPE")
	ErrAuthParseModelFail    = xerrors.New("ERROR.AUTH.PARSE_MODEL_FAIL")
	ErrAuthTimestampInvalid  = xerrors.New("ERROR.AUTH.INVALID_TIMESTAMP")
)
