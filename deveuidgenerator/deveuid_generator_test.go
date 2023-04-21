package deveuidgenerator

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"machinemax/mocks"
)

func TestDevEUIGenerator_GenAndRegisterDevEUI(t *testing.T) {
	t.Run("If Lorawan registers successfully than deveui is sent to the out channel", func(t *testing.T) {
		mockLorawan := new(mocks.LoRaWAN)

		devEuis := make(chan string, 5)

		generator := DevEUIGenerator{
			lorawan: mockLorawan,
			out:     devEuis,
		}

		mockLorawan.EXPECT().RegisterDevEUI(mock.Anything).Once().Return(true, nil)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := generator.GenAndRegisterDevEUI()
			assert.Nil(t, err)

			close(devEuis)
		}()

		count := 0
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case _, ok := <-devEuis:
					if !ok {
						return
					}
					count++
				}
			}
		}()

		wg.Wait()

		assert.Equal(t, 1, count)
	})

	t.Run("If Lorawan fails to register then try to register a different eui", func(t *testing.T) {
		mockLorawan := new(mocks.LoRaWAN)

		devEuis := make(chan string, 5)

		generator := DevEUIGenerator{
			lorawan: mockLorawan,
			out:     devEuis,
		}

		var firstArg, secondArg string
		mockLorawan.EXPECT().RegisterDevEUI(mock.Anything).Run(func(v1 string) {
			firstArg = v1
		}).Once().Return(false, nil)
		mockLorawan.EXPECT().RegisterDevEUI(mock.Anything).Run(func(v2 string) {
			secondArg = v2
		}).Once().Return(true, nil)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := generator.GenAndRegisterDevEUI()
			assert.Nil(t, err)

			close(devEuis)
		}()

		count := 0
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case _, ok := <-devEuis:
					if !ok {
						return
					}
					count++
				}
			}
		}()

		wg.Wait()

		assert.Equal(t, 1, count)
		// assert the 2 args are different. Currently, there's nothing to stop 2 ids to be gened the same out of pure happenstance
		assert.False(t, firstArg == secondArg)
	})

	t.Run("If Lorawan returns an error whilst trying to register then retry with a different deveui", func(t *testing.T) {
		mockLorawan := new(mocks.LoRaWAN)

		devEuis := make(chan string, 5)

		generator := DevEUIGenerator{
			lorawan: mockLorawan,
			out:     devEuis,
		}

		var firstArg, secondArg string
		mockLorawan.EXPECT().RegisterDevEUI(mock.Anything).Run(func(v1 string) {
			firstArg = v1
		}).Once().Return(true, errors.New("some dummy error"))
		mockLorawan.EXPECT().RegisterDevEUI(mock.Anything).Run(func(v2 string) {
			secondArg = v2
		}).Once().Return(true, nil)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := generator.GenAndRegisterDevEUI()
			assert.Nil(t, err)

			close(devEuis)
		}()

		count := 0
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case _, ok := <-devEuis:
					if !ok {
						return
					}
					count++
				}
			}
		}()

		wg.Wait()

		assert.Equal(t, 1, count)
		// assert the 2 args are different. Currently, there's nothing to stop 2 ids to be gened the same out of pure happenstance
		assert.False(t, firstArg == secondArg)
	})
}
