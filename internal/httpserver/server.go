package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Craftbec/Shortener_link/config"
	pb "github.com/Craftbec/Shortener_link/internal/linkshorter"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func ConnectGRPC(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Failed to dial server: %v", err)
	}
	return conn, nil
}

func GET(w http.ResponseWriter, r *http.Request, client pb.ShortenerClient) {
	short := mux.Vars(r)["shortLink"]
	response, err := client.Get(r.Context(), &pb.ShortLink{Link: short})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.Link))
}

func POST(w http.ResponseWriter, r *http.Request, client pb.ShortenerClient) {
	original := r.URL.Query()
	_, err := url.ParseRequestURI(original["url"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect url\n"))
		return
	}
	response, err := client.Post(r.Context(), &pb.OriginalLink{Link: original["url"][0]})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.Link))
}

func HTTPServer(ctx context.Context, conf *config.Config) error {
	r := mux.NewRouter()
	connect, err := ConnectGRPC(fmt.Sprintf(":%v", (*conf).GRCP.Port))
	if err != nil {
		return err
	}
	client := pb.NewShortenerClient(connect)
	r.HandleFunc("/{shortLink}", func(w http.ResponseWriter, r *http.Request) {
		GET(w, r, client)
	}).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		POST(w, r, client)
	}).Methods("POST")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.HTTP.Port),
		Handler: r,
	}
	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println(err)
		}
	}()
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil

}
