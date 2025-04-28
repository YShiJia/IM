### 1. N/A

1. route definition

- Url: /v1/file/upload
- Method: POST
- Request: `FileUploadRequest`
- Response: `FileUploadReply`

2. request definition



```golang
type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext string `json:"ext,optional"`
	Size int64 `json:"size,optional"`
	Path string `json:"path,optional"`
}
```


3. response definition



```golang
type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext string `json:"ext"`
	Name string `json:"name"`
}
```

### 2. N/A

1. route definition

- Url: /v1/file/upload/chunk
- Method: POST
- Request: `FileUploadChunkRequest`
- Response: `FileUploadChunkReply`

2. request definition



```golang
type FileUploadChunkRequest struct {
}
```


3. response definition



```golang
type FileUploadChunkReply struct {
	Etag string `json:"etag"` // MD5
}
```

### 3. N/A

1. route definition

- Url: /v1/file/upload/chunk/complete
- Method: POST
- Request: `FileUploadChunkCompleteRequest`
- Response: `FileUploadChunkCompleteReply`

2. request definition



```golang
type FileUploadChunkCompleteRequest struct {
	Md5 string `json:"md5"`
	Name string `json:"name"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	Key string `json:"key"`
	UploadId string `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}
```


3. response definition



```golang
type FileUploadChunkCompleteReply struct {
	Identity string `json:"identity"` // 存储池identity
}
```

### 4. N/A

1. route definition

- Url: /v1/file/upload/prepare
- Method: POST
- Request: `FileUploadPrepareRequest`
- Response: `FileUploadPrepareReply`

2. request definition



```golang
type FileUploadPrepareRequest struct {
	Md5 string `json:"md5"`
	Name string `json:"name"`
	Ext string `json:"ext"`
}
```


3. response definition



```golang
type FileUploadPrepareReply struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key string `json:"key"`
}
```

