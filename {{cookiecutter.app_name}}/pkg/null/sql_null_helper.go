package null

import (
	"database/sql"
	"time"
)

func StrToNullString(s string) sql.NullString {
	if s != "" {
		return sql.NullString{
			String: s,
			Valid:  true,
		}
	}
	return sql.NullString{}
}

func SQLNullStrToStrPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func TimeToNullTime(s time.Time) sql.NullTime {
	if !s.IsZero() {
		return sql.NullTime{
			Time:  s,
			Valid: true,
		}
	}
	return sql.NullTime{}
}

func TimePtrToNullTime(s *time.Time) sql.NullTime {
	if s != nil && !s.IsZero() {
		return sql.NullTime{
			Time:  *s,
			Valid: true,
		}
	}
	return sql.NullTime{}
}

func NullTimeToTimePtr(s sql.NullTime) *time.Time {
	if s.Valid {
		return &s.Time
	}
	return nil
}
