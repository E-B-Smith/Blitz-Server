

//----------------------------------------------------------------------------------------
//
//                                                                 BlitzMessage : Types.go
//                                                        The back-end server to BlitzHere
//
//                                                                  E.B. Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package BlitzMessage


import (
    "fmt"
    "math"
    "time"
    "strconv"
    "strings"
    "unicode"
    "database/sql"
    "github.com/lib/pq"
)


//----------------------------------------------------------------------------------------
//                                                                           Numeric Types
//----------------------------------------------------------------------------------------


func Float64Ptr(value interface{}) *float64 {
    switch f := value.(type) {
    case sql.NullFloat64:
        if f.Valid {
            return &f.Float64
        } else {
            return nil
        }
    case *sql.NullFloat64:
        if f == nil {
            return nil
        } else if f.Valid {
            return &f.Float64
        } else {
            return nil
        }
    case float64:
        return &f
    case *float64:
        return f
    case float32:
        var f64 float64 = float64(f)
        return &f64
    case *float32:
        if f == nil {
            return nil
        } else {
            var f64 float64 = float64(*f)
            return &f64
        }
    default:
        panic(fmt.Errorf("Un-handled type %+v.", f))
    }
}


func Float64FromPtr(f *float64) float64 {
    if f == nil { return 0.0 }
    return *f
}


//----------------------------------------------------------------------------------------
//                                                         Time / Timestamp / sql.NullTime
//----------------------------------------------------------------------------------------


func TimestampPtr(inputValue interface{}) *Timestamp {
    var t *time.Time = nil

    switch val := inputValue.(type) {
    case pq.NullTime:
        if val.Valid && val.Time.Unix() != 0 {
            t = &val.Time
        }

    case *pq.NullTime:
        if val.Valid && val.Time.Unix() != 0 {
            t = &val.Time
        }

    case time.Time:
        t = &val

    case *time.Time:
        t = val

    default:
        panic(fmt.Errorf("Unhandled type: %+v.", val))
    }

    if t == nil { return nil }
    var fepoch float64 = float64(t.UnixNano()) / float64(1000000000)
    ts := Timestamp{ Epoch: &fepoch }
    return &ts
}


func (timestamp *Timestamp) Time() time.Time {
    i, f := math.Modf(*timestamp.Epoch)
    var  sec int64 = int64(math.Floor(i))
    var nsec int64 = int64(f * 1000000.0)
    return time.Unix(sec, nsec)
}


func (timestamp *Timestamp) TimePtr() *time.Time {
    t := timestamp.Time()
    return &t
}


func (timestamp *Timestamp) NullTime() pq.NullTime {
    var nt pq.NullTime
    if timestamp != nil && timestamp.Epoch != nil && *timestamp.Epoch > -22135596800 {
        nt.Valid = true
        nt.Time = timestamp.Time()
    }
    return nt
}


func (timestamp *Timestamp) NullTimePtr() *pq.NullTime {
    nt := timestamp.NullTime()
    return &nt
}


func (timespan *Timespan) NullTimeStart() pq.NullTime {
    if timespan == nil || timespan.StartTimestamp == nil { return pq.NullTime{} }
    return timespan.StartTimestamp.NullTime()
}


func (timespan *Timespan) NullTimeStop() pq.NullTime {
    if timespan == nil || timespan.StopTimestamp == nil { return pq.NullTime{} }
    return timespan.StopTimestamp.NullTime()
}


func TimespanFromNullTimes(startDate, stopDate pq.NullTime) *Timespan {
    if !startDate.Valid && !stopDate.Valid { return nil }

    var t Timespan
    if startDate.Valid {
        t.StartTimestamp = TimestampPtr(startDate)
    }
    if stopDate.Valid {
        t.StopTimestamp = TimestampPtr(stopDate)
    }

    return &t
}


//----------------------------------------------------------------------------------------
//                                                                             Coordinates
//----------------------------------------------------------------------------------------


func CoordinatePtr(lat, lng float64) *Coordinate {
    return &Coordinate{ Latitude: &lat, Longitude: &lng }
}


func CoordinatePtrFromStrings(lat, lng *string) *Coordinate {
    if lat == nil || lng == nil { return nil }
    latf, _ := strconv.ParseFloat(*lat, 64)
    lngf, _ := strconv.ParseFloat(*lng, 64)
    return &Coordinate{ Latitude: &latf, Longitude: &lngf }
}


func CoordinatePtrFromString(lngLat *string) *Coordinate {
    if lngLat == nil { return nil }

    splitFunc := func(c rune) bool {
        if unicode.IsNumber(c) || c == '-' || c == '.' { return false }
        return true
    }
    s := strings.FieldsFunc(*lngLat, splitFunc)
    if len(s) < 2 { return nil }

    return CoordinatePtrFromStrings(&s[1], &s[0])
}


func CoordinateRegionFromCoordinates(latitude1, longitude1, latitude2, longitude2 float64) CoordinateRegion {
    slat := math.Abs(latitude1 - latitude2)
    slng := math.Abs(longitude1 - longitude2)
    clat := math.Min(latitude1,  latitude2) + slat / 2.0
    clng := math.Min(longitude1, longitude2) + slng / 2.0
    cr := CoordinateRegion {
        Center: &Coordinate {Latitude: &clat, Longitude: &clng},
        Span:   &Coordinate {Latitude: &slat, Longitude: &slng},
    }
    return cr
}


func (c Coordinate) IsValid() bool {
    if   c.Latitude == nil   ||  c.Longitude == nil { return false }
    if  *c.Latitude == 0     && *c.Longitude == 0   { return false }
    if  *c.Latitude > 90.0   || *c.Latitude  < -90.0 { return false }
    if  *c.Longitude > 180.0 || *c.Longitude < -180.0 { return false }
    return true
}

/*
func (c *Coordinate) String() string {
    if c == nil { return "<nil>" }
    return fmt.Sprintf("(%1.4f, %1.4f)", c.Latitude, c.Longitude)
}
*/

func (cr CoordinateRegion) Contains(c Coordinate) bool {
    if  c.Latitude != nil && c.Longitude != nil &&
        cr.Center.Latitude != nil && cr.Center.Longitude != nil &&
        math.Abs(*c.Latitude - *cr.Center.Latitude) < (*cr.Span.Latitude / 2.0) &&
        math.Abs(*c.Longitude - *cr.Center.Longitude) < (*cr.Span.Longitude / 2.0) {
        return true
    }
    return false
}

