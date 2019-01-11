package go_push_pro

import (
	// "all_push_pro/common"

	"fmt"
	"log"
	"path/filepath"

	"github.com/NaySoftware/go-fcm"
	"github.com/anachronistic/apns"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	// "github.com/sideshow/apns2"
	// "github.com/sideshow/apns2/certificate"
	// "github.com/sideshow/apns2/token"
)

func initiateAndroidPush(device_token string, message string, fcm_key string) {
	// try.This(func() {
	var device_arr []string
	// device_arr = append(device_arr, "dW4chJBem8A:APA91bEDIpHYQEmA9wAuA4uomgiVtaKWA86fkwCogLpseigHEpkmQUv5Ql7BpHcGUbTH5akzKjTv0TLHlKFUeZyzkUAH0DBt9cFxbPh93Xx78QiZ3T1vUOFGTXkgZzDUW8CFnxE2CXzA725r3FYfyCPZJs8wNrfpuw")
	device_arr = append(device_arr, device_token)
	data := map[string]string{"message": message}

	// c := fcm.NewFcmClient("AIzaSyCzIkuC8XZrkK5xVUAUguk23Up2hz_bzmw")
	c := fcm.NewFcmClient(fcm_key)
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

func initiateIosPushWithPem(pem_path string, message string, device_token string, apns_environment string) {
	// try.This(func() {
	// Certificate_path, patherr := filepath.Abs(pem_path)
	Certificate_path, patherr := filepath.Abs(pem_path)
	if patherr != nil {
		fmt.Println(patherr)
	}
	payload := apns.NewPayload()
	payload.Badge = 0
	payload.Sound = "Default"
	payload.Alert = message

	pn := apns.NewPushNotification()
	pn.DeviceToken = device_token
	pn.AddPayload(payload)

	var client *apns.Client
	if apns_environment == "development" || apns_environment == "Development" {
		client = apns.NewClient("gateway.sandbox.push.apple.com:2195", Certificate_path, Certificate_path)
	} else {
		client = apns.NewClient("gateway.push.apple.com:2195", Certificate_path, Certificate_path)
	}
	resp := client.Send(pn)

	alert, _ := pn.PayloadString()
	fmt.Println("Alert:", alert)
	fmt.Println("Success:", resp.Success)
	fmt.Println("Error:", resp.Error)
}

func initiateIosPushWithP8(file_path string, message string, device_token string, keyID string, teamID string, topic string, push_mode string) {

	notification := &apns2.Notification{}
	notification.DeviceToken = device_token
	notification.Topic = topic
	notification.Payload = []byte(`{"aps":{"alert":"` + message + `"}}`)

	authKey, err := token.AuthKeyFromFile(file_path)
	if err != nil {
		log.Fatal("token error:", err)
	}

	token := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID: keyID,
		// TeamID from developer account (View Account -> Membership)
		TeamID: teamID,
	}

	if push_mode == "development" || push_mode == "Development" {
		client := apns2.NewTokenClient(token)
		res, err := client.Push(notification)
	} else {
		client := apns2.NewTokenClient(token).Production()
		res, err := client.Push(notification)
	}

	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

func SendAndroidPush(device_token string, message string, fcm_key string) {
	initiateAndroidPush(device_token, message, fcm_key)
}

func SendIosPushWithPem(pem_path string, message string, device_token string, apns_environment string) {
	initiateIosPushWithPem(pem_path, message, device_token, apns_environment)
}

func SendIosPushWithP8(file_path string, message string, device_token string, key_id string, team_id string, topic string, push_mode string) {
	initiateIosPushWithP8(file_path, message, device_token, key_id, team_id, topic, push_mode)
}

// func main() {
// 	SendIosPush("home/sotsys-229/rails_projects/interstride_app/interstride/config/pushcert/production_12_dec.pem", "HI There", "40ae989435eea69c8b724de427874aaf6f375492cc5b2d46b4db5e1ebeb53400")
// }

// 40ae989435eea69c8b724de427874aaf6f375492cc5b2d46b4db5e1ebeb53400
// home/sotsys-229/rails_projects/interstride_app/interstride/config/pushcert/production_12_dec.pem
