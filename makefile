build-gcr:
	gcloud builds submit --tag gcr.io/sandbox-fmalbranque/api-streaming-data

deploy-cloud-run:
	gcloud beta run deploy api-streaming-data \
	--image=gcr.io/sandbox-fmalbranque/api-streaming-data \
	--concurrency=5 \
	--cpu=1 \
	--memory=128M \
	--platform=managed \
	--port=8080 \
	--timeout=10 \
	--region=europe-west1 \
	--max-instances=2 \

build-local:
	docker image build . -t data-streaming

run-local:
	docker run --rm -e PORT=8080 -p 8080:8080 data-streaming



build-run-local:
	make build-local && make run-local
	#make sure docker deamon is working

build-and-deploy:
	make build-gcr && make deploy-cloud-run

