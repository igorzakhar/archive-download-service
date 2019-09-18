# Archive download service
Микросервис для загрузки файлов единым архивом. Реализация на Go. 

Создание архива происходит на лету по запросу от пользователя. Архив не сохраняется на диске, вместо этого по мере упаковки он сразу отправляется пользователю на скачивание.  

Данный репозиторий является попыткой переписать код сервиса c Python на Go: [https://github.com/igorzakhar/async-download-service](https://github.com/igorzakhar/async-download-service).

# Установка

Для работы требуется предустановленный архиватор zip.  
Установка в ОС Debian:
```
sudo apt-get install zip
```

Так же потребуется установка стороннего go пакета  ```httprouter```([github.com/julienschmidt/httprouter](github.com/julienschmidt/httprouter)):
```bash
$ go get github.com/julienschmidt/httprouter
```

# Использование

Скопируйте данный репозиторий в каталог ```$GOPATH/src/```.

```bash
$ git clone https://github.com/igorzakhar/archive-download-service.git
```

Перейдите в каталог ```archive-download-service```:
```bash
$ cd archive-download-service
```

Запуск программы:
```bash
$ go run archive_service.go
```

После запуска сервис будет доступен по адресу [http://127.0.0.1:8080/](http://127.0.0.1:8080)

# Цели проекта

Проект создан в учебных целях. 