package common

import (
	"golang.org/x/xerrors"
	"google.golang.org/grpc/status"
)

var (
	ErrAuthBodyInvalid       = xerrors.New("ERROR.AUTH.INVALID_BODY")
	ErrAuthActionTypeInvalid = xerrors.New("ERROR.AUTH.INVALID_ACTION_TYPE")
	ErrAuthParseModelFail    = xerrors.New("ERROR.AUTH.PARSE_MODEL_FAIL")
	ErrAuthTimestampInvalid  = xerrors.New("ERROR.AUTH.INVALID_TIMESTAMP")
	ErrNotLoginYet           = xerrors.New("ERROR.AUTH.NOT_LOGIN_YET")
	ErrMissingCertificate    = xerrors.New("ERROR.AUTH.MISSING_CERTIFICATE")
	ErrParseError            = xerrors.New("ERROR.COMMON.PARSE_ERROR_FAIL")
)

func ParseGrpcError(err error) error {
	s, ok := status.FromError(err)
	if !ok {
		return ErrParseError
	}

	return xerrors.New(s.Message())
}
