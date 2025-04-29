package common

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time" // Example for handling specific types like time.Time
)

const (
	// tagName is the struct tag used for customization.
	tagName = "url"
	// omitEmptyOption is the tag option to omit zero-value fields.
	omitEmptyOption = "omitempty"
	// timeFormatOption is a potential tag option to specify time format (example).
	timeFormatOption = "format"
	// defaultTimeFormat is the default format if none is specified.
	defaultTimeFormat = time.RFC3339
)

// EncodeURLValues encodes the given struct into url.Values.
// It handles nested structs and slices of structs.
// Use the `url` struct tag for customization:
//   - `url:"custom_name"`: Use "custom_name" as the key.
//   - `url:"-"`: Skip this field.
//   - `url:"custom_name,omitempty"`: Use "custom_name" and omit if the field has its zero value.
//   - `url:"-,omitempty"`: Skip if zero value (though `url:"-"` alone is sufficient to skip).
//   - `url:"start_date,format=2006-01-02"`: Example for custom time formatting.
func EncodeURLValues(data interface{}) (url.Values, error) {
	values := make(url.Values)
	val := reflect.ValueOf(data)

	// If it's a pointer, get the underlying element
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return values, nil // Nothing to encode from a nil pointer
		}
		val = val.Elem()
	}

	// Ensure we are dealing with a struct
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct or pointer to struct, got %s", val.Kind())
	}

	err := encodeRecursive("", val, values)
	if err != nil {
		return nil, err
	}
	return values, nil
}

// encodeRecursive is the helper function that processes struct fields.
// MODIFIED to handle PHP-style array notation like field[index][nested_field]
func encodeRecursive(prefix string, val reflect.Value, values url.Values) error {
	// Dereference pointer if necessary
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil // Skip nil pointers
		}
		val = val.Elem()
	}

	// If the value passed is not a struct anymore (e.g., after dereferencing a pointer),
	// we might need to handle its basic type directly here.
	// However, the primary loop assumes 'val' is a struct. Add check for safety.
	if val.Kind() != reflect.Struct {
		// This case might occur if encoding e.g. a *string field directly?
		// Usually pointers are handled within the loop below.
		// If needed, handle basic types here based on 'prefix'.
		// For now, we assume the initial call and recursive calls for structs land here.
		// Let's skip if it's not a struct at this point in recursion.
		// fmt.Printf("Warning: encodeRecursive called with non-struct type %s for prefix '%s'\n", val.Kind(), prefix)
		return nil // Or handle appropriately if a non-struct is expected
	}

	typ := val.Type()

	// Iterate over struct fields
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		structField := typ.Field(i)

		// Skip unexported fields
		if structField.PkgPath != "" {
			continue
		}

		// --- Tag Parsing ---
		tag := structField.Tag.Get(tagName)
		tagParts := strings.Split(tag, ",")
		fieldName := tagParts[0]
		options := tagParts[1:]

		if fieldName == "-" {
			continue
		}
		if fieldName == "" {
			fieldName = structField.Name // Default to field name
		}

		// --- MODIFIED Key Construction ---
		var currentKey string
		if prefix == "" {
			// Top-level field
			currentKey = fieldName
		} else {
			// Check if the prefix already represents an array element (e.g., "actions[0]")
			if strings.HasSuffix(prefix, "]") {
				// Nested field within an array element: use PHP-style array notation
				// e.g., prefix="actions[0]", fieldName="cmd" -> "actions[0][cmd]"
				currentKey = fmt.Sprintf("%s[%s]", prefix, fieldName)
			} else {
				// Standard nested struct field: use dot notation
				// e.g., prefix="lead", fieldName="first_name" -> "lead.first_name"
				currentKey = prefix + "." + fieldName
			}
		}
		// --- End MODIFIED Key Construction ---

		// Handle omitempty (using the original field value before potential dereferencing)
		omitEmpty := false
		for _, opt := range options {
			if opt == omitEmptyOption {
				omitEmpty = true
				break
			}
		}
		// Use the actual field value (fieldVal) for the omitempty check
		if omitEmpty && isEmptyValue(fieldVal) {
			continue
		}

		// Handle pointers specifically *before* the main kind switch if they point to basic types
		// or if we need to recurse differently for pointer vs non-pointer structs/slices
		if fieldVal.Kind() == reflect.Ptr {
			if fieldVal.IsNil() {
				// If nil and not omitted, decide what to do. Typically skip.
				continue
			}
			// Dereference the pointer for further processing
			fieldVal = fieldVal.Elem() // Now fieldVal holds the pointed-to value
		}

		// --- Value Encoding Based on Kind (using potentially dereferenced fieldVal) ---
		switch fieldVal.Kind() {
		case reflect.Struct:
			if fieldVal.Type() == reflect.TypeOf(time.Time{}) {
				t := fieldVal.Interface().(time.Time)
				// Check IsZero on the actual time value
				if !t.IsZero() {
					format := defaultTimeFormat
					for _, opt := range options {
						if strings.HasPrefix(opt, timeFormatOption+"=") {
							format = strings.TrimPrefix(opt, timeFormatOption+"=")
							break
						}
					}
					values.Add(currentKey, t.Format(format))
				}
				// If zero and omitempty wasn't specified, it adds nothing, effectively omitting.
			} else {
				// Recursively encode nested struct, passing the *newly formatted* key
				err := encodeRecursive(currentKey, fieldVal, values)
				if err != nil {
					return err
				}
			}

		case reflect.Slice, reflect.Array:
			sliceLen := fieldVal.Len()
			for j := 0; j < sliceLen; j++ {
				elem := fieldVal.Index(j) // Get the element Value

				// Construct the base key for this slice element using standard index notation
				// e.g., currentKey="actions", j=0 -> "actions[0]"
				// e.g., currentKey="customer.other_addresses", j=1 -> "customer.other_addresses[1]"
				indexedKeyForElement := fmt.Sprintf("%s[%d]", currentKey, j)

				// Handle pointer elements within the slice
				if elem.Kind() == reflect.Ptr {
					if elem.IsNil() {
						continue // Skip nil elements in slice
					}
					elem = elem.Elem() // Dereference for processing
				}

				// Now check the kind of the (potentially dereferenced) element
				if elem.Kind() == reflect.Struct {
					// Recursively encode the struct element.
					// Pass the indexed key as the prefix. The next recursive call's
					// key construction logic will handle field names correctly (e.g., actions[0][cmd])
					err := encodeRecursive(indexedKeyForElement, elem, values)
					if err != nil {
						return err
					}
				} else {
					// Encode basic type element within the slice
					strValue, err := valueToString(elem, options)
					if err != nil {
						return fmt.Errorf("error converting slice element %s to string: %w", indexedKeyForElement, err)
					}
					// Add basic slice element using the indexed key format.
					// Kartra might expect repeated keys for basic slices (e.g., tags=urgent&tags=fragile)
					// If so, use 'currentKey' instead of 'indexedKeyForElement' here.
					// For consistency with struct slices, we'll use indexed keys: tags[0]=urgent
					values.Add(indexedKeyForElement, strValue)
				}
			}

		case reflect.Map:
			// Map handling - Check if Kartra uses a specific map format
			// Assuming key[mapKey]=mapValue for now, might need adjustment
			iter := fieldVal.MapRange()
			for iter.Next() {
				k := iter.Key()
				v := iter.Value()

				kStr, err := valueToString(k, nil)
				if err != nil {
					return fmt.Errorf("error converting map key for field %s to string: %w", currentKey, err)
				}

				// Format map key access: field[key]
				mapKey := fmt.Sprintf("%s[%s]", currentKey, kStr)

				if v.Kind() == reflect.Ptr {
					if v.IsNil() {
						continue
					}
					v = v.Elem()
				}

				// Simplified map value handling (encode complex types as string)
				// Might need recursion if maps can contain structs/slices for Kartra
				if v.Kind() == reflect.Struct || v.Kind() == reflect.Map || v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
					// Decide how Kartra expects nested types in maps. Using %v for simplicity.
					values.Add(mapKey, fmt.Sprintf("%v", v.Interface()))
				} else {
					strValue, err := valueToString(v, options)
					if err != nil {
						return fmt.Errorf("error converting map value for key %s to string: %w", mapKey, err)
					}
					values.Add(mapKey, strValue)
				}
			}

		default: // Basic types (string, int, bool, float, etc.)
			// Use the (potentially dereferenced) fieldVal
			strValue, err := valueToString(fieldVal, options)
			if err != nil {
				return fmt.Errorf("error converting field %s to string: %w", currentKey, err)
			}
			values.Add(currentKey, strValue)
		}
	}
	return nil
}

// IMPORTANT: Keep the rest of the code (EncodeURLValues, isEmptyValue, valueToString, struct definitions, main example)
// from the previous answer, just replace the old encodeRecursive with this modified one.

// isEmptyValue checks if a reflect.Value is its zero value.
// It's like `reflect.Value.IsZero()` but also considers nil pointers/interfaces/maps/slices.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		// Special check for time.Time zero value
		if v.Type() == reflect.TypeOf(time.Time{}) {
			return v.Interface().(time.Time).IsZero()
		}
		// General struct check - consider it non-empty if it exists
		// A more complex check could see if *all* fields are zero, but IsZero() is usually sufficient
		return v.IsZero() // Use built-in IsZero for general structs
	}
	return false // Default to non-empty for unknown kinds
}

// valueToString converts a reflect.Value to its string representation.
// Handles basic types. `options` can be used for formatting hints (e.g., time).
func valueToString(v reflect.Value, options []string) (string, error) {
	// Dereference pointer if necessary
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "", nil // Represent nil pointer as empty string
		}
		v = v.Elem()
	}

	// Handle time.Time specifically for formatting
	if v.Type() == reflect.TypeOf(time.Time{}) {
		t := v.Interface().(time.Time)
		format := defaultTimeFormat
		for _, opt := range options {
			if strings.HasPrefix(opt, timeFormatOption+"=") {
				format = strings.TrimPrefix(opt, timeFormatOption+"=")
				break
			}
		}
		return t.Format(format), nil
	}

	// Use fmt.Sprintf for general conversion of basic types
	// It handles most primitive types reasonably well.
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32), nil
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	default:
		// Fallback for other potential simple types or types that have a String() method
		if v.IsValid() && v.CanInterface() {
			return fmt.Sprintf("%v", v.Interface()), nil
		}
		return "", fmt.Errorf("unsupported type for string conversion: %s", v.Kind())
	}
}
