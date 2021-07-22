package handler

import (
	"content-management-api/cache"
	"content-management-api/driver/mongo"
	"content-management-api/gateway"
	"content-management-api/handler/validator"
	"content-management-api/usecase"
	"content-management-api/util/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type instance struct {
	port   int
	server *echo.Echo
}

type Server interface {
	Start()
}

func (e *instance) Start() {
	fmt.Printf(banner, version, website)
	e.server.Logger.Fatal(e.server.Start(fmt.Sprintf(":%v", e.port)))
}

const (
	// Version of Sleepy-Hollow Content-Management-API
	version = "0.0.0"
	website = "https://sleepy-hollow.io"
	// http://patorjk.com/software/taag/#p=display&f=Small%20Slant&t=Echo
	banner = `
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNkkkkkkkkkkkkMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNkkkkkkkkkkMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#Y=H@@H@@H@HHTWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMM"""~<<("T"Y<~~~~d@H@@H@@H@::<?WMH8Y3<<<YYMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMM"!.~.~~~~~~~_____~:~d@@H@@H@@K::;;<<<;;;;;>>>>>>1TMMMMMMMMMMMMMMM
MMMMMMMMMMMM9~..~.~~.~(jXH@@H@@@@HmH@H@@H@@@MkH@@@@H@MHky+>>>>????zTMMMMMMMMMMMM
MMMMMMMMMM9~..~.~~.~~~~~~~~?TM@H@@H@MMMMMMMMH@@H@HM9C<>>>>>>????????zTMMMMMMMMMM
MMMMMMMM#!..~.~~.~~~~~~~~~~~:~?HBT<::::::;;;<?THBC>;>>>>>>??>???????==zMMMMMMMMM
MMMMMMM@..~..~.~~~~~~~~~~~~:~:::::::::::;:;;;;;;;;>>>>>>>??>??????==????MMMMMMMM
MMMMMM@.~~.~~~~~~~~~~~~~~:~::~:::::::::;;;:;;;;;>>>>>>>??>???????=?==?=?=HMMMMMM
MMMMM#..~.~.~~~.~~~~~~~~:~:~::~:::::::;;;;;;;;>>>>>>>>????????=?=?=?==?=?=MMMMMM
MMMMM>.~.~~~~~~~~~~~~~:~:::::::::::::;;;;;;;;>>>>>>>?????????=?==?=??==?=?vMMMMM
MMMMF.~~~~~~~~~~~~~~MNgJ-:~:::::::::;;;;;;;;>>>>>>>????zugMM=?=?==?==?=?==?dMMMM
MMMM$~~.~~.~~~~~~~~:(MMMMNNgJ<::::;;;;;;;;>>>>>>>?1ggNMMMMM$?==??=?=?=?=?==dMMMM
MMMM~~~~~~~~~~~~~:~::dMMMMMMMMMNe;;;;;;;>>>>>>>qNMMMMMMMMM@=?=?==?=?=?=?=?==MMMM
MMMM~.~~~~~~~~~~:~:~::dMMMMMMMBC;;;;;;>>>>>>>???vTMMMMMMME=?=?=?==?==?==?=??MMMM
MMMM~~~~~~~~~~::~:::::::T""5C;;;;;;;>>>>>>>?>??????vdBW8?=?==?=??=?=?=?=?==?MMMM
MMMM~~~~~~~~:~:~::::::::::;;;;;;;;;>>>>>>>??????????=?=?=?=?==?==?=?=?=?=?==MMMM
MMMM~~~~~~:~::~::::~:::::;;;;;;;;>>>>>>>??>???????==?=?==?=??=?=?==?==?=?=??MMMM
MMMMF~~~~:~:~:::::?MMMMb;;;;;;MMMMMM#>?>???MMMMMMM?==?=?dMMMM#=?=?==?=?==?=dMMMM
MMMMN_~::~:::::::::TMMM#;;;;;;MMMMMM#>?????MMMMMMM?=?==?dMMM5=?=?=??=?=?==1MMMMM
MMMMMp~:~::~::::::::?HMNJJJJJ(MMMMMMNgJJJJ+MMMMMMMe+++++dM#6=?=?==?==?=??=qMMMMM
MMMMMMx~:::::::::::;;;?MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM6=?=?==?==?==?==dMMMMMM
MMMMMMN+:::::::::;;;;;;;<?HM>>>>??dMMMMMMMMMM#=====?MM9z?=?=?=?=??=??=?=1MMMMMMM
MMMMMMMMJ:::::::;;;;;;;>>>>>>?????dMMMMMMMMMM#?==?==?==?==?==?=?==?==?=aMMMMMMMM
MMMMMMMMMR::::;;;;;;;;>>>>>>?>??????=??TTT6==?=??=?=?=?=?==?=?==?==?=?dMMMMMMMMM
MMMMMMMMMMNe;;;;;;;;>>>>>>>???>?????==?==?=?=?==?==?=?=?=??=?=?=??=?gMMMMMMMMMMM
MMMMMMMMMMMMNg+;;;>>>>>>>>??>?????=?=?=?==?=?=?==?=?==?==?==?==?=1gMMMMMMMMMMMMM
MMMMMMMMMMMMMMMNgJ>>>>>>????????=?==?=?=?=?==?=??=?=?=?=?==?=?ugMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMN+gggggggNNgz?=?=?==?=?=?=?=1ggMNNgggggggNMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNaggzzzzzzugggMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
--------------------------------------------------------------------------------
Version: %s
Web: %s
`
)

func NewServer(container cache.Cache) Server {
	// Echo instance
	e := echo.New()
	e.HideBanner = true
	// set Validator
	e.Validator = validator.NewValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/v1/systems/ping", pong)

	routing(e, container)

	return &instance{
		port:   config.Config.Server.Port,
		server: e,
	}
}

func routing(e *echo.Echo, container cache.Cache) *echo.Echo {
	// TODO move to outside of handler
	db, err := container.Load(cache.MongoDB)
	if err != nil {
		return nil
	}
	mongoContentDriver := mongo.NewContentDriver(db.(*mongo.Client))

	contentModelGateway := gateway.NewContentModel(mongoContentDriver)
	contentModelResource := NewContentModelResource(usecase.NewContentModel(contentModelGateway))
	entryResource := NewEntryResource(usecase.NewEntry(gateway.NewEntry(mongoContentDriver), contentModelGateway))
	spaceResource := NewSpaceResource(usecase.NewSpace(gateway.NewSpace()))

	contentModelResource.Routing(e)
	entryResource.Routing(e)
	spaceResource.Routing(e)

	return e
}

type ErrorResponse struct {
	Message string `json:"message"`
}
