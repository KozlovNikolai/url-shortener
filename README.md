<a name="readme-top"></a>
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">сокращатель ссылок по примеру Николая Тузова</h3>

<div align="left">

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Содержание</summary>
  <ol>
    <li><a href="#о-проекте">О проекте</a></li>
    <li><a href="#начало">Начало</a></li>
    <li><a href="#лицензия">Лицензия</a></li>
    <li><a href="#контакты">Контакты</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## О проекте

![Product Name Screen Shot][front-img]

<p align="left">
Написал "полноценный" REST API сервис — URL Shortener — и задеплоил его на виртуальный сервер с помощью GitHub Actions.</p>
<p align="left">
Говоря «полноценный», имеется в виду, что это будет не игрушечный проект, а готовый к использованию:</p>

   * <p align="left">выбран актуальный http-роутер chi,</p>
   * <p align="left">сделано полноценное логирование,</p>
   * <p align="left">написаны тесты: unit-тесты, тесты хэндлеров и функциональные,</p>
   * <p align="left">настроен автоматический деплой через GitHub Actions и др.</p>


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Начало

Для запуска проекта требуется:

1. Установить git и golang:
```
sudo apt update
sudo apt install git
wget https://go.dev/dl/go1.21.10.linux-amd64.tar.gz
mkdir go1.21.10
sudo rm -rf /usr/local/go 
sudo tar -xvf go1.21.10.linux-amd64.tar.gz -C ~/go1.21.10
export PATH=$PATH:/usr/local/go/bin
```
2. Проверить версию Golang:
```
go version
```
3. Склонировать репозиторий к себе на компьютер:
```
git clone https://github.com/KozlovNikolai/url-shortener.git
```
4. перейти в каталог с проектом:
```
cd url-shortener
```
5. установить значение переменной окружения:
```
export CONFIG_PATH="$HOME/url-shortener/config/local.yaml"
```
6. удалить файлы:
```
rm go.mod
rm go.sum
```
7. обновить зависимости:
```
go mod init url-shortener
go mod tidy
```
8. запустить сервер:
```
go run cmd/url-shortener/main.go
```


<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## Лицензия

Распостраняется свободно.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Контакты

Николай Козлов:\
Telegram: @gremiha3\
e-mail: gremiha3@yandex.ru

Project Link: [https://github.com/KozlovNikolai/manipulator](https://github.com/KozlovNikolai/manipulator)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[arduino-img]: https://cdn.arduino.cc/header-footer/prod/assets/headerLogo-arduino.svg
[arduino-link]: https://www.arduino.cc/
[front-img]: images/url-shortener.png