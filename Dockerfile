##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.24-alpine AS build

ARG PORT

ENV PORT=$PORT

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY . .

# RUN apk add --no-cache make

# RUN make generate.client name=email
# RUN make generate.client name=wuzapi

# download Go modules and dependencies
RUN go mod tidy

EXPOSE ${PORT}

# compile application
RUN go build -o /binary

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /binary /binary

ENTRYPOINT ["/binary"]