gen:
	oapi-codegen -package openapi -generate types openapi.yaml > internal/openapi/types.gen.go
deploy:
	 vercel --prod