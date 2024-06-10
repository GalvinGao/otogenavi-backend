package postgis

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Structs respresenting varying types of points
type Point struct {
	X, Y float64
}

type PointZ struct {
	X, Y, Z float64
}

type PointM struct {
	X, Y, M float64
}

type PointZM struct {
	X, Y, Z, M float64
}

type PointS struct {
	SRID int32
	X, Y float64
}

type PointZS struct {
	SRID    int32
	X, Y, Z float64
}

type PointMS struct {
	SRID    int32
	X, Y, M float64
}

type PointZMS struct {
	SRID       int32
	X, Y, Z, M float64
}

/** Point functions **/
func (p *Point) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p Point) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p Point) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p Point) GetType() uint32 {
	return 1
}

/** PointZ functions **/
func (p *PointZ) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointZ) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointZ) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointZ) GetType() uint32 {
	return 0x80000001
}

/** PointM functions **/
func (p *PointM) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointM) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointM) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointM) GetType() uint32 {
	return 0x40000001
}

/** PointZM functions **/
func (p *PointZM) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointZM) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointZM) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointZM) GetType() uint32 {
	return 0xC0000001
}

/** PointS functions **/
func (p *PointS) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointS) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointS) MarshalJSON() ([]byte, error) {
	return []byte(`{"type":"Point","coordinates":[` + strconv.FormatFloat(p.X, 'f', -1, 64) + `,` + strconv.FormatFloat(p.Y, 'f', -1, 64) + `]}`), nil
}

func (p *PointS) UnmarshalJSON(data []byte) error {
	// Unmarshal the JSON data into a map
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	// Check if the type is a Point
	if m["type"] != "Point" {
		return errors.New("invalid type")
	}

	// Get the coordinates
	coordinates, ok := m["coordinates"].([]interface{})
	if !ok {
		return errors.New("invalid coordinates")
	}
	p.X = coordinates[0].(float64)
	p.Y = coordinates[1].(float64)

	return nil
}

func (p *PointS) UnmarshalGQL(v interface{}) error {
	data, ok := v.(map[string]interface{})
	if !ok {
		return errors.New("invalid data")
	}

	// Check if the type is a Point
	if data["type"] != "Point" {
		return errors.New("invalid type")
	}

	// Get the coordinates
	coordinates, ok := data["coordinates"].([]interface{})
	if !ok {
		return errors.New("invalid coordinates")
	}
	p.X = coordinates[0].(float64)
	p.Y = coordinates[1].(float64)

	return nil
}

func (p PointS) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, `{"latitude":%f,"longitude":%f}`, p.Y, p.X)
}

func (p PointS) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointS) GetType() uint32 {
	return 0x20000001
}

/** PointZS functions **/
func (p *PointZS) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointZS) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointZS) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointZS) GetType() uint32 {
	return 0xA0000001
}

/** PointMS functions **/
func (p *PointMS) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointMS) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointMS) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointMS) GetType() uint32 {
	return 0x60000001
}

/** PointZMS functions **/
func (p *PointZMS) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointZMS) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointZMS) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointZMS) GetType() uint32 {
	return 0xE0000001
}
