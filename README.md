1) Калькулятор
   
AsyncCalculationWebServiceProject — это проект представляет собой асинхронный веб-сервис, разработанный на языке Go. Он предназначен для обработки HTTP-запросов и предоставления данных пользователям через API. Этот сервис для выполнения математических вычислений на основе переданного выражения. Сервис предоставляет API для обработки запросов на вычисления, а также включает в себя набор тестов для проверки корректности работы.

2) Проект написан на языке Go, организован в модульной структуре и содержит примеры использования.

3) Установка и запуск
   
Для запуска проекта выполните следующие шаги:

4) Склонируйте репозиторий:
   
git clone https://github.com/Vlod-and-whot/AsyncCalculationWebServiceProject.git
cd calc_go
Убедитесь, что Go установлен и находится в $PATH (проверить версию можно командой go version).

5) Запустите API-сервер:

go run ./cmd/main.go
Сервер запустится на порту 8080 по умолчанию. Если необходимо изменить порт, установите переменную окружения PORT перед запуском:

В Linux и macOS
export PORT=9090
go run ./cmd/main.go
В Windows (PowerShell)
$env:PORT=9090
go run ./cmd/main.go

6) Использование API
Эндпоинт
POST /api/v1/calculate
Заголовки
Content-Type: application/json
Тело запроса
Пример:

{
  "expression": "2+2*2"
}
Ответы
Успешный запрос

Статус-код: 200 OK
Пример ответа:

{
  "result": "6"
}
Ошибка обработки выражения

Статус-код: 422 Unprocessable Entity
Пример ответа:

{
  "error": "Error calculation"
}
Неподдерживаемый метод

Статус-код: 405 Method Not Allowed
Пример ответа:

{
  "error": "Wrong Method"
}
Некорректное тело запроса

Статус-код: 400 Bad Request
Пример ответа:

{
  "error": "Invalid Body"
}
Примеры использования
Успешный запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
Ответ:

{
  "result": "6"
}
Ошибка: некорректное выражение:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2*(2+2{)"
}'
Ответ:

{
  "error": "Error calculation"
}
Ошибка: неверный метод:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json'
Ответ:

{
  "error": "Wrong Method"
}
