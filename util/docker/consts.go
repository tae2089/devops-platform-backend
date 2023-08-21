package docker

const dockerJava = `
	FROM adoptopenjdk/openjdk11:x86_64-alpine-jdk-11.0.20_8-slim
	ENV APP_HOME=/apps
	ARG JAR_FILE_PATH=build/libs/demo-0.0.1-SNAPSHOT.jar
	WORKDIR $APP_HOME
	COPY $JAR_FILE_PATH app.jar
	EXPOSE 8080
	ENTRYPOINT ["java", "-jar", "app.jar"]
	`
const dockerGolang = `
    FROM golang:1.21.0-alpine3.18 as build
	RUN apk --no-cache add tzdata ca-certificates
	WORKDIR /src/
	COPY . /src/
	RUN CGO_ENABLED=0 go build -o /bin/main run/main.go
	
	FROM scratch
	COPY --from=build /bin/main /bin/main
	COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
	COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
	ENV TZ=Asia/Seoul
	ENTRYPOINT ["/bin/main"]
    `

const dockerNode = `
	FROM node:16 as builder
	WORKDIR /app
	COPY . .
	RUN npm ci 
	RUN npm run build

	FROM node:16
	WORKDIR /app
	COPY --from=builder /app/dist /app/dist
	COPY --from=builder /app/node_modules /app/node_modules
	USER node
	CMD ["node", "dist/src/main"]
	`
