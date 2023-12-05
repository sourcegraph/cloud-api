package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	cloudapiv1 "github.com/sourcegraph/cloud-api/go/cloudapi/v1"
	"github.com/sourcegraph/cloud-api/go/cloudapi/v1/cloudapiv1connect"
	"google.golang.org/api/idtoken"

	"connectrpc.com/connect"
)

func main() {
	client := cloudapiv1connect.NewInstanceServiceClient(
		http.DefaultClient,
		os.Getenv("CAPI_SERVER"),
	)

	req := connect.NewRequest(&cloudapiv1.ListInstancesRequest{})
	req.Header().Add(getAuthHeader())

	res, err := client.ListInstances(
		context.Background(),
		req,
	)
	if err != nil {
		log.Fatalln(err)
	}
	out := fmt.Sprintf("# of instances: %d", len(res.Msg.Instances))
	if res.Msg.Drift != nil {
		out += fmt.Sprintf(" drift - onlyRepo: %d onlyProject: %d", len(res.Msg.Drift.OnlyRepo), len(res.Msg.Drift.OnlyProject))
	}
	log.Println(out)
}

func getAuthHeader() (string, string) {
	if path, ok := os.LookupEnv("CAPI_GCP_SA_KEY"); ok {
		if aud, ok := os.LookupEnv("CAPI_GCP_SA_AUDIENCE"); ok {
			ts, err := idtoken.NewTokenSource(context.Background(), aud, idtoken.WithCredentialsFile(path))
			if err != nil {
				log.Fatalf("unable to create TokenSource: %v", err)
			}
			tok, err := ts.Token()
			if err != nil {
				log.Fatalf("unable to obtain GCP SA Token: %v", err)
			}
			validTok, err := idtoken.Validate(context.Background(), tok.AccessToken, aud)
			if err != nil {
				log.Fatalf("token validation failed: %v", err)
			}
			fmt.Println("Using GCP SA token for:", validTok.Claims["email"])
			return "X-ROBOT-TOKEN", tok.AccessToken
		} else {
			log.Fatalln("CAPI_GCP_SA_KEY set but CAPI_GCP_SA_AUDIENCE not provided")
		}
	}
	if _, ok := os.LookupEnv("CAPI_TOKEN"); !ok {
		log.Fatalln("Neither key (CAPI_GCP_SA_KEY) nor token (CAPI_TOKEN) provided")
	}
	fmt.Println("Using per-user token")
	return "X-USER-TOKEN", os.Getenv("CAPI_TOKEN")
}
