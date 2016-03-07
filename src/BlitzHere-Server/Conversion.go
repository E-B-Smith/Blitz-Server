//  Conversion.go  -  Simple type conversion utilities.
//
//  E.B.Smith  -  October, 2015.


package main


import (
    "time"
    "database/sql"
)


func MinInt(a, b int) int {
    if a < b { return a }
    return b
}


func MaxInt(a, b int) int {
    if a > b { return a }
    return b
}


func Int32PtrFromNullInt64(v sql.NullInt64) *int32 {
    v32 := int32(v.Int64)
    return &v32
}


func Int32PtrFromInt32(i int32) *int32 {
    return &i
}


func Float32PtrFromNullFloat64(v sql.NullFloat64) *float32 {
    v32 := float32(v.Float64)
    return &v32
}


func Float32FromFloat32Ptr(value *float32) float32 {
    if value == nil {
        return 0.0
    } else {
        return *value
    }
}


func Float64FromFloat64Ptr(value *float64) float64 {
    if value == nil {
        return 0.0
    } else {
        return *value
    }
}


func Float64PtrFromFloat64(value float64) *float64 {
    return &value
}


func Float64PtrFromNullFloat(value sql.NullFloat64) *float64 {
    if value.Valid {
        return &value.Float64
    } else {
        return nil
    }
}


func StringFromStringPtr(value *string) string {
    if value == nil {
        return ""
    } else {
        return *value;
    }
}


func Int32FromInt32Ptr(value *int32) int32 {
    if value == nil {
        return 0
    } else {
        return *value;
    }
}


func StringPtrFromNullString(s sql.NullString) *string {
    if s.Valid { return &s.String }
    return nil
}


func StringPtrFromString(s string) *string {
    return &s
}


func BoolPtrFromBool(val bool) *bool {
    return &val;
}


func StringFromPtr(s *string) string {
    if s == nil { return ""; }
    return *s
}


func TimeFromRFC3339(s *string) time.Time {
    if s == nil { return time.Time{} }
    t, error := time.Parse(time.RFC3339, *s)
    if error == nil { return t }
    return time.Time{}
}

// type RowScanner interface {
//     Scan(dest ...interface{}) error
// }

