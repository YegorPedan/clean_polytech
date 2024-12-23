# Название проекта

Краткое описание проекта.

## Установка

Инструкции по установке:

```bash
# Клонировать репозиторий
git clone https://github.com/YegorPedan/clean_polytech.git

# Перейти в директорию проекта
cd clean_polytech

# Установить зависимости
go mod download
```

## API Эндпоинты

### GET /api/v1/resource

- **Описание**: Получить список всех ресурсов.
- **Параметры запроса**:
  - `limit` (опционально): Максимальное количество возвращаемых записей.
  - `offset` (опционально): Смещение для пагинации.
- **Пример запроса**:
  ```bash
  curl -X GET "https://example.com/api/v1/resource?limit=10&offset=0"
  ```
