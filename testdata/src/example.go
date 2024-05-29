package src

import "context"

func example() {
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		ctx := context.WithValue(ctx, "key", i)
		ctx = context.WithValue(ctx, "other", "val")
	}

	for i := 0; i < 10; i++ {
		ctx = context.WithValue(ctx, "key", i) // want "nested context in loop"
		ctx = context.WithValue(ctx, "other", "val")
	}

	for item := range []string{"one", "two", "three"} {
		ctx = wrapContext(ctx) // want "nested context in loop"
		ctx := context.WithValue(ctx, "key", item)
		ctx = wrapContext(ctx)
	}

	for {
		ctx = wrapContext(ctx) // want "nested context in loop"
		break
	}
}

func wrapContext(ctx context.Context) context.Context {
	return context.WithoutCancel(ctx)
}

// storing contexts in a struct isn't recommended, but local copies of a non-pointer struct should act like local copies of a context.
func inStructs(ctx context.Context) {
	for i := 0; i < 10; i++ {
		c := struct{ Ctx context.Context }{ctx}
		c.Ctx = context.WithValue(c.Ctx, "key", i)
		c.Ctx = context.WithValue(c.Ctx, "other", "val")
	}

	for i := 0; i < 10; i++ {
		c := []struct{ Ctx context.Context }{{ctx}}
		c[0].Ctx = context.WithValue(c[0].Ctx, "key", i)
		c[0].Ctx = context.WithValue(c[0].Ctx, "other", "val")
	}

	c := struct{ Ctx context.Context }{ctx}
	for i := 0; i < 10; i++ {
		c := c
		c.Ctx = context.WithValue(c.Ctx, "key", i)
		c.Ctx = context.WithValue(c.Ctx, "other", "val")
	}

	pc := &struct{ Ctx context.Context }{ctx}
	for i := 0; i < 10; i++ {
		c := pc
		c.Ctx = context.WithValue(c.Ctx, "key", i) // want "nested context in loop"
		c.Ctx = context.WithValue(c.Ctx, "other", "val")
	}

	for i := 0; i < 10; i++ {
		c := []*struct{ Ctx context.Context }{{ctx}}
		c[0].Ctx = context.WithValue(c[0].Ctx, "key", i) // want "nested context in loop"
		c[0].Ctx = context.WithValue(c[0].Ctx, "other", "val")
	}
}
