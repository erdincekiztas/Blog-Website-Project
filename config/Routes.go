package config

import (
	admin "goblog/admin/controllers"
	client "goblog/client/controllers"
	"goblog/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()

	// ? /admin
	r.GET("/admin", middleware.AdminMiddleware(admin.Dashboard{}.Index))

	//?  /admin/yeni-ekle
	r.GET("/admin/yeni-post-ekle", middleware.AdminMiddleware(admin.Dashboard{}.NewItem))
	r.POST("/admin/ekle", middleware.AdminMiddleware(admin.Dashboard{}.Add))
	r.GET("/admin/duzenle/:id", middleware.AdminMiddleware(admin.Dashboard{}.Edit))
	r.POST("/admin/guncelle/:id", middleware.AdminMiddleware(admin.Dashboard{}.Update))
	r.GET("/admin/post-sil/:id", middleware.AdminMiddleware(admin.Dashboard{}.Delete))

	r.GET("/admin/danisanlar/listele", middleware.AdminMiddleware(admin.Users{}.UserClientPage))

	r.GET("/admin/danisan/:userID/grafik", middleware.AdminMiddleware(admin.Measurement{}.Grafik))
	r.POST("/admin/danisan/:userID/grafik/ekle", middleware.AdminMiddleware(admin.Measurement{}.GrafikAdd))

	r.GET("/admin/danisan/:userID/antrenman-programi", middleware.AdminMiddleware(admin.Workout{}.ExercisesPage))
	r.POST("/admin/danisan/:userID/antrenman-programi/ekle", middleware.AdminMiddleware(admin.Workout{}.ExerciseAdd))
	r.GET("/admin/danisan/:userID/antrenman-programi/sil/:id", middleware.AdminMiddleware(admin.Workout{}.ExerciseDelete))

	r.GET("/admin/danisan/:userID/diyet-programi", middleware.AdminMiddleware(admin.Diet{}.DietPage))
	r.POST("/admin/danisan/:userID/diyet-programi/ekle", middleware.AdminMiddleware(admin.Diet{}.DietAdd))
	r.GET("/admin/danisan/:userID/diyet-programi/sil/:id", middleware.AdminMiddleware(admin.Diet{}.DietDelete))

	r.GET("/admin/login", admin.Authorization{}.LoginPage)
	r.POST("/admin/do-login", admin.Authorization{}.Login)
	r.GET("/admin/logout", admin.Authorization{}.Logout)

	r.GET("/diyet-programi", middleware.ClientMiddleware(client.Diet{}.DietPage))
	r.GET("/olcumler", middleware.ClientMiddleware(client.Measurement{}.MeasurementPage))
	r.GET("/antrenman-programi", middleware.ClientMiddleware(client.Workout{}.ExercisesPage))
	r.GET("/anasayfa", middleware.ClientMiddleware(client.Dashboard{}.HomePage))

	r.GET("/login", client.Authorization{}.LoginPage)
	r.POST("/do-login", client.Authorization{}.Login)
	r.GET("/logout", client.Authorization{}.Logout)
	r.GET("/register", client.Authorization{}.RegistrationPage)
	r.POST("/do-register", client.Authorization{}.Registration)

	r.GET("/", client.Blog{}.BlogPage)
	r.GET("/yazilar/:slug", client.Blog{}.Detail)

	//? SERVE FİLES Admin dashboard dosyalarını gösterebilmek için
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
