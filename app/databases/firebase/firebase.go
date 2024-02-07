package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func ConnectFirebase() *messaging.Client {

	opt := option.WithCredentialsFile("firebase-service-account-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Errorf("failed to initialize Firebase app: %v", err)
		return nil
	}

	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		logrus.Errorf("failed to connect to Firebase Cloud Messaging: %v", err)
		return nil
	}

	logrus.Info("connected to Firebase Cloud Messaging")
	return fcmClient
}
