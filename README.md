# GoPushPro

This is a package for sending iOS and Android Push Notification in GoLang

### Getting Started

Inorder to use this package, import it using: 

```
github.com/dhara-thacker/go_push_pro
```

### Installing

To install the package in your system, run this command in your temrinal:

```
go get github.com/dhara-thacker/go_push_pro
```

### Send Android Push with FCM

```
package main

import "github.com/dhara-thacker/go_push_pro"

func main() {
	go_push_pro.SendAndroidPush(device_token, message, fcm_key)
}
```

### Send iOS Push with Pem File

```
package main

import "github.com/dhara-thacker/go_push_pro"

func main() {
	go_push_pro.SendIosPushWithPem(pem_file_path, message, device_token, apns_environment)
}
```

Here APNS Environment can be any of the 2 below:

1) development
2) production

### Send iOS Push using P8

```
package main

import "github.com/dhara-thacker/go_push_pro"

func main() {
	go_push_pro.SendIosPushWithP8(pem_file_path, message, device_token, key_id, team_id, topic, push_mode)
}
```

Here push Mode can be any of the 2 below:

1) development
2) production
