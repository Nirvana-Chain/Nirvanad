package consensus

import (
	"github.com/Nirvana-Chain/nirvanad/infrastructure/logger"
	"github.com/Nirvana-Chain/nirvanad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
