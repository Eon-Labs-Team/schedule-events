package checkEvents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"schedule-events/packages/common/config"
	"schedule-events/packages/eventSchedule/models"
	"sync"
)

func CheckEvents() {
	fmt.Println("Checking Events")
	events, err := models.LocalGetEventSchedules()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("events", events)
	startSendingProcess(events)
	return
}

func startSendingProcess(events []models.EventSchedule) {
	wg := &sync.WaitGroup{}

	for i := 0; i < len(events); i++ {
		wg.Add(1)
		go sendAndDeleteEvents(events[i], wg)
	}

	wg.Wait()
}

func sendAndDeleteEvents(event models.EventSchedule, wg *sync.WaitGroup) {
	// HTTP endpoint
	url := config.GoDotEnvVariable("TARGET_API_POST")

	fmt.Println("URL:>", url)

	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println("Error with the parsing event")
		wg.Done()
		return
	}

	fmt.Println("Event to send", string(b))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with the request")
		wg.Done()
		return
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Non-OK HTTP status:", resp.StatusCode)
		fmt.Println("Error sending register")
		wg.Done()
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	//delete event

	var deleteResult = models.LocalDeleteEventSchedule(event)
	if !deleteResult.Status {
		fmt.Println("Error deleting register")
		wg.Done()
		return
	}

	fmt.Println("Register deleted correctly - end of process")
	wg.Done()
	return
}
