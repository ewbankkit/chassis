//
// Chassis.
//

package chassis

import (
	"errors"
	"testing"
)

func TestNew1(t *testing.T) {
	m, err := NewAmount("123.45")
	if err != nil {
		t.Error(err)
	}
	s := m.String()
	if s != "123.45" {
		t.Errorf("Got %s\n", s)
	}
}

func TestNew2(t *testing.T) {
	m, err := NewAmount("12.09")
	if err != nil {
		t.Error(err)
	}
	s := m.String()
	if s != "12.09" {
		t.Errorf("Got %s\n", s)
	}
}

func TestNew3(t *testing.T) {
	m, err := NewAmount("0")
	if err != nil {
		t.Error(err)
	}
	s := m.String()
	if s != "0.00" {
		t.Errorf("Got %s\n", s)
	}
}

func TestNew4(t *testing.T) {
	m, err := NewAmount("00.00")
	if err != nil {
		t.Error(err)
	}
	s := m.String()
	if s != "0.00" {
		t.Errorf("Got %s\n", s)
	}
}

func TestNew5(t *testing.T) {
	m, err := NewAmount("-99.88")
	if err != nil {
		t.Error(err)
	}
	s := m.String()
	if s != "-99.88" {
		t.Errorf("Got %s\n", s)
	}
}

func TestNew6(t *testing.T) {
	_, err := NewAmount("17.6")
	if err == nil {
		t.Error("Allowed one decimal place")
	}
}

func TestNew7(t *testing.T) {
	_, err := NewAmount("17.654")
	if err == nil {
		t.Error("Allowed three decimal places")
	}
}

func TestNegate1(t *testing.T) {
	m1, err := NewAmount("12.09")
	if err != nil {
		t.Error(err)
	}
	m2 := m1.Negate()
	s := m2.String()
	if s != "-12.09" {
		t.Errorf("Got %s\n", s)
	}
}

func TestNegate2(t *testing.T) {
	m1, err := NewAmount("-34.56")
	if err != nil {
		t.Error(err)
	}
	m2 := m1.Negate()
	s := m2.String()
	if s != "34.56" {
		t.Errorf("Got %s\n", s)
	}
}

func TestAdd1(t *testing.T) {
	m1, err := NewAmount("12.34")
	if err != nil {
		t.Error(err)
	}
	m2, err := NewAmount("5.67")
	if err != nil {
		t.Error(err)
	}
	m3 := m1.Add(*m2)
	s := m3.String()
	if s != "18.01" {
		t.Errorf("Got %s\n", s)
	}
}

func TestAdd2(t *testing.T) {
	m1, err := NewAmount("12.34")
	if err != nil {
		t.Error(err)
	}
	m2, err := NewAmount("-5.67")
	if err != nil {
		t.Error(err)
	}
	m3 := m1.Add(*m2)
	s := m3.String()
	if s != "6.67" {
		t.Errorf("Got %s\n", s)
	}
}

func TestPositiveNegativeZero1(t *testing.T) {
	m, err := NewAmount("12.34")
	if err != nil {
		t.Error(err)
	}
	if !m.IsPositive() {
		t.Fail()
	}
	if m.IsNegative() {
		t.Fail()
	}
	if m.IsZero() {
		t.Fail()
	}
}

func TestPositiveNegativeZero2(t *testing.T) {
	m, err := NewAmount("-12.34")
	if err != nil {
		t.Error(err)
	}
	if m.IsPositive() {
		t.Fail()
	}
	if !m.IsNegative() {
		t.Fail()
	}
	if m.IsZero() {
		t.Fail()
	}
}

func TestPositiveNegativeZero3(t *testing.T) {
	m, err := NewAmount("0")
	if err != nil {
		t.Error(err)
	}
	if m.IsPositive() {
		t.Fail()
	}
	if m.IsNegative() {
		t.Fail()
	}
	if !m.IsZero() {
		t.Fail()
	}
}

func TestSqlNullable1(t *testing.T) {
	m, err := NewAmount("12.34")
	if err != nil {
		t.Error(err)
	}
	s := m.SqlNullable()
	if !s.Valid {
		t.Error(errors.New("Expected not NULL"))
	}
}

func TestSqlNullable2(t *testing.T) {
	var m *Amount
	s := m.SqlNullable()
	if s.Valid {
		t.Error(errors.New("Expected NULL"))
	}
}
