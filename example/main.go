package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/google/uuid"
	"github.com/igmagollo/undine/example/ent"
	undine "github.com/igmagollo/undine/pkg/v1"
)

func main() {
	postgres := embeddedpostgres.NewDatabase()
	if err := postgres.Start(); err != nil {
		panic(err)
	}
	defer func() {
		if err := postgres.Stop(); err != nil {
			panic(err)
		}
	}()

	logger := watermill.NewStdLogger(false, false)
	pubsub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable",
		ent.DeduplicatorSchemaAdapter(&undine.DeduplicatorPostgresSchemaAdapter{}),
		ent.WatermillLogger(logger),
		ent.OutboxOffsetsAdapter(&sql.DefaultPostgreSQLOffsetsAdapter{}),
		ent.OutboxSchemaAdapter(&sql.DefaultPostgreSQLSchema{}),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	fmt.Println("initializing schema")
	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println("initializing deduplicator schema")
	if err := client.InitializeDeduplicatorSchema(context.Background()); err != nil {
		panic(err)
	}

	forwarder, err := client.Forwarder("example_consumer_group", pubsub)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := forwarder.Run(context.Background()); err != nil {
			panic(err)
		}
	}()

	// await for forwarder create the outbox table
	time.Sleep(1 * time.Second)

	createUserTopic, _ := pubsub.Subscribe(context.Background(), "create_user")
	userCreatedTopic, _ := pubsub.Subscribe(context.Background(), "user_created")

	go HandleUserCreateMessage(createUserTopic, client)
	go HandleUserCreatedMessage(userCreatedTopic, client)

	stop := make(chan struct{})
	go func() {
		for msg := range generateCreateUserMessages() {
			select {
			case <-stop:
				return
			default:
			}
			fmt.Println("publishing create user", msg.UUID)
			if err := pubsub.Publish("create_user", msg); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	stop <- struct{}{}

	forwarder.Close()
	pubsub.Close()
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) String() string {
	return fmt.Sprintf("User{Name: %s, Email: %s, Password: %s}", u.Name, u.Email, u.Password)
}

func CreateUser(ctx context.Context, client *ent.Client, user User) error {
	fmt.Println("creating user", user)
	return client.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Exec(ctx)
}

type CreateUserMessage struct {
	User User `json:"user"`
}

func NewCreateUserMessage(user User) *message.Message {
	msg := &CreateUserMessage{
		User: user,
	}
	msgb, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return message.NewMessage(watermill.NewUUID(), msgb)
}

type UserCreatedMessage struct {
	User User `json:"user"`
}

func NewUserCreatedMessage(user User) *message.Message {
	msg := &UserCreatedMessage{
		User: user,
	}
	msgb, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return message.NewMessage(watermill.NewUUID(), msgb)
}

func generateCreateUserMessages() <-chan *message.Message {
	ch := make(chan *message.Message)
	lastMessage := NewCreateUserMessage(User{
		Name:     fmt.Sprintf("user-%d", -1),
		Email:    fmt.Sprintf("email-%d", -1),
		Password: fmt.Sprintf("password-%d", -1),
	})
	go func() {
		i := 0
		for true {
			// 10% of the time, we'll send the same user as the last one
			if rand.Intn(10) != 0 {
				lastMessage = NewCreateUserMessage(User{
					Name:     fmt.Sprintf("user-%d", i),
					Email:    fmt.Sprintf("email-%d", i),
					Password: fmt.Sprintf("password-%d", i),
				})
			}
			ch <- lastMessage
			i++
		}
		close(ch)
	}()
	return ch
}

func HandleUserCreateMessage(createUserTopic <-chan *message.Message, client *ent.Client) {
	for msg := range createUserTopic {
		ctx := msg.Context()
		var createUserMessage CreateUserMessage
		if err := json.Unmarshal(msg.Payload, &createUserMessage); err != nil {
			panic(err)
		}
		err := client.WithTx(ctx, func(ctx context.Context) error {
			tx := ent.TxFromContext(ctx)
			deduplicator := tx.Deduplicator()
			publisher, err := tx.OutboxPublisher()
			if err != nil {
				return err
			}

			messageID := uuid.MustParse(msg.UUID)
			err = deduplicator.Deduplicate(ctx, "create_user", messageID)
			if err != nil {
				return err
			}

			if err := CreateUser(ctx, tx.Client(), createUserMessage.User); err != nil {
				return err
			}

			err = publisher.Publish("user_created", NewUserCreatedMessage(createUserMessage.User))
			if err != nil {
				return err
			}

			time.Sleep(1 * time.Second)
			return nil
		})

		if err != nil && undine.IsDuplicationError(err) {
			fmt.Println("user already created, skipping message", createUserMessage.User)
			msg.Ack()
			continue
		}

		if err != nil {
			fmt.Println("error creating user", err)
			msg.Nack()
			continue
		}

		msg.Ack()
	}
}

func HandleUserCreatedMessage(userCreatedTopic <-chan *message.Message, client *ent.Client) {
	for msg := range userCreatedTopic {
		var userCreatedMessage UserCreatedMessage
		if err := json.Unmarshal(msg.Payload, &userCreatedMessage); err != nil {
			panic(err)
		}

		err := client.WithTx(msg.Context(), func(ctx context.Context) error {
			tx := ent.TxFromContext(ctx)
			deduplicator := tx.Deduplicator()

			messageID := uuid.MustParse(msg.UUID)
			err := deduplicator.Deduplicate(ctx, "user_created", messageID)
			if err != nil {
				return err
			}

			fmt.Println("user created", userCreatedMessage.User)
			time.Sleep(1 * time.Second)
			return nil
		})

		if err != nil && undine.IsDuplicationError(err) {
			fmt.Println("user created already handled, skipping message", userCreatedMessage.User)
			msg.Ack()
			continue
		}

		if err != nil {
			fmt.Println("error creating user", err)
			msg.Nack()
			continue
		}

		msg.Ack()
	}
}
