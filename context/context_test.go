package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	fmt.Println(ctx)
}

func process(ctx context.Context) {
	// ambil value dari context
	value := ctx.Value("userID")

	fmt.Println("User ID:", value)
}

func TestWithValue(t *testing.T) {
	ctx := context.Background()

	// tambahkan value ke context
	ctx = context.WithValue(ctx, "userID", 123)

	// kirim context ke fungsi proses
	process(ctx)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func TestWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)

	fmt.Println("Main: canceling...")
	cancel() // kirim sinyal pembatalan ke worker

	time.Sleep(1 * time.Second)
}

func TestWithTimeOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // pastikan cancel dipanggil untuk membersihkan resource

	go worker(ctx)

	time.Sleep(5 * time.Second)
}

func TestWithDeadline(t *testing.T) {
	deadline := time.Now().Add(3 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // pastikan cancel dipanggil untuk membersihkan resource

	go worker(ctx)

	time.Sleep(5 * time.Second)
}
