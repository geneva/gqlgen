package resolvergen

import (
	_ "embed"
	"errors"
	"fmt"
	"go/ast"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/geneva/gqlgen/codegen"
	"github.com/geneva/gqlgen/codegen/config"
	"github.com/geneva/gqlgen/codegen/templates"
	"github.com/geneva/gqlgen/graphql"
	"github.com/geneva/gqlgen/internal/rewrite"
	"github.com/geneva/gqlgen/plugin"
)

//go:embed resolver.gotpl
var resolverTemplate string

func New() plugin.Plugin {
	return &Plugin{}
}

type Plugin struct{}

var _ plugin.CodeGenerator = &Plugin{}

func (m *Plugin) Name() string {
	return "resolvergen"
}

func (m *Plugin) GenerateCode(data *codegen.Data) error {
	if !data.Config.Resolver.IsDefined() {
		return nil
	}

	switch data.Config.Resolver.Layout {
	case config.LayoutSingleFile:
		return m.generateSingleFile(data)
	case config.LayoutFollowSchema:
		return m.generatePerSchema(data)
	}

	return nil
}

func (m *Plugin) generateSingleFile(data *codegen.Data) error {
	file := File{}

	if _, err := os.Stat(data.Config.Resolver.Filename); err == nil {
		// file already exists and we do not support updating resolvers with layout = single so just return
		return nil
	}

	for _, o := range data.Objects {
		if o.HasResolvers() {
			file.Objects = append(file.Objects, o)
		}
		for _, f := range o.Fields {
			if !f.IsResolver {
				continue
			}

			resolver := Resolver{o, f, nil, "", `panic("not implemented")`}
			file.Resolvers = append(file.Resolvers, &resolver)
		}
	}

	resolverBuild := &ResolverBuild{
		File:                &file,
		PackageName:         data.Config.Resolver.Package,
		ResolverType:        data.Config.Resolver.Type,
		HasRoot:             true,
		OmitTemplateComment: data.Config.Resolver.OmitTemplateComment,
	}

	newResolverTemplate := resolverTemplate
	if data.Config.Resolver.ResolverTemplate != "" {
		newResolverTemplate = readResolverTemplate(data.Config.Resolver.ResolverTemplate)
	}

	return templates.Render(templates.Options{
		PackageName: data.Config.Resolver.Package,
		FileNotice:  `// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.`,
		Filename:    data.Config.Resolver.Filename,
		Data:        resolverBuild,
		Packages:    data.Config.Packages,
		Template:    newResolverTemplate,
	})
}

func (m *Plugin) generatePerSchema(data *codegen.Data) error {
	rewriter, err := rewrite.New(data.Config.Resolver.Dir())
	if err != nil {
		return err
	}

	files := map[string]*File{}

	objects := make(codegen.Objects, len(data.Objects)+len(data.Inputs))
	copy(objects, data.Objects)
	copy(objects[len(data.Objects):], data.Inputs)

	for _, o := range objects {
		if o.HasResolvers() {
			fn := gqlToResolverName(data.Config.Resolver.Dir(), o.Position.Src.Name, data.Config.Resolver.FilenameTemplate)
			if files[fn] == nil {
				files[fn] = &File{}
			}

			caser := cases.Title(language.English, cases.NoLower)
			rewriter.MarkStructCopied(templates.LcFirst(o.Name) + templates.UcFirst(data.Config.Resolver.Type))
			rewriter.GetMethodBody(data.Config.Resolver.Type, caser.String(o.Name))
			files[fn].Objects = append(files[fn].Objects, o)
		}
		for _, f := range o.Fields {
			if !f.IsResolver {
				continue
			}

			structName := templates.LcFirst(o.Name) + templates.UcFirst(data.Config.Resolver.Type)
			comment := strings.TrimSpace(strings.TrimLeft(rewriter.GetMethodComment(structName, f.GoFieldName), `\`))

			implementation := strings.TrimSpace(rewriter.GetMethodBody(structName, f.GoFieldName))
			if implementation == "" {
				// Check for Implementer Plugin
				var resolver_implementer plugin.ResolverImplementer
				var exists bool
				for _, p := range data.Plugins {
					if p_cast, ok := p.(plugin.ResolverImplementer); ok {
						resolver_implementer = p_cast
						exists = true
						break
					}
				}

				if exists {
					implementation = resolver_implementer.Implement(f)
				} else {
					implementation = fmt.Sprintf("panic(fmt.Errorf(\"not implemented: %v - %v\"))", f.GoFieldName, f.Name)
				}

			}

			resolver := Resolver{o, f, rewriter.GetPrevDecl(structName, f.GoFieldName), comment, implementation}
			fn := gqlToResolverName(data.Config.Resolver.Dir(), f.Position.Src.Name, data.Config.Resolver.FilenameTemplate)
			if files[fn] == nil {
				files[fn] = &File{}
			}

			files[fn].Resolvers = append(files[fn].Resolvers, &resolver)
		}
	}

	for filename, file := range files {
		file.imports = rewriter.ExistingImports(filename)
		file.RemainingSource = rewriter.RemainingSource(filename)
	}
	newResolverTemplate := resolverTemplate
	if data.Config.Resolver.ResolverTemplate != "" {
		newResolverTemplate = readResolverTemplate(data.Config.Resolver.ResolverTemplate)
	}

	for filename, file := range files {
		resolverBuild := &ResolverBuild{
			File:                file,
			PackageName:         data.Config.Resolver.Package,
			ResolverType:        data.Config.Resolver.Type,
			OmitTemplateComment: data.Config.Resolver.OmitTemplateComment,
		}

		var fileNotice strings.Builder
		if !data.Config.OmitGQLGenFileNotice {
			fileNotice.WriteString(`
			// This file will be automatically regenerated based on the schema, any resolver implementations
			// will be copied through when generating and any unknown code will be moved to the end.
			// Code generated by github.com/geneva/gqlgen`,
			)
			if !data.Config.OmitGQLGenVersionInFileNotice {
				fileNotice.WriteString(` version `)
				fileNotice.WriteString(graphql.Version)
			}
		}

		err := templates.Render(templates.Options{
			PackageName: data.Config.Resolver.Package,
			FileNotice:  fileNotice.String(),
			Filename:    filename,
			Data:        resolverBuild,
			Packages:    data.Config.Packages,
			Template:    newResolverTemplate,
		})
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(data.Config.Resolver.Filename); errors.Is(err, fs.ErrNotExist) {
		err := templates.Render(templates.Options{
			PackageName: data.Config.Resolver.Package,
			FileNotice: `
				// This file will not be regenerated automatically.
				//
				// It serves as dependency injection for your app, add any dependencies you require here.`,
			Template: `type {{.}} struct {}`,
			Filename: data.Config.Resolver.Filename,
			Data:     data.Config.Resolver.Type,
			Packages: data.Config.Packages,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type ResolverBuild struct {
	*File
	HasRoot             bool
	PackageName         string
	ResolverType        string
	OmitTemplateComment bool
}

type File struct {
	// These are separated because the type definition of the resolver object may live in a different file from the
	// resolver method implementations, for example when extending a type in a different graphql schema file
	Objects         []*codegen.Object
	Resolvers       []*Resolver
	imports         []rewrite.Import
	RemainingSource string
}

func (f *File) Imports() string {
	for _, imp := range f.imports {
		if imp.Alias == "" {
			_, _ = templates.CurrentImports.Reserve(imp.ImportPath)
		} else {
			_, _ = templates.CurrentImports.Reserve(imp.ImportPath, imp.Alias)
		}
	}
	return ""
}

type Resolver struct {
	Object         *codegen.Object
	Field          *codegen.Field
	PrevDecl       *ast.FuncDecl
	Comment        string
	Implementation string
}

func gqlToResolverName(base string, gqlname, filenameTmpl string) string {
	gqlname = filepath.Base(gqlname)
	ext := filepath.Ext(gqlname)
	if filenameTmpl == "" {
		filenameTmpl = "{name}.resolvers.go"
	}
	filename := strings.ReplaceAll(filenameTmpl, "{name}", strings.TrimSuffix(gqlname, ext))
	return filepath.Join(base, filename)
}

func readResolverTemplate(customResolverTemplate string) string {
	contentBytes, err := os.ReadFile(customResolverTemplate)
	if err != nil {
		panic(err)
	}
	return string(contentBytes)
}
