package errr

import (
	"errors"
	"strings"
)

var (
	// Database error
	ErrNoRows       = errors.New("data not found")
	ErrDBSortFilter = errors.New("invalid filter or sort value")
	ErrDBConn       = errors.New("ErrDBConn")

	// other error
	ErrBindRequest     = errors.New("failed to bind request body")
	ErrValidateRequest = errors.New("request body not valid")
	ErrCopier          = errors.New("failed to copy data from request payload")
	ErrUnauthorized    = errors.New("unauthorized request")
)

// GetLastNErrorMessage return last n wrapper error message,
// adjust deep which is safe to consume for the user
// or set to 0 to bring all error message.
// for example GetLastNErrorMessage(
// "insert lead:: bad integrity:: ERROR #23505 duplicate key value violates unique constraint \"leads_locations_pkey\"",
// 2).
// will return: "error insert lead: bad integrity"
func GetLastNErrorMessage(err error, deep int) string {
	deep = deep - 1
	if err != nil {
		errsMsg := strings.Split(err.Error(), "::")
		deepErr := len(errsMsg)
		var msgBuilder strings.Builder
		if deep < deepErr && deep >= 0 { // check if all index available and deep inserted is not 0
			for i := 0; i <= deep; i++ {
				// modify message if not contains error word in first error
				if i == 0 && !strings.Contains(strings.ToLower(errsMsg[i]), "error") {
					msgBuilder.WriteString("error ")
				}
				msgBuilder.WriteString(errsMsg[i])
				msgBuilder.WriteString(":")
			}
		} else {
			return err.Error()
		}
		return strings.TrimRight(msgBuilder.String(), ":")
	}
	return ""
}
