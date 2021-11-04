package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"source-server/util"

	"github.com/gin-gonic/gin"
)

func GetParams(ctx *gin.Context) {
	fmt.Print("hello world")
	util.LoadConfig("../conf/user.conf")
	ctx.String(200, "success hhh")
}

func Upload(ctx *gin.Context) {
	util.LoadConfig("../conf/user.conf")
	// 单文件
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	// 上传文件至指定目录
	filename := filepath.Base(file.Filename)
	dst := path.Join("./upload", filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
