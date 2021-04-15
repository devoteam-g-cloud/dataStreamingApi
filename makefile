build-container-repo:
	gcloud builds submit --tag gcr.io/sandbox-fmalbranque/api-streaming-data

deploy-cloud-run:
	gcloud beta run deploy api-streaming-data \
	--image=gcr.io/sandbox-fmalbranque/api-streaming-data \
	--concurrency=1 \
	--cpu=1 \
	--memory=128M \
	--platform=managed \
	--port=8080 \
	--timeout=10 \
	--region=europe-west1 \
	--max-instances=1 \

build-local:
	docker run --rm -e PORT=8080 -p 8080:8080 gcr.io/sandbox-fmalbranque/api-streaming-data 
	#make sure docker deamon is working

build-and-deploy:
	make build && make deploy-cloud-run

