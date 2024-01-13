package undine

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

type Extension struct {
	entc.DefaultExtension
}

func (Extension) Name() string {
	return "undine"
}

func (e Extension) Hooks() []gen.Hook {
	return []gen.Hook{}
}

func (Extension) Options() []entc.Option {
	return []entc.Option{
		entc.FeatureNames("sql/execquery"),
		entc.Dependency(
			entc.DependencyName("WatermillLogger"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "watermill.LoggerAdapter",
				PkgPath: "github.com/ThreeDotsLabs/watermill",
			}),
		),
		entc.Dependency(
			entc.DependencyName("Publisher"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "message.Publisher",
				PkgPath: "github.com/ThreeDotsLabs/watermill/message",
			}),
		),
		entc.Dependency(
			entc.DependencyName("OutboxSchemaAdapter"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "wsql.SchemaAdapter",
				PkgPath: "github.com/ThreeDotsLabs/watermill-sql/pkg/sql",
				PkgName: "wsql",
			}),
		),
		entc.Dependency(
			entc.DependencyName("OutboxOffsetsAdapter"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "wsql.OffsetsAdapter",
				PkgPath: "github.com/ThreeDotsLabs/watermill-sql/pkg/sql",
				PkgName: "wsql",
			}),
		),
		entc.Dependency(
			entc.DependencyName("DeduplicatorSchemaAdapter"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "undine.DeduplicatorSchemaAdapter",
				PkgPath: "github.com/igmagollo/undine/pkg/v1",
				PkgName: "undine",
			}),
		),
	}
}

func (Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("DB").ParseFiles("../../templates/db.go.tmpl")),
		gen.MustParse(gen.NewTemplate("WithTx").ParseFiles("../../templates/with_tx.go.tmpl")),
		gen.MustParse(gen.NewTemplate("UndineOutbox").ParseFiles("../../templates/undine_outbox.go.tmpl")),
		gen.MustParse(gen.NewTemplate("UndineDeduplicator").ParseFiles("../../templates/undine_deduplicator.go.tmpl")),
		gen.MustParse(gen.NewTemplate("WatermillContextExecutor").ParseFiles("../../templates/watermill_context_executor.go.tmpl")),
	}
}
