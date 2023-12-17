package undine

import (
	"fmt"

	"entgo.io/contrib/schemast"
	"entgo.io/ent"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

func generateOutbox(ctx *schemast.Context) error {
	mutations := []schemast.Mutator{
		&schemast.UpsertSchema{
			Name: "Outbox",
			Fields: []ent.Field{
				field.UUID("id", uuid.New()).Immutable(), // schemast does not support Default(uuid.New)
				field.String("topic").Immutable(),        // schemast does not support NotEmpty
				field.Bytes("payload").Immutable(),       // schemast does not support NotEmpty
				field.JSON("headers", map[string]string{}).Optional(),
				field.Time("created_at"), // schemast does not support Default(time.Now)
			},
			Indexes: []ent.Index{
				index.Fields("created_at", "topic"),
			},
		},
	}

	if err := schemast.Mutate(ctx, mutations...); err != nil {
		return err
	}

	return nil
}

func generateProcessedEvent(ctx *schemast.Context) error {
	mutations := []schemast.Mutator{
		&schemast.UpsertSchema{
			Name: "ProcessedMessage",
			Fields: []ent.Field{
				field.UUID("id", uuid.New()).Immutable(),         // schemast does not support Default(uuid.New)
				field.UUID("message_id", uuid.New()).Immutable(), // schemast does not support NotEmpty
				field.String("message_topic").Immutable(),        // schemast does not support NotEmpty
				field.Time("created_at"),                         // schemast does not support Default(time.Now)
			},
			Indexes: []ent.Index{
				index.Fields("message_id", "message_topic").Unique(),
			},
		},
	}

	if err := schemast.Mutate(ctx, mutations...); err != nil {
		return err
	}

	return nil
}

func NewExtension(schemaPath string) Extension {
	return Extension{
		schemaPath: schemaPath,
	}
}

type Extension struct {
	entc.DefaultExtension

	schemaPath string
}

func (Extension) Name() string {
	return "undine"
}

func (e Extension) Hooks() []gen.Hook {
	return []gen.Hook{
		func(next gen.Generator) gen.Generator {
			return gen.GenerateFunc(func(g *gen.Graph) error {
				fmt.Println("Generating undine hooks", e.schemaPath)
				ctx, err := schemast.Load(e.schemaPath)
				if err != nil {
					return fmt.Errorf("loading schema: %w", err)
				}

				if err := generateOutbox(ctx); err != nil {
					return fmt.Errorf("generating outbox: %w", err)
				}

				if err := generateProcessedEvent(ctx); err != nil {
					return fmt.Errorf("generating processed event: %w", err)
				}

				if err := ctx.Print(e.schemaPath); err != nil {
					return fmt.Errorf("printing schema: %w", err)
				}

				newGraph, err := entc.LoadGraph(e.schemaPath, g.Config)
				if err != nil {
					return fmt.Errorf("loading graph: %w", err)
				}

				return next.Generate(newGraph)
			})
		},
	}
}

func (Extension) Options() []entc.Option {
	return []entc.Option{}
}

func (Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("WithTx").ParseFiles("../../templates/with_tx.go.tmpl")),
		gen.MustParse(gen.NewTemplate("OutboxStorer").ParseFiles("../../templates/outbox_storer.go.tmpl")),
		gen.MustParse(gen.NewTemplate("OutboxWatermill").ParseFiles("../../templates/outbox_watermill.go.tmpl")),
		gen.MustParse(gen.NewTemplate("ProcessedmessageWatermill").ParseFiles("../../templates/processedmessage_watermill.go.tmpl")),
		gen.MustParse(gen.NewTemplate("Watermill").ParseFiles("../../templates/watermill.go.tmpl")),
	}
}
