package playwright

import (
	"strings"
)

type NameValue struct {
	Name  string
	Value string
}
type HeadersArray []NameValue
type rawHeaders struct {
	headersArray HeadersArray
	headersMap   map[string]map[string]bool
}

func (r *rawHeaders) Get(name string) string {
	values := r.GetAll(name)
	if len(values) == 0 {
		return ""
	}
	sep := ", "
	if strings.ToLower(name) == "set-cookie" {
		sep = "\n"
	}
	return strings.Join(values, sep)
}
func (r *rawHeaders) GetAll(name string) []string {
	name = strings.ToLower(name)
	out := make([]string, 0)
	for value := range r.headersMap[name] {
		out = append(out, value)
	}
	return out
}
func (r *rawHeaders) Headers() map[string]string {
	out := make(map[string]string)
	for key := range r.headersMap {
		out[key] = r.Get(key)
	}
	return out
}

func (r *rawHeaders) HeadersArray() HeadersArray {
	return r.headersArray
}
func newRawHeaders(headers interface{}) *rawHeaders {
	r := &rawHeaders{}
	r.headersArray = make([]NameValue, 0)
	r.headersMap = make(map[string]map[string]bool)
	for _, header := range headers.([]interface{}) {
		entry := header.(map[string]interface{})
		name := entry["name"].(string)
		value := entry["value"].(string)
		r.headersArray = append(r.headersArray, NameValue{
			Name:  name,
			Value: value,
		})
		if _, ok := r.headersMap[strings.ToLower(name)]; !ok {
			r.headersMap[strings.ToLower(name)] = make(map[string]bool)
		}
		r.headersMap[strings.ToLower(name)][value] = true
	}
	return r
}
