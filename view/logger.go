package view

import (
	"github.com/blue-army-2017/knight/util"
	"go.uber.org/zap"
)

var l *zap.SugaredLogger

func init() {
	l = util.GetLogger()
}
