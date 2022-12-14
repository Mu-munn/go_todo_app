package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	// err := http.ListenAndServe(":18080",
	// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	// 	}),
	// )
	// if err != nil {
	// 	fmt.Printf("faild to start http server : %v", err)
	// 	os.Exit(1)
	// }
	if err := run(context.Background()); err != nil {
		log.Printf("faild to start http server : %v", err)
	}
}
func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンでHTTPサーバーを起動する
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("faild to start http server : %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("faild to shutdown : %+v", err)
	}

	return eg.Wait()
}
