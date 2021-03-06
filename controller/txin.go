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

func GetTxInputsByTxId(ctx *gin.Context) {
	log.Printf("GetTxInputsByTxId enter")

	txIdHex := ctx.Param("txid")
	// check
	txIdReverse, err := hex.DecodeString(txIdHex)
	if err != nil {
		log.Printf("txid invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "txid invalid"})
		return
	}
	txId := utils.ReverseBytes(txIdReverse)

	result, err := service.GetTxInputsByTxId(hex.EncodeToString(txId))
	if err != nil {
		log.Printf("get block failed: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "get txin failed"})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "ok",
		Data: result,
	})
}

func GetTxInputsByTxIdInsideHeight(ctx *gin.Context) {
	log.Printf("GetTxInputsByTxIdInsideHeight enter")

	blkHeightString := ctx.Param("height")
	blkHeight, err := strconv.Atoi(blkHeightString)
	if err != nil {
		log.Printf("height invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "height invalid"})
		return
	}

	txIdHex := ctx.Param("txid")
	// check
	txIdReverse, err := hex.DecodeString(txIdHex)
	if err != nil {
		log.Printf("txid invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "txid invalid"})
		return
	}
	txId := utils.ReverseBytes(txIdReverse)

	result, err := service.GetTxInputsByTxIdInsideHeight(blkHeight, hex.EncodeToString(txId))
	if err != nil {
		log.Printf("get block failed: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "get txin failed"})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "ok",
		Data: result,
	})
}

func GetTxInputByTxIdAndIdx(ctx *gin.Context) {
	log.Printf("GetTxInputByTxIdAndIdx enter")

	// check tx
	txIdHex := ctx.Param("txid")
	txIdReverse, err := hex.DecodeString(txIdHex)
	if err != nil {
		log.Printf("txid invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "txid invalid"})
		return
	}
	txId := utils.ReverseBytes(txIdReverse)

	// check index
	txIndexString := ctx.Param("index")
	txIndex, err := strconv.Atoi(txIndexString)
	if err != nil || txIndex < 0 {
		log.Printf("txindex invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "txindex invalid"})
		return
	}

	result, err := service.GetTxInputByTxIdAndIdx(hex.EncodeToString(txId), txIndex)
	if err != nil {
		log.Printf("get block failed: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "get txin failed"})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "ok",
		Data: result,
	})
}

func GetTxInputByTxIdAndIdxInsideHeight(ctx *gin.Context) {
	log.Printf("GetTxInputByTxIdAndIdxInsideHeight enter")

	blkHeightString := ctx.Param("height")
	blkHeight, err := strconv.Atoi(blkHeightString)
	if err != nil {
		log.Printf("height invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "height invalid"})
		return
	}

	// check tx
	txIdHex := ctx.Param("txid")
	txIdReverse, err := hex.DecodeString(txIdHex)
	if err != nil {
		log.Printf("txid invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "txid invalid"})
		return
	}
	txId := utils.ReverseBytes(txIdReverse)

	// check index
	txIndexString := ctx.Param("index")
	txIndex, err := strconv.Atoi(txIndexString)
	if err != nil || txIndex < 0 {
		log.Printf("txindex invalid: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "txindex invalid"})
		return
	}

	result, err := service.GetTxInputByTxIdAndIdxInsideHeight(blkHeight, hex.EncodeToString(txId), txIndex)
	if err != nil {
		log.Printf("get block failed: %v", err)
		ctx.JSON(http.StatusOK, model.Response{Code: -1, Msg: "get txin failed"})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "ok",
		Data: result,
	})
}
