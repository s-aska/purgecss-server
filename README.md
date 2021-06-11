# purgecss-server

purgecss on GAE/SE

## How to deploy purgecss server

```
make deploy-purgecss PROJECT_NAME=YOUR_GCP_PROJECT_ID
```

## How to run local purgecss server

```
node app.js
```

## How to get purgecss

```
curl -sS http://localhost:8080 -X POST -F "html=@index.html" -F "css=@main.css"
```

```
go run main.go https://office-aska.com https://office-aska.com/static/website/css/style.css http://localhost:8080
```

```
go run main.go https://office-aska.com https://office-aska.com/static/website/css/style.css https://***.an.r.appspot.com
```

