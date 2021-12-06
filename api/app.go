package main

import "fmt"

func main() {
	var userinfo = get("https://api.opco.dev/userinfo?userid=1")
	fmt.Println(userinfo)

	var userInfoPost = post("https://api.opco.dev/userinfo?userid=1&usersecret=1", "")
	fmt.Println(userInfoPost)

	var progress = get("https://api.opco.dev/progress?op=op012")
	fmt.Println(progress)
}
