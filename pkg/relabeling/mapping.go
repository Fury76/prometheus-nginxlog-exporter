package relabeling

import (
	"strings"
)

// Map maps a sourceValue from the access log line according to the relabeling
// config (matching against whitelists, regular expressions etc.)
func (r *Relabeling) Map(sourceValue string) (string, error) {
	if r.Split > 0 {
		separator := r.Separator
		if separator == "" {
			separator = " "
		}

		values := strings.Split(sourceValue, separator)

		if len(values) >= r.Split {
			sourceValue = values[r.Split-1]
		} else {
			sourceValue = ""
		}
	}

	if r.WhitelistExists {
		if _, ok := r.WhitelistMap[sourceValue]; ok {
			return sourceValue, nil
		}

		return "other", nil
	}

	if len(r.Matches) > 0 {
		for i := range r.Matches {
			if r.Matches[i].CompiledRegexp.MatchString(sourceValue) {
				sourceValue = r.Matches[i].CompiledRegexp.ReplaceAllString(sourceValue, r.Matches[i].Replacement)
				break
			}
		}
	}

	return sourceValue, nil
}
