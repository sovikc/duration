package duration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDuration(t *testing.T) {
	d, err := New("12/02/2000", "14/02/2000")

	if err != nil {
		t.Fail()
	}

	interval, err := d.GetDays()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 1, interval, "they should be equal")
}

func TestGetDays_1(t *testing.T) {
	d, err := New("02/06/1983", "22/06/1983")

	if err != nil {
		t.Fail()
	}

	interval, err := d.GetDays()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 19, interval, "they should be equal")
}

func TestGetDays_2(t *testing.T) {
	d, err := New("04/07/1984", "25/12/1984")

	if err != nil {
		t.Fail()
	}

	interval, err := d.GetDays()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 173, interval, "they should be equal")
}
func TestGetDays_3(t *testing.T) {
	d, err := New("03/01/1989", "03/08/1983")

	if err != nil {
		t.Fail()
	}

	interval, err := d.GetDays()
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 1979, interval, "they should be equal")
}

func TestIncorrectDateSeparator(t *testing.T) {
	_, err := New("12-02/2000", "14/02-2000")

	if err == nil {
		t.Fail()
	}
}
func TestIncorrectDateFormat(t *testing.T) {
	_, err := New("12/2/2000", "14/2/2000")

	if err == nil {
		t.Fail()
	}
}
func TestIncorrectDate(t *testing.T) {
	_, err := New("29/02/2000", "29/2/2001")

	if err == nil {
		t.Fail()
	}
}
