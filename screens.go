package main

import "github.com/racingmars/go3270"

var loginScreen = go3270.Screen {
	{Row: 0, Col: 35, Intense: true, Content: "Route/3270", Color: go3270.Yellow},
	{Row: 6, Col: 3, Content: "Username: "},
	{Row: 8, Col: 3, Content: "Password: "},
	{Row: 6, Col: 15, Name: "username", Write: true, Highlighting: go3270.Underscore},
	{Row: 6, Col: 30},
	{Row: 8, Col: 15, Name: "password", Write: true, Highlighting: go3270.Underscore, Hidden: true},
	{Row: 8, Col: 30},
	{Row: 12, Col: 3, Content: "Press "},
	{Row: 12, Col: 10, Content: "ENTER", Intense: true},
	{Row: 12, Col: 17, Content: "to log in."},
	{Row: 10, Col: 3, Intense: true, Color: go3270.Red, Name: "errormsg"},
	{Row: 22, Col: 0, Content: "PF3 Exit", Color: go3270.Blue},
}

var loginScreenRules = go3270.Rules{
	"username": {Validator: go3270.NonBlank},
	"password": {Validator: go3270.NonBlank, Reset: true},
}

var connectionErrorScreen = go3270.Screen{
	{Row: 0, Col: 35, Intense: true, Content: "Route/3270", Color: go3270.Yellow},
	{Row: 4, Col: 3, Intense: true, Content: "Error", Color: go3270.Red},
	{Row: 6, Col: 3, Name: "errormsg"},
	{Row: 6, Col: 79},
	{Row: 22, Col: 0, Content: "ENTER Back", Color: go3270.Blue},
	{Row: 20, Col: 1, Content: "Hostname: "},
	{Row: 20, Col: 40, Content: "Port"},
	{Row: 20, Col: 11, Name: "destination", Intense: true},
	{Row: 20, Col: 39},
	{Row: 20, Col: 46, Name: "port", Intense: true},
	{Row: 20, Col: 79},
}

var serverSelectionScreen = go3270.Screen{
	{Row: 0, Col: 35, Intense: true, Content: "Route/3270", Color: go3270.Yellow},
	{Row: 3, Col: 3, Content: "Service       Description", Intense: true},
	{Row: 4, Col: 3, Name: "svc1"},
	{Row: 4, Col: 17, Name: "desc1"},
	{Row: 5, Col: 3, Name: "svc2"},
	{Row: 5, Col: 17, Name: "desc2"},
	{Row: 6, Col: 3, Name: "svc3"},
	{Row: 6, Col: 17, Name: "desc3"},
	{Row: 7, Col: 3, Name: "svc4"},
	{Row: 7, Col: 17, Name: "desc4"},
	{Row: 8, Col: 3, Name: "svc5"},
	{Row: 8, Col: 17, Name: "desc5"},
	{Row: 9, Col: 3, Name: "svc6"},
	{Row: 9, Col: 17, Name: "desc6"},
	{Row: 10, Col: 3, Name: "svc7"},
	{Row: 10, Col: 17, Name: "desc7"},
	{Row: 11, Col: 3, Name: "svc8"},
	{Row: 11, Col: 17, Name: "desc8"},
	{Row: 12, Col: 3, Name: "svc9"},
	{Row: 12, Col: 17, Name: "desc9"},
	{Row: 13, Col: 3, Name: "svc10"},
	{Row: 13, Col: 17, Name: "desc10"},
	{Row: 14, Col: 3, Name: "svc11"},
	{Row: 14, Col: 17, Name: "desc11"},
	{Row: 15, Col: 3, Name: "svc12"},
	{Row: 15, Col: 17, Name: "desc12"},
	{Row: 16, Col: 3, Name: "svc13"},
	{Row: 16, Col: 17, Name: "desc13"},
	{Row: 17, Col: 3, Name: "svc14"},
	{Row: 17, Col: 17, Name: "desc14"},
	{Row: 19, Col: 3, Content: "Please select the service you want to connect to."},
	{Row: 20, Col: 3, Content: "Service > "},
	{Row: 20, Col: 14, Name: "service", Write: true, Highlighting: go3270.Underscore},
	{Row: 20, Col: 21},
	{Row: 21, Col: 3, Intense: true, Color: go3270.Red, Name: "errormsg"},
	{Row: 22, Col: 0, Content: "PF3 Exit          PF12 Logout", Color: go3270.Blue},
}

var serverSelectionScreenRules = go3270.Rules{
	"service": {Validator: go3270.NonBlank},
}