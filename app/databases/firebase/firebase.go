package firebase

import (
	"context"
	"talkspace/app/configs"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func ConnectFirebase() *messaging.Client {

	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load firebase configuration: %v", err)
		return nil
	}

	opt := option.WithCredentialsFile(config.FIREBASE.FIREBASE_CREDENTIALS_FILE)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logrus.Errorf("failed to initialize firebase app: %v", err)
		return nil
	}

	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		logrus.Errorf("failed to connect to firebase cloud messaging: %v", err)
		return nil
	}

	logrus.Info("connected to firebase cloud messaging")
	return fcmClient
}