package switcheroo

import (
	"context"
	"strconv"
	"testing"
)

func TestExample(t *testing.T) {
	ctx := context.WithValue(context.Background(), "total", 1)

	r := New(ctx)

	r.Add("/add/{number}", func(ctx context.Context, params map[string]string) {
		total := ctx.Value("total").(int)
		number, _ := strconv.Atoi(params["number"])
		total += number

		if total != 10 {
			t.Fatalf("total should have been 10. Got: %v", total)
		}
	})

	r.Run("/add/9")
}

func TestMatchedRun(t *testing.T) {
	ctx := context.WithValue(context.Background(), "ctxkey", "prefix|")

	toBeTested := ""

	r := New(ctx)

	r.Add("/a/{name}", func(ctx context.Context, params map[string]string) {
		toBeTested += ctx.Value("ctxkey").(string) + params["name"]
	})

	r.Run("/a/didip")

	if toBeTested != "prefix|didip" {
		t.Fatalf("Failed to properly match and execute handler. Got %v", toBeTested)
	}
}

func TestExactMatchRun(t *testing.T) {
	ctx := context.WithValue(context.Background(), "ctxkey", "prefix|")

	toBeTested := ""

	r := New(ctx)

	r.Add("/a", func(ctx context.Context, params map[string]string) {
		toBeTested = "matched"
	})

	r.Run("/a")

	if toBeTested != "matched" {
		t.Fatalf("Failed to properly match and execute handler. Got %v", toBeTested)
	}
}

func TestMissRun(t *testing.T) {
	ctx := context.WithValue(context.Background(), "ctxkey", "prefix|")

	toBeTested := ""

	r := New(ctx)

	r.Add("/a/{name}", func(ctx context.Context, params map[string]string) {
		toBeTested += ctx.Value("ctxkey").(string) + params["name"]
	})

	r.Run("/total/random")

	if toBeTested != "" {
		t.Fatalf("input string should never match anything.")
	}
}
