gen:
	oapi-codegen -package openapi -generate types openapi.yaml > openapi/types.gen.go
dev:
	vercel dev
deploy:
	 vercel --prod