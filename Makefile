.PHONY: docker-image

docker-image:
	docker build -t us.gcr.io/sourcegraph-dev/gddo .
	docker push us.gcr.io/sourcegraph-dev/gddo
