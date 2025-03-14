# Домашнее задание — Тестирование приложения wallet

## Дано

На неделе мы окунулись во внутрянку тестирования на Golang: узнали про тот самый testing.T, рассмотрели allure-go и чутка попрактиковались

## Задание

1. Необходимо переписать тесты из ДЗ-2 с использованием фреймворка allure-go
    - https://github.com/ozontech/allure-go
2. Для удобства запуска тестов необходимо добавить make команду (в Makefile), запускающую ваши е2е-тесты
    - пример - `make e2e-tests`
    - после прогона тестов должен формировать сырой отчет - `allure-results`
3. *Написать е2е-тесты, использующие kafka
    - Логика сервиса при работе с Kafka
      - сервис с определенным интервалом вычитывает сообщеиня из Kafka и сохраняет в `локальное хранилище`
      - вместимость локального хранилища - 1000 элементов
        - после переполнения хранилища, старые записи удаляются
          - Например: было 950 элементов, добавили 60 элементов -> первые 10 элементов удаляются
    - Для работы с Kafka, в сервисе реализованы 2 ручки
      - produce - записывает сообщение в кафку
      - consume - возвращает указанное кол-во элементов из локального хранилища
    - сценарий на ваше усмотрение
4. *Добавить make команду (в Makefile), открывающую allure-отчет в браузере
    - подсказка: используй `allure serve`


ps: задачи с * - дополнительные

pss: необходимо перенести себе изменения из МР - https://gitlab.ozon.dev/qa/classroom-14/students/e2e-backend/-/merge_requests/1/diffs
 - app/*
 - internal/pb/*


## Общие критерии приемки
### Обязательные задачи
1. Решены 2 обязательные задачи: п.1 и п.2
2. Чистота кода 
3. Корректность написанных тестов
4. После прогона тестов из п.2 должен формироваться сырой отчет - `allure-results`

### Дополнительные задачи *
1. п.3 - все ожидания должны быть реализованы через pooling
2. п.4 - у пользователя, который захочет посмотреть отчет, может не быть необходимых доп зависимостей, таких как `allure`. Необходимо добавить в readme.md инструкцию для их установки
    - например: "скачать отсюда и положи сюда"

При выставлении баллов учитывается чистота и правильность написанного кода.

### Дедлайны сдачи и проверки задания: 
- 21 сентября 23:59 (сдача) / 24 сентября, 23:59 (проверка)