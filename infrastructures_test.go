package deferr_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/marconi/deferr"
)

type FakeStoreHandler struct {
	deferr.StoreHandler
}

func (fsh *FakeStoreHandler) Size() int {
	return len(fsh.Items)
}

func TestStoreHandlerSpec(t *testing.T) {
	handler := &FakeStoreHandler{}

	Convey("testing store handler", t, func() {
		Convey("should be able to push item", func() {
			t := &deferr.Todo{Name: "Wash clothes"}
			err := handler.Push(t)
			So(err, ShouldBeNil)
			So(handler.Size(), ShouldEqual, 1)
		})

		Convey("should be able to query items", func() {
			items := handler.Query()
			So(items, ShouldNotBeNil)
			So(len(items), ShouldEqual, 1)
		})

		Convey("should be able to defer item", func() {
			t := &deferr.Todo{Name: "Sweep the floor"}
			handler.Push(t)
			So(handler.Size(), ShouldEqual, 2)

			items := handler.Query()
			So(items[0].(*deferr.Todo).Name, ShouldEqual, "Wash clothes")
			So(items[1].(*deferr.Todo).Name, ShouldEqual, "Sweep the floor")

			err := handler.Defer()
			So(err, ShouldBeNil)

			items = handler.Query()
			So(items[0].(*deferr.Todo).Name, ShouldEqual, "Sweep the floor")
			So(items[1].(*deferr.Todo).Name, ShouldEqual, "Wash clothes")
		})

		Convey("should be able to pop item", func() {
			item, err := handler.Pop()
			So(err, ShouldBeNil)
			So(item, ShouldNotBeNil)
			So(handler.Size(), ShouldEqual, 1)

			handler.Pop()
			item, err = handler.Pop()
			So(err, ShouldNotBeNil)
			So(item, ShouldBeNil)
		})
	})
}
