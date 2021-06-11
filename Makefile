
deploy-purgecss:
	gcloud app deploy --project $(PROJECT_NAME) purgecss.yaml
