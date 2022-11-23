# GMG (GUI Mini Game)

Check your attention!!!


## Description

The game is designed to calculate the speed of finding numbers from 1 to 25 in ascending order.

![1](https://user-images.githubusercontent.com/57867110/193422217-9e06a2d6-3532-4f01-b4b5-865e595337aa.png)
![2](https://user-images.githubusercontent.com/57867110/193422224-25741311-2805-402a-a7cb-3c267b2a6020.png)
![3](https://user-images.githubusercontent.com/57867110/193422230-261d0eda-2713-4aa2-b4db-621c55c57b95.png)


## Getting Started

### Dependencies

* Go 1.16+
* Windows, Linux or Mac OS
* MinGW-W64 for building

### Installing

```
Server:
docker pull postgres
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD={PASSWORD} -e POSTGRES_USER={USERNAME} -d postgres
git clone https://github.com/zagart47/GMG.git
cd GMG/server
docker build -t server .
docker run --env DBHOST=postgres://{USERNAME}:{PASSWORD}@{DBHOST}:5432/postgres --name server -p 80:80
go build .
```

```
Client:
git clone https://github.com/zagart47/GMG.git
cd GMG
go build.go
```

### Executing program

```
go run .
```

## Authors

Artur Zagirov  
[@zagart47](https://t.me/zagart47)