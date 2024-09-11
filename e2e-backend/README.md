# E2E Backend

Проект с Е2Е тестами для проекта wallet.

### Установка зависимостей

```bash
make install
```

### Запуск тестов

```bash
make test
```

Чтобы запустить сервис в docker-compose необходимо:
1. Перейти в директорию с compose.yaml
    ```shell
    cd ./app
    ```
2. Авторизоваться под своей учеткой
    ```shell
    docker login gitlab.ozon.dev
    ```
3. Спулить =)
    ```shell
    make app-up
    ```
