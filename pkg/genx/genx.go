package genx

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type GenX struct{}

func (f *GenX) templates(pkg string) map[string]string {
	return map[string]string{
		"datasource.go":           f.dataSourceTemplate(pkg),
		"handler.go":              f.handlerTemplate(pkg),
		"provider.go":             f.providerTemplate(pkg),
		"repository.go":           f.repositoryTemplate(pkg),
		"router.go":               f.routerTemplate(pkg),
		"usecase.go":              f.useCaseTemplate(pkg),
		fmt.Sprintf("%s.go", pkg): f.modelTemplate(pkg),
	}
}

func (f *GenX) ensureDir(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func (f *GenX) Process(feature string) {
	log.Println("--> START")
	for filename := range f.templates(feature) {
		f.Generate(feature, filename)
	}
	log.Println("<-- END")
}

func (f *GenX) Generate(feature string, filename string) {
	template := f.getTemplate(feature, filename)
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	currentDir = fmt.Sprintf("%s/%s", currentDir, feature)

	if f.ensureDir(currentDir) != nil {
		log.Fatal("Create directory error")
	}

	target := fmt.Sprintf("%s/%s", currentDir, filename)

	if err := ioutil.WriteFile(target, []byte(template), 0755); err != nil {
		log.Println("Generate file error", err)
	} else {
		log.Println(fmt.Sprintf("Generate file %s success", filename))
	}
}

func (f *GenX) dataSourceTemplate(pkg string) string {
	return fmt.Sprintf(`package %s

type DataSource interface {
}

type dataSource struct {
	DbSource database.DataSource
}

func NewDataSource(dbSource database.DataSource) DataSource {
	return &dataSource{
		DbSource: dbSource,
	}
}`, pkg)
}

func (f *GenX) handlerTemplate(pkg string) string {
	return fmt.Sprintf(`package %s

type Handler interface {
}

type handler struct {
	Uc UseCase
}

func NewHandler(uc UseCase) Handler {
	return &handler{
		Uc: uc,
	}
}`, pkg)
}

func (f *GenX) providerTemplate(pkg string) string {
	return fmt.Sprintf(`package %s

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewDataSource,
	NewRepository,
	NewUseCase,
	NewHandler,
	NewRouter,
)`, pkg)
}

func (f *GenX) repositoryTemplate(pkg string) string {
	return fmt.Sprintf(`package %s

type Repository interface {
}

type repository struct {
	Ds DataSource
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}`, pkg)
}

func (f *GenX) routerTemplate(pkg string) string {
	return fmt.Sprintf(`package %s

import "github.com/labstack/echo"

type Router interface {
	Initial(e *echo.Echo)
}

type router struct {
	Handle Handler
}

func (r *router) Initial(e *echo.Echo) {

}

func NewRouter(handle Handler) Router {
	return &router{Handle: handle}
}`, pkg)
}

func (f *GenX) useCaseTemplate(pkg string) string {
	return fmt.Sprintf(`package %s

type UseCase interface {
}

type useCase struct {
	Repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}`, pkg)
}

func (f *GenX) modelTemplate(pkg string) string {
	model := f.modelName(pkg)

	return fmt.Sprintf(`package %s

type %s struct  {

}`, pkg, model)
}

func (f *GenX) getTemplate(pkg string, filename string) string {
	return f.templates(pkg)[filename]
}

func (f *GenX) modelName(feature string) string {
	first := strings.ToUpper(feature[:1])
	last := feature[1:]
	modelName := fmt.Sprintf("%s%s", first, last)
	return modelName
}

func New() GenX {
	return GenX{}
}
