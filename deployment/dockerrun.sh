mkdir -p data
docker run -d \
  --name goslug \
  -p 8080:8080 \
  -e HOST=0.0.0.0 \
  -e USERNAME="admin" \
  -e PASSWORD="password" \
  -e STORAGE_DRIVER="json" \
  -v $(pwd)/data:/app/data \
  goslug:latest