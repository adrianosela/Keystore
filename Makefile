all: build

clean:
	rm -rf pkg bin

deploy: dockerbuild down
	docker run -d --name keystore_service -p 8080:8888 keystore

up: build
	./Keystore

dockerbuild:
	./dockerbuild.sh

build:
	go build -o ./Keystore

down:
	(docker stop keystore_service || true) && (docker rm keystore_service || true)
