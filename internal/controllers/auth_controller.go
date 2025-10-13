package controllers

import (
	"net/http"
	"os"
	"time"

	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	log.LogInfo("🔍 Acceso a Index", nil)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"status": "OK",
	})
}

func RegisterForm(c *gin.Context) {
	log.LogInfo("📝 Acceso a formulario de registro", nil)
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegisterUser(c *gin.Context) {
	var input struct {
		Username string `form:"username" binding:"required"`
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		log.LogWarn("⚠️ Datos inválidos en el registro", map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Datos Inválidos",
		})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.LogError("❌ Error al hashear la contraseña", map[string]interface{}{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"error": "Error Interno",
		})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashed),
		Role:     "USER",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		log.LogError("❌ Error al crear usuario en la base de datos", map[string]interface{}{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
			"user":   input.Username,
		})
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"error": "Error al crear usuario",
		})
		return
	}
	log.LogInfo("✅ Usuario registrado correctamente", map[string]interface{}{
		"status": http.StatusSeeOther,
		"user":   user.Username,
	})
	c.Redirect(http.StatusSeeOther, "/login")
}

func LoginForm(c *gin.Context) {
	log.LogInfo("📝 Acceso a formulario de Login", nil)
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		log.LogWarn("⚠️ Datos inválidos en el login", map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Datos inválidos"})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		log.LogWarn("⚠️ Intento de login con usuario no existente", map[string]interface{}{
			"status": http.StatusUnauthorized,
			"email":  input.Email,
			"error":  err.Error(),
		})
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Usuario no encontrado",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.LogWarn("⚠️ Contraseña incorrecta en login", map[string]interface{}{
			"status": http.StatusUnauthorized,
			"email":  user.Email,
			"error":  err.Error(),
		})
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Contraseña Incorrecta",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.LogError("❌ Error al generar JWT", map[string]interface{}{
			"status": http.StatusInternalServerError,
			"email":  user.Email,
			"error":  err.Error(),
		})
	}

	c.SetCookie("jwt", tokenString, 3600*24, "", "", false, true)

	log.LogInfo("✅ Login exitoso", map[string]interface{}{
		"status": http.StatusSeeOther,
		"email":  user.Email,
	})

	c.Redirect(http.StatusSeeOther, "/")
}

func LogoutUser(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	log.LogInfo("👋 Usuario deslogueado correctamente", nil)
	c.Redirect(http.StatusSeeOther, "/login")
}

func Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		log.LogError("❌ Intento no autorizado de acceso a /profile", map[string]interface{}{
			"status": http.StatusUnauthorized,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Usuario no autenticado",
		})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.LogWarn("⚠️ Usuario no encontrado", map[string]interface{}{
				"status": http.StatusNotFound,
				"id":     userID,
				"error":  err.Error(),
			})
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Usuario no encontrado",
			})
		} else {
			log.LogError("❌ Error al buscar usuario en Profile()", map[string]interface{}{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Error interno del servidor",
			})
		}
		return
	}

	log.LogInfo("✅ Perfil accedido por usuario", map[string]interface{}{
		"id":     userID,
		"status": http.StatusOK,
	})

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"Name":  user.Username,
		"Email": user.Email,
	})
}
