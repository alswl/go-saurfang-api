# go-saurfang

```
curl 'http://localhost:8080/api-docs/swagger.json' > api/openapi.yaml

rm client/zz_generated_*.go;rm client/*/zz_generated_*.go; rm models/zz_generated_*.go; swagger generate client -c client -f ./api/openapi.yaml -A saurfang --template-dir ./api/templates --allow-template-override -C ./api/config.yaml
```
