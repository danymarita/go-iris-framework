package controller

import (
	"fmt"
	"go_mongo_iris/src/modules/profile/usecase"
	"io"
	"time"

	"os"

	"go_mongo_iris/src/modules/profile/model"
	storyModel "go_mongo_iris/src/modules/story/model"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	uuid "github.com/satori/go.uuid"
	"github.com/wuriyanto48/replacer"
)

const ProfileIDKey = "ProfileID"

type ProfileController struct {
	Ctx            iris.Context
	Session        *sessions.Session
	ProfileUsecase usecase.ProfileUsecase
}

// Get Current Profile ID
func (c *ProfileController) getCurrentProfileID() string {
	return c.Session.GetString(ProfileIDKey)
}

// Check Session is not empty
func (c *ProfileController) isProfileLoggedIn() bool {
	return c.getCurrentProfileID() != ""
}

func (c *ProfileController) logout() {
	c.Session.Destroy()
}

// localhost:3000/profile/story GET
func (c *ProfileController) GetStory() mvc.Result {
	if !c.isProfileLoggedIn() {
		return mvc.Response{
			Path: "/profile/login",
		}
	}

	return mvc.View{
		Name: "profile/story.html",
		Data: iris.Map{"Title": "New Story"},
	}
}

// localhost:3000/profile/story POST
func (c *ProfileController) PostStory() mvc.Result {
	if !c.isProfileLoggedIn() {
		return mvc.Response{
			Path: "/profile/login",
		}
	}

	title := c.Ctx.FormValue("title")
	content := c.Ctx.FormValue("content")

	if title == "" || content == "" {
		return mvc.Response{
			Path: "/profile/story",
		}
	}

	id := uuid.NewV4()
	profile, err := c.ProfileUsecase.GetByID(c.getCurrentProfileID())
	if err != nil {
		c.logout()
		// Agar diredirect ke halaman login
		c.GetMe()
	}

	var story storyModel.Story
	story.ID = id.String()
	story.Title = title
	story.Content = content
	story.Profile = profile
	story.CreatedAt = time.Now()
	story.UpdatedAt = time.Now()

	// Tidak menggunakan := karena variable err sudah dideklarasikan sebelumnya
	_, err = c.ProfileUsecase.CreateStory(&story)
	if err != nil {
		return mvc.Response{
			Path: "/profile/story",
		}
	}

	return mvc.Response{
		Path: "/profile/story",
	}
}

// Iris framework secara otomatis akan melakukan mapping GetRegister sebagai url dengan method GET
// localhost:3000/profile/register
func (c *ProfileController) GetRegister() mvc.Result {
	if c.isProfileLoggedIn() {
		c.logout()
	}

	return mvc.View{
		Name: "profile/register.html",
		Data: iris.Map{"Title": "Profile Registration"},
	}
}

func (c *ProfileController) PostRegister() mvc.Result {
	firstName := c.Ctx.FormValue("first_name")
	lastName := c.Ctx.FormValue("last_name")
	email := c.Ctx.FormValue("email")
	password := c.Ctx.FormValue("password")

	if firstName == "" || lastName == "" || email == "" || password == "" {
		return mvc.Response{
			Path: "/profile/register",
		}
	}

	id := uuid.NewV4()
	imageProfile, err := c.uploadImage(c.Ctx, id.String())
	if err != nil {
		return mvc.Response{
			Path: "/profile/register",
		}
	}

	var profile model.Profile
	profile.ID = id.String()
	profile.FirstName = firstName
	profile.LastName = lastName
	profile.Email = email
	profile.Password = password
	profile.ImageProfile = imageProfile
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()

	// Karena variable err sudah ada diatas jadi kita isi saja dengan value baru, tidak perlu dideklarasikan lagi
	_, err = c.ProfileUsecase.SaveProfile(&profile)

	if err != nil {
		return mvc.Response{
			Path: "/profile/register",
		}
	}

	c.Session.Set(ProfileIDKey, profile.ID)
	return mvc.Response{
		Path: "/profile/me",
	}
}

func (c *ProfileController) GetLogin() mvc.Result {
	if c.isProfileLoggedIn() {
		c.logout()
	}

	return mvc.View{
		Name: "profile/login.html",
		Data: iris.Map{"Title": "Login"},
	}
}

func (c *ProfileController) PostLogin() mvc.Result {
	email := c.Ctx.FormValue("email")
	password := c.Ctx.FormValue("password")

	if email == "" || password == "" {
		return mvc.Response{
			Path: "/profile/login",
		}
	}

	profile, err := c.ProfileUsecase.GetByEmail(email)
	if err != nil {
		// fmt.Println(err)
		return mvc.Response{
			Path: "/profile/login",
		}
	}

	if !profile.IsValidPassword(password) {
		return mvc.Response{
			Path: "/profile/login",
		}
	}

	c.Session.Set(ProfileIDKey, profile.ID)

	return mvc.Response{
		Path: "/profile/me",
	}
}

func (c *ProfileController) GetMe() mvc.Result {
	if !c.isProfileLoggedIn() {
		return mvc.Response{
			Path: "/profile/login",
		}
	}

	profile, err := c.ProfileUsecase.GetByID(c.getCurrentProfileID())
	if err != nil {
		c.logout()
		// arahkan ke getMe agar diredirect ke halaman login
		c.GetMe()
	}

	return mvc.View{
		Name: "profile/me.html",
		Data: iris.Map{
			"Title":   "My Profile",
			"Profile": profile,
		},
	}
}

// Fungsi ini akan dimapping ke url localhost:3000/profile/logout
func (c *ProfileController) AnyLogout() {
	if c.isProfileLoggedIn() {
		c.logout()
	}

	c.Ctx.Redirect("/profile/login")
}

// Fungsi upload image
// iris.Context berfungsi untuk mengambil field inputan form
func (c *ProfileController) uploadImage(ctx iris.Context, id string) (string, error) {
	file, info, err := ctx.FormFile("image_profile")
	if err != nil {
		return "", err
	}

	defer file.Close()
	// replacer.Replace digunakan untuk mereplace karakter uncommon di file name
	filename := fmt.Sprintf("%s%s%s", id, "_", replacer.Replace(info.Filename, "_"))
	out, err := os.OpenFile("./web/public/images/profile/"+filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return "", err
	}

	defer out.Close()

	io.Copy(out, file)

	return filename, nil
}
