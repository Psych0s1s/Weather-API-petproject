Структура проекта:

Weather-API-petproject/
├── cmd/
│   ├── cli/
│   │   └── main.go  # CLI версия
│   └── server/
│       └── main.go  # Серверная версия
└── weather/
    └── weather.go   # Логика пакета weather


Реализация:
CLI версия (cmd/cli/main.go): содержит код для версии приложения, которая работает как командно-строковое приложение. Здесь пользователь может ввести название города через командную строку, и приложение выведет данные о погоде.
Серверная версия (cmd/server/main.go): содержит код для версии приложения, работающего как веб-сервер. Пример вызова: http://localhost:8080/weather?city=Название города

Запуск:
Для запуска каждой из версий вам нужно будет перейти в соответствующую директорию и выполнить команду go run main.go или скомпилировать каждую версию в исполняемый файл с помощью go build, который затем можно будет запускать напрямую.

Обратите внимание, что для правильной работы пакета weather.go необходимо получить ключ API на сайте OpenWeatherMap API ( https://openweathermap.org/api ) и установить его в переменные среды вашей ОС командой:
для Windows: set OPENWEATHERMAP_API_KEY=ваш_api_ключ
для MacOS или Linux: export OPENWEATHERMAP_API_KEY=ваш_api_ключ
