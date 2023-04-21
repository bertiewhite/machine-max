package deveuidgenerator

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"machinemax/lorawan"
)

const allowedChars = "ABCDEF0123456789"

type DevEUIGenerator struct {
	lorawan lorawan.LoRaWAN
	out     chan string
}

func NewDeveuiGenerator(lorawan lorawan.LoRaWAN) *DevEUIGenerator {
	out := make(chan string)
	return &DevEUIGenerator{
		lorawan: lorawan,
		out:     out,
	}
}

func (deg *DevEUIGenerator) GetRegisteredIDChan() chan string {
	return deg.out
}

func generateHexString(length int) (string, error) {
	max := big.NewInt(int64(len(allowedChars)))
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b[i] = allowedChars[n.Int64()]
	}
	return string(b), nil
}

func (deg *DevEUIGenerator) GenAndRegisterDevEUI() error {
	for {
		devEUI, err := generateHexString(5)
		if err != nil {
			// reasonably confident this error won't ever be reached so not concerned with handling 'properly'
			return err
		}

		success, err := deg.lorawan.RegisterDevEUI(devEUI)
		if err != nil {
			// For a cli task this seems like reasonable error handling, log the error and retry with a new deveui however,
			// an actual system could justify retries/more traceability on frequency and types of errors being seen here.
			fmt.Println(fmt.Sprintf("encountered error: %s", err.Error()))
			continue
		}

		if !success {
			// if unsuccessful retry recreation of deveui and register
			continue
		}

		deg.out <- devEUI

		return nil
	}
}
