//
// Chassis.
//

package chassis

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	format = "%d.%02d"
)

var (
	re = regexp.MustCompile(`\A(-?)(\d+)\.(\d{2})\z`)
)

// Represents a USD amount.
type Amount struct {
	amount int64 // In cents
}

func NewAmount(s string) (*Amount, error) {
	if s == "0" {
		return &Amount{0}, nil
	}

	groups := re.FindStringSubmatch(s)
	if groups == nil {
		return nil, errors.New("No match")
	}
	negative := groups[1] == "-"
	dollars, err := strconv.Atoi(groups[2])
	if err != nil {
		return nil, err
	}
	cents, err := strconv.Atoi(groups[3])
	if err != nil {
		return nil, err
	}
	amount := int64(dollars)*100 + int64(cents)
	if negative {
		amount = -1 * amount
	}
	return &Amount{amount}, nil
}

func (a Amount) Add(n Amount) *Amount {
	return &Amount{a.amount + n.amount}
}

func (a Amount) Negate() *Amount {
	return &Amount{-1 * a.amount}
}

func (a *Amount) SqlNullable() *sql.NullString {
	if a != nil {
		return &sql.NullString{a.String(), true}
	}
	return &sql.NullString{Valid: false}
}

func (a *Amount) String() string {
	amount := a.amount
	negative := amount < 0
	if negative {
		amount = -1 * amount
	}
	dollars := int(amount / 100)
	if negative {
		dollars = -1 * dollars
	}
	cents := int(amount % 100)

	return fmt.Sprintf(format, dollars, cents)
}
