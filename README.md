# go-anime-image

* Automatically download anime images and crop face

## Sample 
* original image  
<img src="./download/rail_romanesque_1.jpg" width="240"/>

* cropped image  
<img src="./face/0_rail_romanesque_1.jpg" width="120px"/> <img src="./face/1_rail_romanesque_1.jpg" width="120px"/> <img src="./face/2_rail_romanesque_1.jpg" width="120px"/> <img src="./face/3_rail_romanesque_1.jpg" width="120px"/> <img src="./face/4_rail_romanesque_1.jpg" width="120px"/> <img src="./face/5_rail_romanesque_1.jpg" width="120px"/> <img src="./face/6_rail_romanesque_1.jpg" width="120px"/>

## Requirement

* [Docker](https://www.docker.com/get-started)
* build docker images first at `./go-anime-image`

```
docker-compose build
```

## Supported Site

* https://tsundora.com 
* https://wallpaperboys.com

## How to download images

* Images are downloaded in `./download`  
  
* Download from https://tsundora.com
```
export SCRAPING_SITE=tsundora.com
docker-compose up scraping  
```

* Download from https://wallpaperboys.com
```
export SCRAPING_SITE=wallpaperboys.com 
docker-compose up scraping 
```

* Download with keyword
```
export SCRAPING_SITE=tsundora.com
export SCRAPING_KEYWORD=iDOLM@STER
docker-compose up scraping 
```

## How to crop face

* crop face for all files is `./face`

```
docker-compose up cropping
```
