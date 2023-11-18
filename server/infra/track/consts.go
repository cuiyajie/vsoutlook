package track

// event keys
const (
	EventType = "type"
	UserID    = "user"
	At        = "at"
	Site      = "site"

	// api
	Path      = "path"
	FileOwner = "file_owner"
	Issue     = "issue"
	FileID    = "file"
	ReqParams = "reqParams"

	IsInternalUser = "internal"
	IsEarlyAccess  = "eaccess"
	RegisteredDays = "rdays"

	Browser        = "browser"
	BrowserVersion = "bsv"

	BrowserID    = "bsid"
	SessionID    = "ssid"
	SessionAge   = "ssage"
	RuntimeAge   = "rtage"
	WebSocketAge = "wsage"

	TabsCount      = "tabsc"
	CurrentPage    = "cpage"
	CurrentProject = "project"
	CurrentDesign  = "vthread"
)

// helper keys:
// not in events, but stored in server to help form an event
const (
	RuntimeSince    = "rtSince"
	SessionSince    = "sessSince"
	SocketSince     = "socketSince"
	UserAgent       = "userAgent"
	Platform        = "platform"
	OperatingSystem = "os"
	OSVersion       = "osVersion"
)
