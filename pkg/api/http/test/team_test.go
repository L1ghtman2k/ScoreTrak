package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/server/gorilla"
	"github.com/ScoreTrak/ScoreTrak/pkg/api/http/client"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	. "github.com/ScoreTrak/ScoreTrak/pkg/logger/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"

	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestTeamSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_api_team"
	c.Logger.FileName = "team_test.log"
	db := storage.SetupDB(c.DB)
	l := SetupLogger(c.Logger)
	rtr := gorilla.NewRouter()
	routes := gorilla.Routes{
		gorilla.Route{
			Name:        "Index",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: gorilla.Index,
		},
	}
	cr := orm.NewTeamRepo(db, l)
	teamSvc := team.NewTeamServ(cr)
	routes = append(routes, gorilla.TeamRoutes(l, teamSvc)...)
	for _, route := range routes {
		hdler := route.HandlerFunc
		rtr.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(hdler)
	}
	rtr.Use(gorilla.JsonHeader)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	go http.Serve(listener, rtr)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Initializing team repo and controller", t, func() {
		CreateAllTables(db)
		DataPreload(db)
		s := client.NewScoretrakClient(&url.URL{Host: fmt.Sprintf("localhost:%d", port), Scheme: "http"}, "", http.DefaultClient)
		cli := client.NewTeamClient(s)
		Convey("Retrieving a team by Name", func() {
			retTeam, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
			So(err, ShouldBeNil)
			So(retTeam.Name, ShouldEqual, "TeamOne")
			So(*(retTeam.Enabled), ShouldBeTrue)
		})
		//Convey("Retrieving a team by wrong ID", func() {
		//	retTeam, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111122"))
		//	So(err, ShouldNotBeNil)
		//	So(retTeam, ShouldBeNil)
		//	seer, ok := err.(*client.InvalidResponse)
		//	So(ok, ShouldBeTrue)
		//	So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusNotFound)
		//})
		//
		//Convey("Updating a team by Name", func() {
		//	fls := false
		//	t := team.Team{ID: uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"), Name: "TeamOneZZZ", Enabled: &fls}
		//	err := cli.Update(&t)
		//	So(err, ShouldBeNil)
		//	Convey("Retrieving a team by Name", func() {
		//		retTeam, err := cli.GetByID(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
		//		So(err, ShouldBeNil)
		//		So(retTeam.Name, ShouldEqual, "TeamOneZZZ")
		//		So(*(retTeam.Enabled), ShouldBeFalse)
		//	})
		//})
		//
		//Convey("Getting all teams", func() {
		//	teams, err := cli.GetAll()
		//	So(err, ShouldBeNil)
		//	So(len(teams), ShouldEqual, 4)
		//	var IDs []string
		//	for _, tm := range teams {
		//		IDs = append(IDs, tm.Name)
		//	}
		//	So(IDs, ShouldContain, "TeamTwo")
		//})
		//
		//Convey("Deleting a team that doesnt have child hosts by Name", func() {
		//	err := cli.Delete(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
		//	So(err, ShouldBeNil)
		//	Convey("Getting all teams", func() {
		//		teams, err := cli.GetAll()
		//		So(err, ShouldBeNil)
		//		So(len(teams), ShouldEqual, 3)
		//	})
		//})
		//
		//Convey("Deleting a team that does have child hosts by Name", func() {
		//	err := cli.Delete(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
		//	So(err, ShouldNotBeNil)
		//	seer, ok := err.(*client.InvalidResponse)
		//	So(ok, ShouldBeTrue)
		//	So(seer.ResponseCode, ShouldHaveSameTypeAs, http.StatusConflict)
		//	Convey("Getting all teams", func() {
		//		teams, err := cli.GetAll()
		//		So(err, ShouldBeNil)
		//		So(len(teams), ShouldEqual, 4)
		//	})
		//})
		//
		//Convey("Deleting a non existent team", func() {
		//	err := cli.Delete(uuid.FromStringOrNil("11111111-1111-1111-1111-111111111111"))
		//	So(err, ShouldBeNil)
		//})
		//
		//Convey("Storing a new team", func() {
		//	fls := false
		//	t := []*team.Team{{Name: "TeamFive", Enabled: &fls}}
		//	err := cli.Store(t)
		//	So(err, ShouldBeNil)
		//	Convey("Getting all teams", func() {
		//		teams, err := cli.GetAll()
		//		So(err, ShouldBeNil)
		//		So(len(teams), ShouldEqual, 5)
		//	})
		//})

		Reset(func() {
			CleanAllTables(db)
		})
	})
	DropDB(db, c)

}