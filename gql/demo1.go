package gql

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func GqlRun() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
			{
				hello
			}
		`
	params := graphql.Params{Schema: schema, RequestString: query}

	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}

type DirectionsResponse struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
}

type DirectionsRequest struct {
	Waypoints   string
	AccessToken string
}

func Mapbox() {
	access_token := "sk.eyJ1IjoiZGVzdGVrIiwiYSI6ImNsZ3oyeHYxNTBmazkzZW05NWg4MjM3YXcifQ.uKkAmQL-tYu93iqhVda6Qg"
	waypoints := "-122.42,37.78;-77.03,38.91"
	directions_request := DirectionsRequest{waypoints, access_token}
	api_url := fmt.Sprintf("https://api.mapbox.com/directions/v5/mapbox/driving/%s.json?access_token=%s", directions_request.Waypoints, directions_request.AccessToken)

	response, error := http.Get(api_url)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	var directions_response DirectionsResponse
	error = json.NewDecoder(response.Body).Decode(&directions_response)
	if error != nil {
		panic(error)
	}

	route := directions_response.Routes[0]
	distance := route.Distance
	duration := route.Duration

	fmt.Printf("Distance: %v meters\n", distance)
	fmt.Printf("Duration: %v seconds\n", duration)
}
