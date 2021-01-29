package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"rs/cors"
	"strconv"
	"strings"
	"sync"
	"time"
)

type message struct {
	Command  string
	Sensor   string
	SensorID string
	Update   string
	Value    string
}

type Response struct {
	Status string
}

var jsonMsg message
var mu sync.Mutex
var txBuffer chan string
var stopThread1 chan string
var stopThread2 chan string
var stopThread3 chan string
var stopThread4 chan string

func autoSensorUpdate(sensor string, sensorID string) {
	temp := 25000
	hum := 50
	battery := 85
	threadID, _ := strconv.Atoi(sensorID)
	var msg string
	var zwaveMsg string
	var msgLength int
	for {
		flag := "TRUE"
		msg = fmt.Sprintf("heartbeat ZW:0:%s:1", sensor)
		msgLength = len(msg)
		zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
		mu.Lock()
		txBuffer <- zwaveMsg
		mu.Unlock()
		time.Sleep(250 * time.Millisecond)
		msg = fmt.Sprintf("battery ZW:0:%s:1 %s", sensor, strconv.Itoa(battery))
		msgLength = len(msg)
		zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
		mu.Lock()
		txBuffer <- zwaveMsg
		mu.Unlock()
		time.Sleep(250 * time.Millisecond)
		msg = fmt.Sprintf("temperature ZW:0:%s:1 %s", sensor, strconv.Itoa(temp))
		msgLength = len(msg)
		zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
		mu.Lock()
		txBuffer <- zwaveMsg
		mu.Unlock()
		time.Sleep(250 * time.Millisecond)
		msg = fmt.Sprintf("heartbeat ZW:0:%s:2", sensor)
		msgLength = len(msg)
		zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
		mu.Lock()
		txBuffer <- zwaveMsg
		mu.Unlock()
		time.Sleep(250 * time.Millisecond)
		msg = fmt.Sprintf("battery ZW:0:%s:2 %s", sensor, strconv.Itoa(battery))
		msgLength = len(msg)
		zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
		mu.Lock()
		txBuffer <- zwaveMsg
		mu.Unlock()
		time.Sleep(250 * time.Millisecond)
		msg = fmt.Sprintf("humidity ZW:0:%s:2 %s", sensor, strconv.Itoa(hum))
		msgLength = len(msg)
		zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
		mu.Lock()
		txBuffer <- zwaveMsg
		mu.Unlock()
		for i := 0; i < 4; i++ {

			if threadID == 1 {
				select {
				case cmd := <-stopThread1:
					switch cmd {
					case "stop1":
						fmt.Println("Stopping thread1")
						time.Sleep(300 * time.Millisecond)
						deleteSensor(sensor)
						return
					}
				default:
					if strings.Compare(flag, "TRUE") == 0 {
						temp = temp + 500
						hum = hum + 1
					}
				}
			}

			if threadID == 2 {
				select {
				case cmd := <-stopThread2:
					switch cmd {
					case "stop2":
						fmt.Println("Stopping thread2")
						time.Sleep(300 * time.Millisecond)
						deleteSensor(sensor)
						return
					}
				default:
					if strings.Compare(flag, "TRUE") == 0 {
						temp = temp + 500
						hum = hum + 1
					}
				}
			}

			if threadID == 3 {
				select {
				case cmd := <-stopThread3:
					switch cmd {
					case "stop3":
						fmt.Println("Stopping thread3")
						time.Sleep(300 * time.Millisecond)
						deleteSensor(sensor)
						return
					}
				default:
					if strings.Compare(flag, "TRUE") == 0 {
						temp = temp + 500
						hum = hum + 1
					}
				}
			}

			if threadID == 4 {
				select {
				case cmd := <-stopThread4:
					switch cmd {
					case "stop4":
						fmt.Println("Stopping thread4")
						time.Sleep(300 * time.Millisecond)
						deleteSensor(sensor)
						return
					}
				default:
					if strings.Compare(flag, "TRUE") == 0 {
						temp = temp + 500
						hum = hum + 1
					}
				}
			}
			flag = "FALSE"
			time.Sleep(4 * time.Second)
		}
	}
}

func manualBatUpdate(sensor string, value string) {
	var msg string
	var zwaveMsg string
	var msgLength int
	msg = fmt.Sprintf("battery ZW:0:%s:1 %s", sensor, value)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(250 * time.Millisecond)
	msg = fmt.Sprintf("battery ZW:0:%s:2 %s", sensor, value)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
}

func createSensor(sensor string) {
	var msg string
	var zwaveMsg string
	var msgLength int
	msg = fmt.Sprintf("createTemperatureSensor ZW:0:%s:1 1 ZW:0:%s", sensor, sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(250 * time.Millisecond)
	msg = fmt.Sprintf("description ZW:0:%s:1 %s", sensor, sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(250 * time.Millisecond)
	msg = fmt.Sprintf("presence ZW:0:%s:1 1", sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(250 * time.Millisecond)
	msg = fmt.Sprintf("heartbeat ZW:0:%s:1", sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(250 * time.Millisecond)
	msg = fmt.Sprintf("createHumiditySensor ZW:0:%s:2 1 ZW:0:%s", sensor, sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(250 * time.Millisecond)
	msg = fmt.Sprintf("description ZW:0:%s:2 %s", sensor, sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(250 * time.Millisecond)
	msg = fmt.Sprintf("presence ZW:0:%s:2 1", sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(250 * time.Millisecond)
	msg = fmt.Sprintf("heartbeat ZW:0:%s:2", sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
}

func deleteSensor(sensor string) {
	var msg string
	var zwaveMsg string
	var msgLength int
	msg = fmt.Sprintf("destroy ZW:0:%s:1", sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
	time.Sleep(1 * time.Second)
	msg = fmt.Sprintf("destroy ZW:0:%s:2", sensor)
	msgLength = len(msg)
	zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
	mu.Lock()
	txBuffer <- zwaveMsg
	mu.Unlock()
}

func convertJsontoZwaveMsg(command string, sensor string, sensorID string, update string, value string) string {
	var zwaveMsg string
	var msg string
	var msgLength int

	if strings.Compare(command, "CreateSensor") == 0 {
		createSensor(sensor)
	}

	if strings.Compare(command, "DeleteSensor") == 0 {
		if strings.Compare(update, "manual") == 0 {
			deleteSensor(sensor)
		}
		if strings.Compare(update, "auto") == 0 {
			if strings.Compare(sensorID, "1") == 0 {
				stopThread1 <- "stop1"
			}
			if strings.Compare(sensorID, "2") == 0 {
				stopThread2 <- "stop2"
			}
			if strings.Compare(sensorID, "3") == 0 {
				stopThread3 <- "stop3"
			}
			if strings.Compare(sensorID, "4") == 0 {
				stopThread4 <- "stop4"
			}
		}
	}

	if strings.Compare(command, "TempUpdate") == 0 {
		if strings.Compare(update, "manual") == 0 {
			convertedValue, _ := strconv.Atoi(value)
			milliCelsius := int(((float32(convertedValue) - 32.0) * (5.0 / 9.0)) * 1000)
			msg = fmt.Sprintf("temperature ZW:0:%s:1 %s", sensor, strconv.Itoa(milliCelsius))
			msgLength = len(msg)
			zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
		}
	}

	if strings.Compare(command, "HumUpdate") == 0 {
		if strings.Compare(update, "manual") == 0 {
			msg = fmt.Sprintf("humidity ZW:0:%s:2 %s", sensor, value)
			msgLength = len(msg)
			zwaveMsg = fmt.Sprintf("ZW:\"%s\",size: %d", msg, msgLength+1)
		}
	}

	if strings.Compare(command, "BatUpdate") == 0 {
		if strings.Compare(update, "manual") == 0 {
			manualBatUpdate(sensor, value)
		}
	}

	if strings.Compare(command, "SensorUpdate") == 0 {
		if strings.Compare(update, "auto") == 0 {
			go autoSensorUpdate(sensor, sensorID)
		}
	}
	return zwaveMsg
}

func webRequestHandler(rw http.ResponseWriter, request *http.Request) {
	fmt.Println("Received message from websocket")
	decoder := json.NewDecoder(request.Body)

	var commandMsg string
	var sensorMsg string
	var sensorIDMsg string
	var updateMsg string
	var valueMsg string
	var zwaveMsg string
	err := decoder.Decode(&jsonMsg)
	if err != nil {
		panic(err)
	}
	commandMsg = jsonMsg.Command
	sensorMsg = jsonMsg.Sensor
	sensorIDMsg = jsonMsg.SensorID
	updateMsg = jsonMsg.Update
	valueMsg = jsonMsg.Value
	zwaveMsg = convertJsontoZwaveMsg(commandMsg, sensorMsg, sensorIDMsg, updateMsg, valueMsg)
	if len(zwaveMsg) > 0 {
		mu.Lock()
		txBuffer <- zwaveMsg
		mu.Unlock()
	}
	Res := "Success"
	profile := Response{Res}
	json.NewEncoder(rw).Encode(profile)

}

func zwaveThread() {
	conn, err := net.Dial("tcp", "127.0.0.1:4448")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Fprintf(conn, <-txBuffer+"\n")
		fmt.Println("Message was sent to XBRIDGE")
	}
}

func webHandlerThread() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", webRequestHandler)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}

func main() {
	txBuffer = make(chan string)
	stopThread1 = make(chan string)
	stopThread2 = make(chan string)
	stopThread3 = make(chan string)
	stopThread4 = make(chan string)
	fmt.Println("Starting Sensor Simulator")
	go zwaveThread()
	go webHandlerThread()
	/*	for {
			time.Sleep(1000 * time.Second)
		}
	*/
	select {}
}
