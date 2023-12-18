# undine

An opinionated library that combines Ent and Watermill into a set of powerful utilities for transactionally handle events.

[Ent](https://entgo.io/) is a entity library that generates strongly typed query builders based on the schema. It is simple and its benefits easily overcome the fact that you need to learn how it works instead of just writing raw SQL. It also offers extensions capabilities were you can leverage the schema and generate your own code, e.g. Proto files, OpenAPI specs or GraphQL schemas matching your database schema.

[Watermill](https://watermill.io/) is a library for building event-driven applications. It offers Publisher and Subscriber interfaces and implementations of them for a lot of [message brokers](https://watermill.io/pubsubs/). It has a powerfull and flexible router and offers plenty of other features that makes our lives a lot easier.

The thing is, when we are working with event-driven systems, it is very common to have to deal with at-least-once guarantee and make the consumers
idempotent. It has known solutions, but they are not so trivial to implement and they adds a lot of boilerplate and complexity to our code. We can find outbox pattern implementation packages to be used with SQL, but they don't fit with Ent due to its strong schema coupling. The same can be applied to every pattern that uses the database to add more guarantees to the delivery. So, this library aims to combine the best of both worlds, creating reusable patterns that leverages the Ent's code generation to couple it with the schema and using the watermill as the event-driven lib.

# What this library offers

- [x] Outbox pattern implementation
  - [x] Outbox table creation
  - [x] Outbox event storing
  - [x] Outbox event publisher (just the storer implementing watermill's publisher interface)
  - [ ] Outbox event relay
- [x] Idempotency through deduplication
  - [x] Deduplication table creation
  - [x] Deduplication event storing
- [x] Watermill Router wrapper to transactionally handle events
  - [x] AddTransactionalHandler (just a wrapper to the router's AddHandler)
  - [x] AddPublisherDecorators (wraps the AddPublisherDecorators to decorate outbox publishers too)
  - [ ] Make Middlewares to be applied to the inner handler instead of the wrapped one in the AddTransactionalHandler
- [ ] Inbox pattern implementation
  - [ ] Inbox table creation
  - [ ] Inbox event storing
  - [ ] Inbox event relay
- [ ] OpenTelemetry
  - [ ] Metrics
  - [ ] Tracing

# How to use

Add the extension to the entc command:

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

Now, the entc will generate all the tools that undine provides in your ent package.
