package wego

import (
	"github.com/godcong/wego/core"
)

type Auth interface {
	Session(code string) core.Map
}

type DataCube interface {
	UserPortrait(from, to string) core.Map
	SummaryTrend(from, to string) core.Map
	DailyVisitTrend(from, to string) core.Map
	WeeklyVisitTrend(from, to string) core.Map
	MonthlyVisitTrend(from, to string) core.Map
	VisitDistribution(from, to string) core.Map
	DailyRetainInfo(from, to string) core.Map
	WeeklyRetainInfo(from, to string) core.Map
	MonthlyRetainInfo(from, to string) core.Map
	VisitPage(from, to string) core.Map
}

type AppCode interface {
	Get(path string, optionals core.Map) core.Map
	GetQrCode(path string, width int) core.Map
	GetUnlimit(scene string, optionals core.Map) core.Map
}

type MiniProgram interface {
	Auth() Auth
	AppCode() AppCode
	//Client() core.Client
	DataCube() DataCube
	AccessToken() AccessToken
}

//func NewAppCode(application core.Application, config core.Config) AppCode {
//	return &mini_program.AppCode{
//		Config: config,
//		//mini_program:   application.MiniProgram(),
//	}
//}

func GetMiniProgram() MiniProgram {
	obj := GetApp().Get("mini_program").(MiniProgram)
	core.Debug("GetMiniProgram|obj:", obj)
	return obj
}

func GetAuth() Auth {
	return GetMiniProgram().Auth()
}

func GetAppCode() AppCode {
	return GetMiniProgram().AppCode()
}

func GetDataCube() DataCube {
	return GetMiniProgram().DataCube()
}
