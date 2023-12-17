
all: run-apps
build: build-apps

.PHONY: run-apps build-apps
run-apps: app1 app2 app3
build-apps: build-app1 build-app2 build-app3

.PHONY: build-app1 build-app2 build-app3
build-app1: APP=app1
build-app2: APP=app2
build-app3: APP=app3
build-app1 build-app2 build-app3:
	@echo building ${APP}
	@echo --------------------------
	@go build -o build/${APP} cmd/${APP}/main.go
	@echo ==========================

.PHONY: app1 app2 app3
app1: APP=app1
app2: APP=app2
app3: APP=app3
app1 app2 app3:
	@echo running ${APP}
	@echo --------------------------
	@go run cmd/${APP}/main.go
	@echo ==========================

test:
	for s in $$(go list ./...); do if ! ${ENV} go test -count=1 -failfast -v -p 1 $$s; then break; fi; done