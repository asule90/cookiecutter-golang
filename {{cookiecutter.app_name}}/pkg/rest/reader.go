package rest

import (
	"net/url"
	"strconv"
	"strings"
)

// GetString helper returns a string value from the query string, or the provided
// default value if no matching key could be found.
func GetString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}
	return s
}

// GetBool helper returns a bool value from the query string, or the provided
// default value if no matching key could be found.
func GetBool(qs url.Values, key string, defaultValue bool) bool {
	s := qs.Get(key)
	b, err := strconv.ParseBool(s)
	if err != nil {
		return defaultValue
	}
	return b
}

// GetCSV helper reads a string value from the query string and then splits it
// into a slice on the comma character. If no matching key could be found, it returns
// the provided default value.
func GetCSV(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)
	if csv == "" {
		return defaultValue
	}
	return strings.Split(csv, ",")
}

// GetInt helper reads a string value from the query string and converts it to an
// integer before returning. If no matching key could be found it returns the provided
// default value. If the value couldn't be converted to an integer, then we record an
// error message in the provided Validator instance.
func GetInt(qs url.Values, key string, defaultValue int) int {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return i
}

// GetTraceID helper reads a trace_id value from r.Context() injected by chi middleware.RequestID.
/* func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return ""
	}
	return span.SpanContext().TraceID().String()
} */
