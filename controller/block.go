package controller

import (
	"encoding/hex"
	"log"
	"net/http"
	"satoblock/lib/utils"
	"satoblock/model"
	"satoblock/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBlocksByHeightRange(ctx *gin.Context) {
	log.Printf("GetBlocksByHeightRange enter")

	// check height
	blkStartHeightString := ctx.Param("start")
	blkStartHeight, err := strconv.Atoi(blkStartHeightString)
	if err != nil || blkStartHeight < 0 {
		log.Printf("blk start height invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "blk start height invalid"})
		return
	}
	blkEndHeightString := ctx.Param("end")
	blkEndHeight, err := strconv.Atoi(blkEndHeightString)
	if err != nil || blkEndHeight < 0 {
		log.Printf("blk end height invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "blk end height invalid"})
		return
	}

	if blkEndHeight <= blkStartHeight || (blkEndHeight-blkStartHeight > 1000) {
		log.Printf("blk end height invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "blk end height invalid"})
		return
	}

	result, err := service.GetBlocksByHeightRange(blkStartHeight, blkEndHeight)
	if err != nil {
		log.Printf("get blocks failed: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "get blocks failed"})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "ok",
		Data: result,
	})
}

func GetBlockByHeight(ctx *gin.Context) {
	log.Printf("GetBlockByHeight enter")

	// check height
	blkHeightString := ctx.Param("height")
	blkHeight, err := strconv.Atoi(blkHeightString)
	if err != nil || blkHeight < 0 {
		log.Printf("blk height invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "blk height invalid"})
		return
	}

	result, err := service.GetBlockByHeight(blkHeight)
	if err != nil {
		log.Printf("get block failed: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "get block failed"})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "ok",
		Data: result,
	})
}

func GetBlockById(ctx *gin.Context) {
	log.Printf("GetBlockById enter")

	blkIdHex := ctx.Param("blkid")
	// check
	blkIdReverse, err := hex.DecodeString(blkIdHex)
	if err != nil {
		log.Printf("blkid invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "blkid invalid"})
		return
	}
	blkId := utils.ReverseBytes(blkIdReverse)

	result, err := service.GetBlockById(hex.EncodeToString(blkId))
	if err != nil {
		log.Printf("get block failed: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "get block failed"})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "ok",
		Data: result,
	})
}
