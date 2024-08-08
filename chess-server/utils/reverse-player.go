package utils

func ReversePlayer(color string) string {
	if color == "white" {
		return "black"
	}
	return "white"
}
