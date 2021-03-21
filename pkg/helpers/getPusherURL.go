package helpers

// getPusherURL returns string needed to connect to pusher
func GetPusherURL() string {
	pu := GetEnv("PUSHER_URL", "http://0c4069fe8f4e7e474bef:b4f13a0194c5a12efd8e@api.pusherapp.com/apps/1174681")
	return pu
}
