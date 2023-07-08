# Le Monde 3.0 - Web

## Run with npm

- Install dependendies
```shell
npm install
```

- ENV setup (don't forget to update the `.env` file created, if needed)
```shell
cp .env.example .env
```

- Run server
```shell
npm start
```

## Run with docker

- Build docker image
```shell
docker build . -t lemonde-web:latest
```

- Run docker image
```shell
docker run -p 3000:80 lemonde-web:latest
```