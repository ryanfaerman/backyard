package picket

var blacklist = []string{}

func AddToBlacklist(url string) {
	blacklist = append(blacklist, url)
}

func ListBlacklist() []string {
	return blacklist
}
