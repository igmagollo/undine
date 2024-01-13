# undine

An opinionated library that combines Ent and Watermill into a set of powerful utilities for transactionally handle events.

[Ent](https://entgo.io/) is a entity library that generates strongly typed query builders based on the schema. It is simple and its benefits easily overcome the fact that you need to learn how it works instead of just writing raw SQL. It also offers extensions capabilities were you can leverage the schema and generate your own code, e.g. Proto files, OpenAPI specs or GraphQL schemas matching your database schema.

[Watermill](https://watermill.io/) is a library for building event-driven applications. It offers Publisher and Subscriber interfaces and implementations of them for a lot of [message brokers](https://watermill.io/pubsubs/). It has a powerfull and flexible router and offers plenty of other features that makes our lives a lot easier.

This package is meant to combine the watermill with Ent in a way that we can use the Outbox pattern (forwarder component) and a deduplication easily

# What this library offers

- [x] Client exporting WithTx utility function
- [x] Outbox pattern (with watermill forwarder over Ent)
- [x] Deduplication
- [x] All glue code needed generated inside Ent using this Extension

# How to use

First, add the extension to the entc command:

```go
package main

import (
	"log"

	undine "github.com/igmagollo/undine/pkg/v1"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{}, entc.Extensions(
		undine.NewExtension("./schema"),
	))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

```

Now we need to initialize the database passing some dependencies as follows:
```go
client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable",
		ent.DeduplicatorSchemaAdapter(&undine.DeduplicatorPostgresSchemaAdapter{}), // Deduplicator sql adapter
		ent.WatermillLogger(logger),
		ent.Publisher(pubsub), // outside publisher so forwarder can forward messages
		ent.OutboxOffsetsAdapter(&sql.DefaultPostgreSQLOffsetsAdapter{}), // outbox sql adapter
		ent.OutboxSchemaAdapter(&sql.DefaultPostgreSQLSchema{}), // outbox sql adapter
)
```

Then everything that we need is inside `ent.Client` and `ent.Tx` structs:
```go
forwarder := client.Forwarder(consumerGroup)

go func() {
  if err := forwarder.Run(context.Background()); err != nil {
    panic(err)
  }
}()

...

err := client.WithTx(ctx, func(ctx context.Context) error {
      tx := ent.TxFromContext(ctx)
      deduplicator := tx.Deduplicator()
      outboxPublisher, err := tx.OutboxPublisher()

      ...

      err = deduplicator.Deduplicate(ctx, topic, msgID)

      ...

      err = outboxPublisher.Publish(topic, msg)

      ...
})

if undine.IsDuplicationError(err) {
  ...
}
```
