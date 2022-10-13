package main

import (
	"CNSA_CONCERT2019/controller"
	"html/template"
	"io"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Template is a custom html/template renderer for Echo framework
type Template struct {
	templates *template.Template
}

// Render renders a template document
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.New("").Delims("[[", "]]").ParseFiles(
			"view/login.html", "view/changePassword.html", "view/changePasswordFirst.html", "view/index.html", "view/reserveCN.html", "view/reserveSA.html", "view/checkReserve.html", "view/aboutUs.html", "view/success.html",
		)),
	}

	e := echo.New()
	//models.Insert("172188", "강시현")
	/*
		for i := 181001; i <= 181157; i++ {
			models.ChangePasswordAdmin(strconv.Itoa(i), strconv.Itoa(i))
		}
		for i := 182001; i <= 182213; i++ {
			models.ChangePasswordAdmin(strconv.Itoa(i), strconv.Itoa(i))
		}
	*/

	// Set middlewares
	// Logger: loging all request and responses
	// Recover: Recover main thread if it fails
	e.Use(middleware.Logger(), middleware.Recover())

	// Session 설정
	store := session.NewCookieStore([]byte("secret"))
	e.Use(session.Sessions("CASESSION", store))

	// Set template renderer
	// We uses standard golang template
	e.Renderer = t

	// Set static serve files
	e.Static("/assets", "static")

	// Handle requests
	// Filter by path
	// ================ 학생 페이지 ================
	// 로그인 페이지
	e.GET("/login", controller.Login)
	e.POST("/login", controller.LoginPost)

	// 로그아웃
	e.GET("/logout", controller.Logout)

	// 비밀번호 변경 페이지
	e.GET("/changePasswordFirst", controller.ChangePasswordFirst, controller.AuthAPI)
	e.GET("/changePassword", controller.ChangePassword, controller.AuthAPI)

	// 메인 페이지
	e.GET("/", controller.Index, controller.AuthAPI)
	// 예매 페이지
	e.GET("/reserve", controller.Reserve, controller.AuthAPI)
	e.GET("/checkReserve", controller.CheckReserve, controller.AuthAPI)
	e.GET("/success", controller.Success, controller.AuthAPI)

	// 어바웃 어스
	e.GET("/aboutUs", controller.AboutUS, controller.AuthAPI)

	// ================ 학생 API ================
	// 신청한 좌석 가져오기
	e.GET("/api/getApplys", controller.GetApplysAPI, controller.AuthAPI)
	// 좌석 신청하기
	e.POST("/api/apply", controller.AddApplyAPI, controller.AuthAPI)
	// 비밀번호 변경
	e.POST("/api/changePasswordFirst", controller.ChangePasswordFirstAPI, controller.AuthAPI)
	e.POST("/api/changePassword", controller.ChangePasswordAPI, controller.AuthAPI)

	// 학생 입력
	e.POST("/api/insert", controller.InsertAPI)
	// 미신청 학생입력
	e.POST("/api/insertNope", controller.InsertNopeAPI)

	// Start web server
	e.Start(":80")
}
