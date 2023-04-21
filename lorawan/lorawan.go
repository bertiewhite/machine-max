package lorawan

//go:generate mockery --name LoRaWAN --output ../mocks --with-expecter
type LoRaWAN interface {
	RegisterDevEUI(string) (bool, error)
}

