# harbor-models

harbor models from https://github.com/goharbor/harbor/tree/main/api

```
swagger generate model -f swagger/swagger.yaml --template-dir swagger/templates --additional-initialism=CVE --additional-initialism=GC

mv models/* ../
rm -rf models
```