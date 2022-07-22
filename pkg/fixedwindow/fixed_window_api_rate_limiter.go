package fixedwindow



type fixedWindow struct {
	UserID     string `json:"userid"`
	Count      int    `json:"count"`
	WindowSize int    `json:"windowsize"`
}

var fixedWindowUsers = []fixedWindow{
	{UserID: "1", Count: 0, WindowSize: 60000},
}

func GetFixedWindowUsers() []fixedWindow {
	return fixedWindowUsers
}

