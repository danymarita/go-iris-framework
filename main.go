package main

import (
	"go_mongo_iris/src/config"
	"go_mongo_iris/src/modules/profile/controller"
	"go_mongo_iris/src/modules/profile/repository"
	"go_mongo_iris/src/modules/profile/usecase"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"

	// Kita berikan alias karena nama packagenya sama dengan profile
	storyController "go_mongo_iris/src/modules/story/controller"
	storyRepo "go_mongo_iris/src/modules/story/repository"
	storyUC "go_mongo_iris/src/modules/story/usecase"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	views := iris.HTML("./web/views", ".html").Layout("layout.html").Reload(true)

	app.RegisterView(views)
	app.StaticWeb("/public", "./web/public")

	db, err := config.GetMongoDb()
	if err != nil {
		os.Exit(2)
	}

	// Dependency Injection di Go
	// Inisialisasi repository dengan argument db connection dan nama collection mongodb
	// Story
	storyRepository := storyRepo.NewStoryRepositoryMongo(db, "stories")
	storyUsecase := storyUC.NewStoryUsecase(storyRepository)
	// Profile
	profileRepository := repository.NewProfileRepositoryMongo(db, "profiles")
	profileUsecase := usecase.NewProfileUsecase(profileRepository, storyRepository)

	// declare session manager
	sessionManager := sessions.New(sessions.Config{
		Cookie:  "cookieprofile",
		Expires: time.Minute * 10,
	})

	// Profile
	// Route grouping
	profile := mvc.New(app.Party("/profile"))
	profile.Register(profileUsecase, sessionManager.Start)
	// Handle profile controller
	profile.Handle(new(controller.ProfileController))

	// Story
	story := mvc.New(app)
	story.Register(storyUsecase)
	story.Handle(new(storyController.StoryController))

	app.Run(iris.Addr(":3000"))
}
