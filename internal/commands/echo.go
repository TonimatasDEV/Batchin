package commands

var Echo = true

func SetEcho(args string) {
	Echo = args == "on"
}

func PrintEcho(args string) {
	println(args)
}
