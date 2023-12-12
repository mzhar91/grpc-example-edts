package repository

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v3"
)

const defaultUUID = "00000000-0000-0000-0000-000000000000"
const defaultErr = "column not specified"
const defaultDateFormat = "2006-01-02 15:04:05"

// Values setup args for insert
func Values(val reflect.Value) ([]string, []interface{}, error) {
	var columns []string
	var values []interface{}
	
	err := errors.New(defaultErr)
	
	for i := 0; i < val.NumField(); i++ {
		isEmpty := true
		column := val.Type().Field(i).Tag.Get("db")
		cse := val.Type().Field(i).Tag.Get("case")
		value := reflect.Indirect(val).FieldByName(val.Type().Field(i).Name).Interface()
		
		if column != "" {
			switch value.(type) {
			case int:
				if value.(int) > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(int) > 0 {
						value = value.(int)
					} else {
						value = nil
					}
				}
				
				break
			case int64:
				if value.(int64) > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(int64) > 0 {
						value = value.(int64)
					} else {
						value = nil
					}
				}
				
				break
			case null.Int:
				if value.(null.Int).ValueOrZero() > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(null.Int).ValueOrZero() > 0 {
						value = value.(null.Int).Int64
					} else {
						value = nil
					}
				}
				
				break
			case bool:
				isEmpty = false
				
				if cse == "nullable" {
					value = nil
				} else {
					value = value.(bool)
				}
				
				break
			case null.Bool:
				isEmpty = false
				
				if cse == "nullable" {
					value = nil
				} else {
					value = value.(null.Bool)
				}
				
				break
			case float64:
				if value.(float64) > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(float64) > 0 {
						value = value.(float64)
					} else {
						value = nil
					}
				}
				
				break
			case null.Float:
				if value.(null.Float).ValueOrZero() > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(null.Float).ValueOrZero() > 0 {
						value = value.(null.Float).Float64
					} else {
						value = nil
					}
				}
				
				break
			case string:
				if len(value.(string)) > 0 || cse == "nullable" {
					isEmpty = false
					
					if len(value.(string)) > 0 {
						value = strings.Replace(value.(string), "'", "''", -1)
					} else {
						value = nil
					}
				}
				
				break
			case null.String:
				if len(value.(null.String).ValueOrZero()) > 0 || cse == "nullable" {
					isEmpty = false
					
					if len(value.(null.String).ValueOrZero()) > 0 {
						value = strings.Replace(value.(null.String).String, "'", "''", -1)
					} else {
						value = nil
					}
				}
				
				break
			case time.Time:
				if !value.(time.Time).IsZero() || cse == "nullable" {
					isEmpty = false
					
					if !value.(time.Time).IsZero() {
						value = value.(time.Time).Format(defaultDateFormat)
					} else {
						value = nil
					}
				}
				
				break
			case null.Time:
				if !value.(null.Time).ValueOrZero().IsZero() || cse == "nullable" {
					isEmpty = false
					
					if !value.(null.Time).ValueOrZero().IsZero() {
						value = value.(null.Time).ValueOrZero().Format(defaultDateFormat)
					} else {
						value = nil
					}
				}
				
				break
			case uuid.UUID:
				if (value.(uuid.UUID).String() != defaultUUID && len(value.(uuid.UUID).String()) > 0) || cse == "nullable" {
					isEmpty = false
					
					if value.(uuid.UUID).String() != defaultUUID && len(value.(uuid.UUID).String()) > 0 {
						value = value.(uuid.UUID).String()
					} else {
						value = nil
					}
				}
				
				break
			case uuid.NullUUID:
				if (value.(uuid.NullUUID).UUID.String() != defaultUUID && len(value.(uuid.NullUUID).UUID.String()) > 0) || cse == "nullable" {
					isEmpty = false
					
					if value.(uuid.NullUUID).UUID.String() != defaultUUID && len(value.(uuid.NullUUID).UUID.String()) > 0 {
						value = value.(uuid.NullUUID).UUID.String()
					} else {
						value = nil
					}
				}
				
				break
			default:
				break
			}
			
			if !isEmpty {
				columns = append(columns, column)
				
				if value == nil {
					values = append(values, "null")
				} else {
					values = append(values, value)
				}
			}
		}
	}
	
	if len(columns) > 0 && len(values) > 0 {
		err = nil
	}
	
	return columns, values, err
}

func BulkValues(kind reflect.Kind, value reflect.Value) ([]string, [][]interface{}, error) {
	var cols, columns []string
	var values [][]interface{}
	
	err := errors.New(defaultErr)
	
	for i := 0; i < value.Len(); i++ {
		var valuesTmp []interface{}
		
		val := value.Index(i).Elem()
		
		for i := 0; i < val.NumField(); i++ {
			isEmpty := true
			column := val.Type().Field(i).Tag.Get("db")
			cse := val.Type().Field(i).Tag.Get("case")
			value := reflect.Indirect(val).FieldByName(val.Type().Field(i).Name).Interface()
			
			if column != "" {
				switch value.(type) {
				case int:
					if value.(int) > 0 || cse == "nullable" {
						isEmpty = false
						
						if value.(int) > 0 {
							value = value.(int)
						} else {
							value = nil
						}
					}
					
					break
				case int64:
					if value.(int64) > 0 || cse == "nullable" {
						isEmpty = false
						
						if value.(int64) > 0 {
							value = value.(int64)
						} else {
							value = nil
						}
					}
					
					break
				case null.Int:
					if value.(null.Int).ValueOrZero() > 0 || cse == "nullable" {
						isEmpty = false
						
						if value.(null.Int).ValueOrZero() > 0 {
							value = value.(null.Int).Int64
						} else {
							value = nil
						}
					}
					
					break
				case bool:
					isEmpty = false
					
					if cse == "nullable" {
						value = nil
					} else {
						value = value.(bool)
					}
					
					break
				case null.Bool:
					isEmpty = false
					
					if cse == "nullable" {
						value = nil
					} else {
						value = value.(null.Bool)
					}
					
					break
				case float64:
					if value.(float64) > 0 || cse == "nullable" {
						isEmpty = false
						
						if value.(float64) > 0 {
							value = value.(float64)
						} else {
							value = nil
						}
					}
					
					break
				case null.Float:
					if value.(null.Float).ValueOrZero() > 0 || cse == "nullable" {
						isEmpty = false
						
						if value.(null.Float).ValueOrZero() > 0 {
							value = value.(null.Float).Float64
						} else {
							value = nil
						}
					}
					
					break
				case string:
					if len(value.(string)) > 0 || cse == "nullable" {
						isEmpty = false
						
						if len(value.(string)) > 0 {
							value = strings.Replace(value.(string), "'", "''", -1)
						} else {
							value = nil
						}
					}
					
					break
				case null.String:
					if len(value.(null.String).ValueOrZero()) > 0 || cse == "nullable" {
						isEmpty = false
						
						if len(value.(null.String).ValueOrZero()) > 0 {
							value = strings.Replace(value.(null.String).String, "'", "''", -1)
						} else {
							value = nil
						}
					}
					
					break
				case time.Time:
					if !value.(time.Time).IsZero() || cse == "nullable" {
						isEmpty = false
						
						if !value.(time.Time).IsZero() {
							value = value.(time.Time).Format(defaultDateFormat)
						} else {
							value = nil
						}
					}
					
					break
				case null.Time:
					if !value.(null.Time).ValueOrZero().IsZero() || cse == "nullable" {
						isEmpty = false
						
						if !value.(null.Time).ValueOrZero().IsZero() {
							value = value.(null.Time).ValueOrZero().Format(defaultDateFormat)
						} else {
							value = nil
						}
					}
					
					break
				case uuid.UUID:
					if (value.(uuid.UUID).String() != defaultUUID && len(value.(uuid.UUID).String()) > 0) || cse == "nullable" {
						isEmpty = false
						
						if value.(uuid.UUID).String() != defaultUUID && len(value.(uuid.UUID).String()) > 0 {
							value = value.(uuid.UUID).String()
						} else {
							value = nil
						}
					}
					
					break
				case uuid.NullUUID:
					if (value.(uuid.NullUUID).UUID.String() != defaultUUID && len(value.(uuid.NullUUID).UUID.String()) > 0) || cse == "nullable" {
						isEmpty = false
						
						if value.(uuid.NullUUID).UUID.String() != defaultUUID && len(value.(uuid.NullUUID).UUID.String()) > 0 {
							value = value.(uuid.NullUUID).UUID.String()
						} else {
							value = nil
						}
					}
					
					break
				default:
					break
				}
				
				if !isEmpty {
					cols = append(cols, column)
					
					if value == nil {
						valuesTmp = append(valuesTmp, "null")
					} else {
						valuesTmp = append(valuesTmp, value)
					}
				}
			}
		}
		
		values = append(values, valuesTmp)
	}
	
	if len(cols) > 0 && len(values) > 0 {
		m := make(map[string]bool)
		
		for _, loop := range cols {
			if _, ok := m[loop]; !ok {
				m[loop] = true
				columns = append(columns, loop)
			}
		}
		
		err = nil
	}
	
	return columns, values, err
}

// Set setup args for update
func Set(val reflect.Value) ([]string, error) {
	var sets []string
	
	err := errors.New(defaultErr)
	
	for i := 0; i < val.NumField(); i++ {
		isEmpty := true
		column := val.Type().Field(i).Tag.Get("db")
		cse := val.Type().Field(i).Tag.Get("case")
		value := reflect.Indirect(val).FieldByName(val.Type().Field(i).Name).Interface()
		
		if column != "" {
			switch value.(type) {
			case int:
				if value.(int) > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(int) > 0 {
						value = value.(int)
					} else {
						value = nil
					}
				}
				
				break
			case int64:
				if value.(int64) > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(int64) > 0 {
						value = value.(int64)
					} else {
						value = nil
					}
				}
				
				break
			case null.Int:
				if value.(null.Int).ValueOrZero() > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(null.Int).ValueOrZero() > 0 {
						value = value.(null.Int).Int64
					} else {
						value = nil
					}
				}
				
				break
			case bool:
				isEmpty = false
				
				if cse == "nullable" {
					value = nil
				} else {
					value = value.(bool)
				}
				
				break
			case null.Bool:
				isEmpty = false
				
				if cse == "nullable" {
					value = nil
				} else {
					value = value.(null.Bool)
				}
				
				break
			case float64:
				if value.(float64) > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(float64) > 0 {
						value = value.(float64)
					} else {
						value = nil
					}
				}
				
				break
			case null.Float:
				if value.(null.Float).ValueOrZero() > 0 || cse == "nullable" {
					isEmpty = false
					
					if value.(null.Float).ValueOrZero() > 0 {
						value = value.(null.Float).Float64
					} else {
						value = nil
					}
				}
				
				break
			case string:
				if len(value.(string)) > 0 || cse == "nullable" {
					isEmpty = false
					
					if len(value.(string)) > 0 {
						value = strings.Replace(value.(string), "'", "''", -1)
					} else {
						value = nil
					}
				}
				
				break
			case null.String:
				if len(value.(null.String).ValueOrZero()) > 0 || cse == "nullable" {
					isEmpty = false
					
					if len(value.(null.String).ValueOrZero()) > 0 {
						value = strings.Replace(value.(null.String).String, "'", "''", -1)
					} else {
						value = nil
					}
				}
				
				break
			case time.Time:
				if !value.(time.Time).IsZero() || cse == "nullable" {
					isEmpty = false
					
					if !value.(time.Time).IsZero() {
						value = value.(time.Time).Format(defaultDateFormat)
					} else {
						value = nil
					}
				}
				
				break
			case null.Time:
				if !value.(null.Time).ValueOrZero().IsZero() || cse == "nullable" {
					isEmpty = false
					
					if !value.(null.Time).ValueOrZero().IsZero() {
						value = value.(null.Time).ValueOrZero().Format(defaultDateFormat)
					} else {
						value = nil
					}
				}
				
				break
			case uuid.UUID:
				if (value.(uuid.UUID).String() != defaultUUID && len(value.(uuid.UUID).String()) > 0) || cse == "nullable" {
					isEmpty = false
					
					if value.(uuid.UUID).String() != defaultUUID && len(value.(uuid.UUID).String()) > 0 {
						value = value.(uuid.UUID).String()
					} else {
						value = nil
					}
				}
				
				break
			case uuid.NullUUID:
				if (value.(uuid.NullUUID).UUID.String() != defaultUUID && len(value.(uuid.NullUUID).UUID.String()) > 0) || cse == "nullable" {
					isEmpty = false
					
					if value.(uuid.NullUUID).UUID.String() != defaultUUID && len(value.(uuid.NullUUID).UUID.String()) > 0 {
						value = value.(uuid.NullUUID).UUID.String()
					} else {
						value = nil
					}
				}
				
				break
			default:
				break
			}
			
			if !isEmpty {
				if value == nil {
					sets = append(sets, fmt.Sprintf("%v = null", column))
				} else {
					sets = append(sets, fmt.Sprintf("%v = '%v'", column, value))
				}
			}
		}
	}
	
	if len(sets) > 0 {
		err = nil
	}
	
	return sets, err
}
