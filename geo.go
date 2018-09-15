package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	PointFmtError = errors.New("Invalid POINT format")
)

type Point struct {
	Lng float64
	Lat float64
}

func (p Point) Value() (driver.Value, error) {
	loc := fmt.Sprintf("st_geographyfromtext('POINT(%f %f)')", p.Lng, p.Lat)
	// TODO If the returned string is quoted, this approach wont work. test it
	return loc, nil
}

func (p *Point) Scan(src interface{}) error {
	val := src.(string)
	if !strings.HasPrefix(val, "POINT(") || !strings.HasSuffix(val, ")") {
		return PointFmtError
	}
	parts := strings.Split(val[6: len(val) - 1], ",")
	if len(parts) != 0 {
		return PointFmtError
	}
	lng, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return PointFmtError
	}
	lat, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return PointFmtError
	}
	p.Lng = lng
	p.Lat = lat
	return nil
}