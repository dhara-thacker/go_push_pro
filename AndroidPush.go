package go_push_pro

import (
	// "all_push_pro/common"

	"fmt"

	"github.com/NaySoftware/go-fcm"
)

func initiateAndroidPush(message string) {
	// try.This(func() {
	var device_arr []string
	device_arr = append(device_arr, "dW4chJBem8A:APA91bEDIpHYQEmA9wAuA4uomgiVtaKWA86fkwCogLpseigHEpkmQUv5Ql7BpHcGUbTH5akzKjTv0TLHlKFUeZyzkUAH0DBt9cFxbPh93Xx78QiZ3T1vUOFGTXkgZzDUW8CFnxE2CXzA725r3FYfyCPZJs8wNrfpuw")
	data := map[string]string{"message": message}

	c := fcm.NewFcmClient("AIzaSyCzIkuC8XZrkK5xVUAUguk23Up2hz_bzmw")
	c.NewFcmRegIdsMsg(device_arr, data)
	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println("Failed to send message:", err)
	}
	// }).Finally(func() {
	// 	fmt.Println("Send Android Notification API Block.")
	// }).Catch(func(e try.E) {
	// 	fmt.Println(e)
	// })
}
func SendAndroidPush() {
	initiateAndroidPush("Hi There")
}
