build:
	gcloud builds submit --tag gcr.io/fader4/hash-slot-demo1
deploy:
	gcloud run deploy hash-slot-demo1 \
		--platform=managed \
		--region=europe-west1 \
		--image gcr.io/fader4/hash-slot-demo
# deploy:
# 	gcloud run deploy hash-slot-demo1 \
# 		--platform=managed \
# 		--region=europe-west1 \
# 		--image gcr.io/fader4/hash-slot-demo \
# 		--vpc-connector=projects/fader4/locations/europe-west1/connectors/demo2
