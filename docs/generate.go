package main

// これは上手く動く
// oapi-codegen --config=models.cfg.yaml openapi.yaml
// oapi-codegen --config=server.cfg.yaml openapi.yaml

// 公式にはこう書いてあるけどこれは上手く動かない気がする...(多分)
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=models.cfg.yaml openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml openapi.yaml
