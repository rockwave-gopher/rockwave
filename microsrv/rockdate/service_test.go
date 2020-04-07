package rockdate

import (
	"context"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()
	s, err := srv.Status(ctx)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	ok := s == "ok"
	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func TestGet(t *testing.T) {
	srv, ctx := setup()
	d, err := srv.Get(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	time := time.Now()
	today := time.Format("02/01/2006")
	ok := today == d

	if !ok {
		t.Errorf("expected dates to be equal")
	}
}

func TestValidate(t *testing.T) {
	srv, ctx := setup()

	b, err := srv.Validate(ctx, "04/07/2020")

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !b {
		t.Errorf("date should be valid")
	}

	b, err = srv.Validate(ctx, "31/31/2020")
	if b {
		t.Errorf("date should be invalid")
	}

	b, err = srv.Validate(ctx, "12/31/2019")
	if b {
		t.Errorf("USA date should be invalid")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
