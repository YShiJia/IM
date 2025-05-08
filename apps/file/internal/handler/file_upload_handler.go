package handler

import (
	"fmt"
	"github.com/YShiJia/IM/lib/data"
	"io"
	"net/http"

	"github.com/YShiJia/IM/apps/file/internal/logic"
	"github.com/YShiJia/IM/apps/file/internal/svc"
	"github.com/YShiJia/IM/apps/file/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//	return
		//}
		file, fileHeader, err := r.FormFile("file")

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("服务端接收文件失败 err: %v", err))
			return
		}
		defer file.Close()
		if fileHeader.Size > svcCtx.Config.MaxFileBytes*data.MB {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("上传文件大小超出单次上传最大限制: %dMB", svcCtx.Config.MaxFileBytes))
			return
		}
		fileData, err := io.ReadAll(file)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, fmt.Errorf("服务端读取文件失败 err: %v", err))
			return
		}
		req.FileName = fileHeader.Filename
		req.Size = fileHeader.Size

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req, fileData)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
