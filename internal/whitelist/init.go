package whitelist

import "sync"

type WhitelistGroup struct {
	wg *sync.WaitGroup
}
