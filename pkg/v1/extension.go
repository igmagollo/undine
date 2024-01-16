package undine

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"

	_ "embed"
)

var (
	//go:embed templates/db.go.tmpl
	dbTemplate string

	//go:embed templates/with_tx.go.tmpl
	withTxTemplate string

	//go:embed templates/undine_outbox.go.tmpl
	undineOutboxTemplate string

	//go:embed templates/undine_deduplicator.go.tmpl
	undineDeduplicatorTemplate string

	//go:embed templates/watermill_context_executor.go.tmpl
	watermillContextExecutorTemplate string
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
		gen.MustParse(gen.NewTemplate("DB").Parse(dbTemplate)),
		gen.MustParse(gen.NewTemplate("WithTx").Parse(withTxTemplate)),
		gen.MustParse(gen.NewTemplate("UndineOutbox").Parse(undineOutboxTemplate)),
		gen.MustParse(gen.NewTemplate("UndineDeduplicator").Parse(undineDeduplicatorTemplate)),
		gen.MustParse(gen.NewTemplate("WatermillContextExecutor").Parse(watermillContextExecutorTemplate)),
	}
}
