/*
 * openapi server
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type Okaserver struct {

	Id string `json:"id"`
}

// AssertOkaserverRequired checks if the required fields are not zero-ed
func AssertOkaserverRequired(obj Okaserver) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertOkaserverConstraints checks if the values respects the defined constraints
func AssertOkaserverConstraints(obj Okaserver) error {
	return nil
}
