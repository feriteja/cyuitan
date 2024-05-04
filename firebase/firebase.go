package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

func Init() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

}
