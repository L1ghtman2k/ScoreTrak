package orm

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/service"
	. "ScoreTrak/test"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func TestCheckSpec(t *testing.T) {
	var c *config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewConfigClone(SetupConfig("../../../configs/test-config.yml"))
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_orm_check"
	c.Logger.FileName = "check_test.log"
	db := SetupDB(c)
	l := SetupLogger(c)
	t.Parallel() //t.Parallel should be placed after SetupDB because gorm has race conditions on Hook register
	Convey("Creating Round, Service and Check tables along with their foreign keys", t, func() {
		db.AutoMigrate(&service.Service{})
		db.AutoMigrate(&round.Round{})
		db.AutoMigrate(&check.Check{})
		db.Model(&check.Check{}).AddForeignKey("service_id", "services(id)", "CASCADE", "RESTRICT")
		db.Model(&check.Check{}).AddForeignKey("round_id", "rounds(id)", "CASCADE", "RESTRICT")
		cr := NewCheckRepo(db, l)
		Convey("When all tables are empty", func() {
			Convey("Should output no entry", func() {
				gt, err := cr.GetAll()
				So(err, ShouldBeNil)
				So(len(gt), ShouldEqual, 0)
			})
			Convey("Creating a sample check should not be allowed", func() {
				c := check.Check{}
				err := cr.Store(&c)
				So(err, ShouldNotBeNil)
				ac, err := cr.GetAll()
				So(err, ShouldBeNil)
				So(len(ac), ShouldEqual, 0)
			})
			Convey("Load sample services and rounds", func() {
				var count int
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES (5, 999, 999, 'TestService')")
				db.Exec("INSERT INTO services (id, service_group_id, host_id, name) VALUES (6, 999, 999, 'TestService')")
				db.Exec("INSERT INTO rounds (id, start) VALUES (1, $1)", time.Now())
				db.Exec("INSERT INTO rounds (id, start) VALUES (2, $1)", time.Now())
				db.Table("services").Count(&count)
				So(count, ShouldEqual, 2)
				db.Table("rounds").Count(&count)
				So(count, ShouldEqual, 2)
				Convey("Creating a sample check and associating with service 5 and round 1", func() {
					c := check.Check{Log: "TestLog", ServiceID: 5, RoundID: 1}
					err := cr.Store(&c)
					Convey("Should be Allowed", func() {
						So(err, ShouldBeNil)
						ac, err := cr.GetAll()
						So(err, ShouldBeNil)
						So(len(ac), ShouldEqual, 1)
						Convey("Then Querying By ID", func() {
							ss, err := cr.GetByID(c.ID)
							So(err, ShouldBeNil)
							So(ss.Log, ShouldEqual, "TestLog")
							So(ss.ServiceID, ShouldEqual, 5)
							So(ss.RoundID, ShouldEqual, 1)
						})

						Convey("Then Querying By wrong ID", func() {
							ss, err := cr.GetByID(c.ID + 1)
							So(err, ShouldNotBeNil)
							So(ss, ShouldBeNil)
						})

						Convey("Then Deleting the check should be allowed", func() {
							err = cr.Delete(c.ID)
							So(err, ShouldBeNil)
							ac, err = cr.GetAll()
							So(err, ShouldBeNil)
							So(len(ac), ShouldEqual, 0)
						})

						Convey("Then Deleting a wrong entry", func() {
							err = cr.Delete(c.ID + 1)
							So(err, ShouldNotBeNil)
							Convey("Should output one entry", func() {
								ac, err := cr.GetAll()
								So(err, ShouldBeNil)
								So(len(ac), ShouldEqual, 1)
							})
						})

						Convey("Then adding more checks. One with similar round, and the other with similar service", func() {
							c1 := check.Check{Log: "TestLog2", ServiceID: 6, RoundID: 1}
							c2 := check.Check{Log: "TestLog", ServiceID: 5, RoundID: 2}
							checks := []*check.Check{&c1, &c2}
							err = cr.StoreMany(checks)
							So(err, ShouldBeNil)
							Convey("Query by round ID should return 2 entries", func() {
								checks, err := cr.GetAllByRoundID(1)
								So(err, ShouldBeNil)
								So(len(checks), ShouldEqual, 2)
							})
							Convey("Query by round ID and Service ID should return 1 entry", func() {
								checks, err := cr.GetByRoundServiceID(1, 5)
								So(err, ShouldBeNil)
								So(len(checks), ShouldEqual, 1)
							})
						})
					})
				})
				Convey("Creating a check with wrong service should not be allowed", func() {
					s := check.Check{Log: "TestLog", ServiceID: 999, RoundID: 1}
					err := cr.Store(&s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

				Convey("Creating a check with wrong round should not be allowed", func() {
					s := check.Check{Log: "TestLog", ServiceID: 5, RoundID: 4}
					err := cr.Store(&s)
					So(err, ShouldNotBeNil)
					ac, err := cr.GetAll()
					So(err, ShouldBeNil)
					So(len(ac), ShouldEqual, 0)
				})

			})
		})
		Reset(func() {
			db.DropTableIfExists(&check.Check{})
			db.DropTableIfExists(&round.Round{})
			db.DropTableIfExists(&service.Service{})
		})
	})
	DropDB(db, c)
	db.Close()
}
