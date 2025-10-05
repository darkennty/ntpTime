package ntpTime

import (
	"errors"
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func GetTime() (string, error) {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error at getting time: %v\n", err))
	}

	return t.Format(time.RFC3339), nil
}
