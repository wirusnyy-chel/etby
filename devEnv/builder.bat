cd ..
SET SERVICE_NAME=login
docker build -f services/%service_name%/server/Dockerfile -t etby-%service_name% .

SET SERVICE_NAME=register
docker build -f services/%service_name%/server/Dockerfile -t etby-%service_name% .