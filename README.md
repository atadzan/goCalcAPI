### GoCalcAPI - cервис подсчёта арифметических выражений<br>

##### Команда для запуска проекта в локальной среде: ```go run cmd/main.go --port=8080```

##### Команда для запуска тестов хэндлера: ```go test ./pkg/handlers -v```

##### Эндпоинты: 
* POST <b>/api/v1/calculate</b><br>
На вход ожидается ```{ "expression": "выражение, которое ввёл пользователь" }```
  * Ответ в случае успешной обработки: HTTP_STATUS: 200 и тело ответа ```{ "result": "результат выражения" }```
  * Ответ в случае ошибок: 
    * Если входные данные не соответствуют требованиям приложения: HTTP_STATUS: 422 и тело ответа  ```{ "error": "Expression is not valid" }```
    * В случае какой-либо иной ошибки («Что-то пошло не так»): HTTP_STATUS: 500 и тело ответа ```{ "error": "Internal server error" }```<br>
  * Примеры запросов:
    * Кейс успешного запроса
      ``` 
        curl --location 'localhost:8080/api/v1/calculate' \
            --header 'Content-Type: application/json' \
            --data '{ "expression": "2+2*2"}' 
      ```
      Ответ: ```{"result":6}```</br>
    * Кейс невалидного выражения 
      ``` 
        curl --location 'localhost:8080/api/v1/calculate' \
            --header 'Content-Type: application/json' \
            --data '{ "expression": "a+2*2"}' 
      ```
      Ответ: ```{"error":"Expression is not valid"}```
    * Кейс для других ошибок
      ``` 
        curl --location 'localhost:8080/api/v1/calculate' \
            --header 'Content-Type: application/json' \
            --data '{ "expression": "10/0"}' 
      ```
      Ответ: ```{"error":"Internal server error"}```
      