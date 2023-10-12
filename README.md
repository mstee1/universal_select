# universal_select

## О проекте
Проект демонстрации универсального select запроса к базе данных.

Посмотреть версию программы можно задав ключ при запуске. Если используется этот ключ, программа только выведет версию, но не запустится:
```bash
--version

Пример вывода
go run main.go --version

1.0.0
```

## Структура конфигурационного файла (Необходимо соблюдать вложенность)

```YAML
# Параметры подключения к БД
database:
  dbPort: 
  dbHost: 
  dbName: 
  dbUser: 
  dbPassword: 
# Sql запросы
sql:
  select1: ""
  select2: ""
  select3: ""
  select4: ""
```

В данном примере представлена функция
```GO
func (r *req) SelectData(ctx context.Context, query string) ([][]interface{}, error) 
```

Она способна выполнять любой запрос select и получать любые типы данных (протестированы типы - int32,int64,string,time.Time)

Пример выполнения первого запроса. Получается поле "id" типа integer(В бд этот тип является 32 разрядным).

```GO
func Select1(ctx context.Context, cfg *config.Config, req psql.Requests) {

	data, err := req.SelectData(ctx, cfg.Sql.Select1)
	if err != nil {
		fmt.Println(err)
		return
	}

	var id int32
	for _, line := range data {
		id = line[0].(int32)
		fmt.Println(id)
	}
}
```