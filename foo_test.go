package datastoreemulatorplayground

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestFoo(t *testing.T) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	Expect(t, client).Not().ToBe(nil)
	Expect(t, err).ToBe(nil)

	n, err := client.Count(ctx, datastore.NewQuery("users"))
	Expect(t, err).ToBe(nil)
	Expect(t, n).ToBe(0)

	err = client.Close()
	Expect(t, err).ToBe(nil)
}
