package category

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/delivery/domain/category"
	"github.com/ninosistemas10/delivery/infrastructure/handler/response"
	"github.com/ninosistemas10/delivery/model"
)

type handler struct {
	useCase  category.UseCase
	response response.API
}

func newHandler(useCase category.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
	m := model.Category{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

func (h handler) Update(c echo.Context) error {
	m := model.Category{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	m.ID = ID

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "h.useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

func (h handler) UpdateImage(c echo.Context) error {
	// 🔹 Parsear el ID de la categoría desde la URL
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	// 🔹 Procesar el archivo de imagen
	file, err := c.FormFile("image")
	if err != nil {
		return h.response.Error(c, "No image file provided", err)
	}

	// 🔹 Abrir el archivo
	src, err := file.Open()
	if err != nil {
		return h.response.Error(c, "Unable to open image file", err)
	}
	defer src.Close()

	// 🔹 Definir la ruta de almacenamiento
	uploadDir := "uploads/categorias/"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return h.response.Error(c, "Unable to create directory", err)
	}

	// 🔹 Generar un nombre de archivo único
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, filename)

	// 🔹 Crear el archivo destino
	dst, err := os.Create(filePath)
	if err != nil {
		return h.response.Error(c, "Unable to save image file", err)
	}
	defer dst.Close()

	// 🔹 Guardar la imagen en el servidor
	if _, err := io.Copy(dst, src); err != nil {
		return h.response.Error(c, "Error copying image to destination", err)
	}

	// 🔹 Verificar si la imagen realmente se guardó
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return h.response.Error(c, "Image file was not saved", err)
	}

	// 🔹 Construir la URL de acceso a la imagen
	imageURL := fmt.Sprintf("http://localhost:8081/categorias/%s", filename)

	// 🔹 Llamar al caso de uso para actualizar la imagen en la base de datos
	if err := h.useCase.UpdateImage(ID, imageURL); err != nil {
		return h.response.Error(c, "useCase.UpdateImage failed", err)
	}

	// 🔹 Retornar respuesta exitosa con la URL correcta
	return c.JSON(h.response.OK(map[string]string{
		"message": "Image updated successfully",
		"images":  imageURL,
	}))
}

func (h handler) Delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	err = h.useCase.Delete(ID)
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.response.Error(c, "uuid.Parse()", err)
	}

	categoryData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.response.Error(c, "useCase.GetBYID", err)
	}

	return c.JSON(h.response.OK(categoryData))
}

func (h handler) GetAll(c echo.Context) error {
	categorys, err := h.useCase.GetAll()
	if err != nil {
		return h.response.Error(c, "useCase.GetAll", err)
	}

	return c.JSON(h.response.OK(categorys))
}
