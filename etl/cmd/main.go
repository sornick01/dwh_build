package main

import (
	"context"
	"etl/etl/api/etl"
	"etl/internal/config"
	"etl/internal/implementation"
	"etl/internal/repository"
	"etl/internal/service"
	"etl/internal/service/migrator"
	"etl/internal/service/validator"
	"etl/pkg/postgres"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.New("./")
	f, err := os.OpenFile(cfg.PathLog+"/log.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open log file: %s", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	log.SetOutput(f)

	poolCfgs, err := postgres.New(cfg.ConfigDBURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer poolCfgs.Close()

	poolSrc, err := postgres.New(cfg.SrcDBURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer poolSrc.Close()

	poolDst, err := postgres.New(cfg.DstDBURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer poolDst.Close()

	// init etl
	repo := repository.NewRepository(poolCfgs, poolSrc, poolDst)
	val := &validator.Validator{}
	migr := migrator.NewMigrator(val, repo)
	svc := service.NewService(migr, repo)
	impl := implementation.NewImplementation(svc, repo)

	lis, err := net.Listen("tcp", cfg.PortGRPC)
	if err != nil {
		log.Fatalln("failed to listen: ", err.Error())
	}
	s := grpc.NewServer()

	etl.RegisterEtlServer(s, impl)

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalln("failed to serve: ", err)
		}
	}()

	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	//Graceful	Shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		s.GracefulStop()

		serverStopCtx()
	}()

	log.Println("Serving gRPC on port ", cfg.PortGRPC)
	//db := &domain.Database{}
	//routes := &domain.Routes{}
	//
	//dbDescr, err := os.ReadFile("jsons/db.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//routesDescr, err := os.ReadFile("jsons/routes.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = json.Unmarshal(dbDescr, db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = json.Unmarshal(routesDescr, routes)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(db.ToSql())
	//fmt.Println(routes.ToSql())
	//conn, err := pgx.Connect(context.Background(), "postgresql://postgres:1234@localhost:5432/dst")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = conn.Exec(context.Background(), db.ToSql())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = conn.Exec(context.Background(), routes.ToSql())
	//if err != nil {
	//	log.Fatal(err)
	//}
}
