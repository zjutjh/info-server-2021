package model

type Options struct {
	ConfigPath string `short:"c" long:"config" description:"[PATH] Config file path"`
	LoadData   string `short:"l" long:"load" description:"[PATH] Read & load data from excel"`
	Passwd     string `short:"p" long:"passwd" description:"[PASS] Password of excel file"`
	Sheet      string `short:"s" long:"sheet" description:"[Sheet] Read sheet"`
	Update     bool   `short:"u" long:"update" description:"Update database by dormitory info"`
	Version    bool   `short:"v" long:"version" description:"Show Info server version & quit"`
}
