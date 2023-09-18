package errr

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
)

// register all constraint here
const (
	// All error we know for key constrain
	// uniqueEmail             = "users_email_key"
	// uniqueKTP               = "unique_leads_ktpnumber_indentifier"
	// uniquePhone             = "unique_leads_phone_indentifier"
	allPKey = "pkey" // more general
)

// translateError do transform error to already know error message for better readability
func translateError(errMsg string) string {
	if strings.Contains(errMsg, allPKey) {
		return "duplicate id:: "
	}
	return ""
}

type GormErr struct {
	Code           string `json:"Code"`
	Message        string `json:"Message"`
	ConstraintName string `json:"ConstraintName"`
}

// ParseDBError mapping error to self definition error,
// use it for every error from gorm
func ParseDBError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNoRows
	}

	// marshaling - unmarshal gorm error
	var gormErr GormErr
	byteError, errMarshal := json.Marshal(err)
	if errMarshal == nil {
		_ = json.Unmarshal((byteError), &gormErr)
	}

	if gormErr.Code != "" && gormErr.Message != "" {
		switch gormErr.Code {
		case pgerrcode.ForeignKeyViolation,
			pgerrcode.InvalidTextRepresentation,
			pgerrcode.RestrictViolation,
			pgerrcode.ExclusionViolation,
			pgerrcode.NotNullViolation,
			pgerrcode.CheckViolation,
			pgerrcode.IntegrityConstraintViolation,
			pgerrcode.CardinalityViolation:
			return WrapF(400, "%s%s:: %v", translateError(err.Error()), "bad integrity", err)
		case pgerrcode.UniqueViolation:
			return WrapF(409, "%s%s:: %v", translateError(err.Error()), "bad integrity", err)
		case pgerrcode.UndefinedColumn:
			return WrapF(500, "%s%s:: %v", translateError(err.Error()), "bad implementation", err)
		}
	}

	return err
}
