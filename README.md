# WBABEProject-11

## 기술 스택

golang (go1.19.4 linux/arm64)
MongoDB

## 프로젝트 이름

띵동 주문이요, 온라인 주문시스템(Online Ordering System)

## 프로젝트 개요

언택트 시대에 급증하고 있는 온라인 주문 시스템은 이미 생활전반에 그 영향을 끼치고 있는 상황에, 가깝게는 배달어플, 매장에는 키오스크, 식당에는 패드를 이용한 메뉴 주문까지 그 사용범위가 점점 확대되어 가고 있습니다. 이런 시대에 해당 시스템을 이해, 경험하고 각 단계별 프로세스를 이해하여 구현함으로써 서비스 구축에 경험을 쌓고, golang의 이해를 돕습니다.

## 프로젝트 목표

1. 학습자는 주문자/피주문자의 역할에서 필수적인 기능을 도출, 구현합니다.
2. 학습자는 해당 시스템에 대해 요구사항을 접수하고 주문자와 피주문자 입장에서 필요한 기능을 도출하여, 기능을 확장하고 주문 서비스를 원할하게 지원할수 있는 기능을 구현합니다.
3. 주문자는 신뢰있는 주문과 배달까지를 원합니다. 또, 피주문자는 주문내역을 관리하고 원할한 서비스가 제공되어야 합니다.

### 프로젝트 설치 및 실행

```
// 설치
git clone git@github.com:codestates/WBABEProject-11.git Project
cd Project
go mod tidy


// 실행
go run main.go

```

### API 명세서

```
// 메뉴
POST /menu -> 메뉴 추가
PUT /menu -> 메뉴 업데이트
DELETE /menu/:name -> 메뉴 삭제
GET /menu/ -> 메뉴 전체 조회
GET  /menu/:name -> 메뉴 조회

// 메뉴 리뷰
GET /menu/review/:name -> 리뷰 조회
POST /menu/review -> 리뷰 작성

// 주문
POST /order -> 주문 추가
PUT /order -> 주문 업데이트
GET  /order/:name -> 주문 조회
GET  /order/status -> 주문 상태 조회

```

### API 테스트

## DB 설계

### DB 구조

```
// Menu DB
type Menu struct {
	Name string `json:"name" bson:"name"`
	Soldout int `json:"soldout" bson:"soldout"`
	Stock int `json:"stock" bson:"stock"`
	Origin string `json:"origin" bson:"origin"`
	Price int `json:"price" bson:"price"`
}

```

```
// Order DB
type Order struct {
	Menu string `json:"menu" bson:"menu"`
	Name string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
	Address	string `json:"address" bson:"address"`
	Status int `json:"status" bson:"status"`
}

```

```
// Review DB
type Review struct {
	Name string `json:"name" bson:"name"`
	Rating int `json:"rating" bson:"rating"`
	OrderNumber int `json:"orderNumber" bson:"ordernumber"`
	Review string `json:"review" bson:"review"`
}

```
