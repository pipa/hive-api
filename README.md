This is the API readme

To directly run this container use the following:
```bash
# For development
docker build -t api .
docker run -p 8888:8888 -v [put your path here]/:/go/src/github.com/pipa/app api

# For production
docker build -t api-prd . --build-arg app_env=production
docker run -i -t -p 8888:8888 api
```
