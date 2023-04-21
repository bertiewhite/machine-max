package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"machinemax/deveuidgenerator"
	"machinemax/jobpool"
	"machinemax/lorawan"
)

func main() {
	errHandler := func(err error) {
		fmt.Println(err.Error())
	}

	lorewan := lorawan.NewLorawanClient(http.DefaultClient)

	generator := deveuidgenerator.NewDeveuiGenerator(lorewan)
	devEuis := generator.GetRegisteredIDChan()

	jp := jobpool.NewJobPool(generator.GenAndRegisterDevEUI, errHandler, 10, 100)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case id, ok := <-devEuis:
				if !ok {
					return
				}
				fmt.Println(fmt.Sprintf("Registered: %s", id))
			}
		}
	}()

	jp.Done = make(chan interface{})

	sigintChan := make(chan os.Signal, 1)
	// Could set jp.Done's channel to accept type os.Signal then we could just do Notify(jp.Done, os.interrupt). And
	// handle the exiting of jobs more directly. But I prefer the re-usability and reduced specificity of this method.
	signal.Notify(sigintChan, os.Interrupt)

	go func() {
		for _ = range sigintChan {
			fmt.Println("Captured os.Interrupt")
			close(jp.Done)
		}
	}()

	jp.Start()
	close(devEuis)
	wg.Wait()
}
