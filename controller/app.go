package controller

import (
	"CNSA_CONCERT2019/models"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// AuthAPI 로그인 인증 middleware
func AuthAPI(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 로그인이 되어 있지 않으면 login page로 redirect
		session := session.Default(c)
		if session.Get("studentNumber") == nil {
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}

		return next(c)
	}
}

// Login : Login Page
func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

// LoginPost : Check a Login Data
func LoginPost(c echo.Context) error {
	isSuccessed, name, permission := models.Login(c.FormValue("studentNumber"), c.FormValue("password"))

	// Login 성공 시
	if isSuccessed {
		// Session에 학번 저장
		session := session.Default(c)
		session.Set("studentNumber", c.FormValue("studentNumber"))
		session.Set("name", name)
		cl, rw, co := models.GetApplyMine(c.FormValue("studentNumber"))
		session.Set("class", cl)
		session.Set("row", rw)
		session.Set("col", co)
		session.Save()

		if c.FormValue("studentNumber") == c.FormValue("password") {
			return c.Redirect(http.StatusMovedPermanently, "/changePasswordFirst")
		}

		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	// Login 실패 시
	return c.Redirect(http.StatusMovedPermanently, "/login?error="+permission)
}

// Logout : 로그아웃 - 세션 초기화
func Logout(c echo.Context) error {
	// Session 초기화
	session := session.Default(c)
	session.Clear()
	session.Save()

	// 로그인 페이지로 빠이빠이
	return c.Redirect(http.StatusMovedPermanently, "/login")
}

// ChangePasswordFirst : 첫 비번 변경 페이지
func ChangePasswordFirst(c echo.Context) error {
	return c.Render(http.StatusOK, "change_password_first", nil)
}

// ChangePassword : 비번 변경 페이지
func ChangePassword(c echo.Context) error {
	return c.Render(http.StatusOK, "change_password", nil)
}

// Index : Main Page
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

// Reserve : 예매 Page
func Reserve(c echo.Context) error {
	session := session.Default(c)
	studentNumber := session.Get("studentNumber").(string)
	grade := c.QueryParam("grade")

	if grade == "1" && studentNumber[1] == '9' {
		// 1학년인 경우
		return c.Render(http.StatusOK, "reserve_cn", map[string]interface{}{
			"name":          session.Get("name").(string),
			"studentNumber": session.Get("studentNumber").(string),
		})
	} else if grade == "2" && (studentNumber[1] == '8' || studentNumber[1] == '7') {
		// 2학년인 경우
		return c.Render(http.StatusOK, "reserve_sa", map[string]interface{}{
			"name":          session.Get("name").(string),
			"studentNumber": session.Get("studentNumber").(string),
		})
	}
	if studentNumber[1] == '9' {
		return c.Redirect(http.StatusMovedPermanently, "/reserve?grade=1")
	}
	return c.Redirect(http.StatusMovedPermanently, "/reserve?grade=2")
}

// CheckReserve : 예매 확인 Page
func CheckReserve(c echo.Context) error {
	session := session.Default(c)
	cla := session.Get("class").(string)
	row := session.Get("row").(string)
	col := session.Get("col").(string)
	if (cla == c.QueryParam("class") && row == c.QueryParam("row")) && col == c.QueryParam("col") {
		return c.Render(http.StatusOK, "check_reserve", map[string]interface{}{
			"name":          session.Get("name").(string),
			"studentNumber": session.Get("studentNumber").(string),
		})
	}
	return c.Redirect(http.StatusMovedPermanently, "/checkReserve?class="+cla+"&row="+row+"&col="+col)
}

// Success : 완료
func Success(c echo.Context) error {
	return c.Render(http.StatusOK, "success", nil)
}

// AboutUS : 어바웃 어스 Page
func AboutUS(c echo.Context) error {
	return c.Render(http.StatusOK, "about_us", nil)
}

// =========== API ===========

// GetApplysAPI 좌석 신청 가져오기
func GetApplysAPI(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetApplysByClass(c.QueryParam("class")))
}

// AddApplyAPI 좌석 신청하기
func AddApplyAPI(c echo.Context) error {
	session := session.Default(c)
	if c.FormValue("class") == "" {
		return c.Redirect(http.StatusMovedPermanently, "/success?success=false")
	}

	err := models.AddApply(session.Get("studentNumber").(string), session.Get("name").(string), c.FormValue("class"), c.FormValue("row"), c.FormValue("col"))

	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/success?success=false")
	}
	session.Set("class", c.FormValue("class"))
	session.Set("row", c.FormValue("row"))
	session.Set("col", c.FormValue("col"))
	session.Save()
	return c.Redirect(http.StatusMovedPermanently, "/success?success=true")
}

// ChangePasswordAPI 비번 변경
func ChangePasswordAPI(c echo.Context) error {
	session := session.Default(c)

	isSuccessed, _, _ := models.Login(session.Get("studentNumber").(string), c.FormValue("password"))

	if !isSuccessed {
		return c.Redirect(http.StatusMovedPermanently, "/changePassword?error=password")
	}
	if c.FormValue("newPassword") == "" {
		return c.Redirect(http.StatusMovedPermanently, "/changePassword?error=blank")
	}

	if c.FormValue("newPassword") != c.FormValue("newPasswordCheck") {
		return c.Redirect(http.StatusMovedPermanently, "/changePassword?error=check")
	}

	models.ChangePassword(session.Get("studentNumber").(string), c.FormValue("newPassword"))
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// ChangePasswordFirstAPI 비번 변경
func ChangePasswordFirstAPI(c echo.Context) error {
	session := session.Default(c)
	if c.FormValue("newPassword") == "" {
		return c.Redirect(http.StatusMovedPermanently, "/changePasswordFirst?error=blank")
	}
	if c.FormValue("newPassword") != c.FormValue("newPasswordCheck") {
		return c.Redirect(http.StatusMovedPermanently, "/changePasswordFirst?error=check")
	}
	models.ChangePassword(session.Get("studentNumber").(string), c.FormValue("newPassword"))
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// InsertAPI 학생 정보 입력
func InsertAPI(c echo.Context) error {
	models.Insert(c.FormValue("studentNumber"), c.FormValue("name"))

	return c.String(http.StatusOK, "hello")
}

// InsertNopeAPI 미신청 정보 입력
func InsertNopeAPI(c echo.Context) error {
	models.MakeNope(c.FormValue("studentNumber"), c.FormValue("name"))

	return c.String(http.StatusOK, "success")
}
