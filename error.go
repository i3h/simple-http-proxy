package main

func check(err error) {
	if err != nil {
		Log.Warn(err)
	}
}

func pin(err error, msg string) {
	if err != nil {
		Log.Warn(msg)
	}
}
