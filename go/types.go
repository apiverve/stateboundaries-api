// Package stateboundaries provides a Go client for the State Boundaries API.
//
// For more information, visit: https://apiverve.com/marketplace/stateboundaries?utm_source=go&utm_medium=readme
package stateboundaries

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// ValidationRule defines validation constraints for a parameter.
type ValidationRule struct {
	Type      string
	Required  bool
	Min       *float64
	Max       *float64
	MinLength *int
	MaxLength *int
	Format    string
	Enum      []string
}

// ValidationError represents a parameter validation error.
type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	return "Validation failed: " + strings.Join(e.Errors, "; ")
}

// Helper functions for pointers
func float64Ptr(v float64) *float64 { return &v }
func intPtr(v int) *int             { return &v }

// Format validation patterns
var formatPatterns = map[string]*regexp.Regexp{
	"email":    regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`),
	"url":      regexp.MustCompile(`^https?://.+`),
	"ip":       regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`),
	"date":     regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
	"hexColor": regexp.MustCompile(`^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`),
}

// Request contains the parameters for the State Boundaries API.
//
// Parameters:
//   - state (required): string - The 2 Letter State Code (e.g., CA) [minLength: 2, maxLength: 2]
type Request struct {
	State string `json:"state"` // Required
}

// ToQueryParams converts the request struct to a map of query parameters.
// Only non-zero values are included.
func (r *Request) ToQueryParams() map[string]string {
	params := make(map[string]string)
	if r == nil {
		return params
	}

	v := reflect.ValueOf(*r)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		// Get the json tag for the field name
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		// Handle tags like `json:"name,omitempty"`
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName == "-" {
			continue
		}

		// Skip zero values
		if field.IsZero() {
			continue
		}

		// Convert to string
		params[jsonName] = fmt.Sprintf("%v", field.Interface())
	}

	return params
}

// Validate checks the request parameters against validation rules.
// Returns a ValidationError if validation fails, nil otherwise.
func (r *Request) Validate() error {
	rules := map[string]ValidationRule{
		"state": {Type: "string", Required: true, MinLength: intPtr(2), MaxLength: intPtr(2)},
	}

	if len(rules) == 0 {
		return nil
	}

	var errors []string
	v := reflect.ValueOf(*r)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		jsonName := strings.Split(jsonTag, ",")[0]

		rule, exists := rules[jsonName]
		if !exists {
			continue
		}

		// Check required
		if rule.Required && field.IsZero() {
			errors = append(errors, fmt.Sprintf("Required parameter [%s] is missing", jsonName))
			continue
		}

		if field.IsZero() {
			continue
		}

		// Type-specific validation
		switch rule.Type {
		case "integer", "number":
			var numVal float64
			switch field.Kind() {
			case reflect.Int, reflect.Int64:
				numVal = float64(field.Int())
			case reflect.Float64:
				numVal = field.Float()
			}
			if rule.Min != nil && numVal < *rule.Min {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be at least %v", jsonName, *rule.Min))
			}
			if rule.Max != nil && numVal > *rule.Max {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be at most %v", jsonName, *rule.Max))
			}

		case "string":
			strVal := field.String()
			if rule.MinLength != nil && len(strVal) < *rule.MinLength {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be at least %d characters", jsonName, *rule.MinLength))
			}
			if rule.MaxLength != nil && len(strVal) > *rule.MaxLength {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be at most %d characters", jsonName, *rule.MaxLength))
			}
			if rule.Format != "" {
				if pattern, ok := formatPatterns[rule.Format]; ok {
					if !pattern.MatchString(strVal) {
						errors = append(errors, fmt.Sprintf("Parameter [%s] must be a valid %s", jsonName, rule.Format))
					}
				}
			}
		}

		// Enum validation
		if len(rule.Enum) > 0 {
			strVal := fmt.Sprintf("%v", field.Interface())
			found := false
			for _, enumVal := range rule.Enum {
				if strVal == enumVal {
					found = true
					break
				}
			}
			if !found {
				errors = append(errors, fmt.Sprintf("Parameter [%s] must be one of: %s", jsonName, strings.Join(rule.Enum, ", ")))
			}
		}
	}

	if len(errors) > 0 {
		return &ValidationError{Errors: errors}
	}
	return nil
}

// ResponseData contains the data returned by the State Boundaries API.
type ResponseData struct {
	Type string `json:"type"`
	Features []FeaturesItem `json:"features"`
}

// FeaturesItem represents an item in the Features array.
type FeaturesItem struct {
	Type string `json:"type"`
	Geometry GeometryData `json:"geometry"`
	Properties PropertiesData `json:"properties"`
}

// GeometryData represents the geometry object.
type GeometryData struct {
	Coordinates []CoordinatesItem `json:"coordinates"`
	Type string `json:"type"`
}

// CoordinatesItem represents an item in the Coordinates array.
type CoordinatesItem struct {
	0 []0Item `json:"0"`
}

// 0Item represents an item in the 0 array.
type 0Item struct {
	0 float64 `json:"0"`
	1 float64 `json:"1"`
}

// PropertiesData represents the properties object.
type PropertiesData struct {
	Gid int `json:"gid"`
	Arealand string `json:"arealand"`
	Division int `json:"division"`
	Intptlat float64 `json:"intptlat"`
	Name string `json:"name"`
	Objectid int `json:"objectid"`
	Areawater string `json:"areawater"`
	Intptlon float64 `json:"intptlon"`
	Oid string `json:"oid"`
	Funcstat string `json:"funcstat"`
	Centlon float64 `json:"centlon"`
	Stusab string `json:"stusab"`
	State string `json:"state"`
	Statens string `json:"statens"`
	Centlat float64 `json:"centlat"`
	Basename string `json:"basename"`
	Mtfcc string `json:"mtfcc"`
	Region int `json:"region"`
	Lsadc string `json:"lsadc"`
	Geoid string `json:"geoid"`
	GeoPoint_2d GeoPoint_2dData `json:"geo_point_2d"`
}

// GeoPoint_2dData represents the geo_point_2d object.
type GeoPoint_2dData struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
