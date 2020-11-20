# TODO
game & externalSystem : many2one relationship
* update createGame
* update gameModel
* update externalSystemModel (add description -> for web interface (select) )

# Описание модели хранения данных

## Таблицы

### games
    Игры
* game_id - id игры
* ext_system_id - id внешней системы
* name - название игры
* start_date - дата начала игры (ISO timestamp)
* end_date - дата конца игры (ISO timestamp)
* answer_type - тип ответа на вопрос 
    - '1' - Текст 
    - '2' - Категориальный
    - '3' - Координаты прямоугольника
    - '4' - Полигональный
* question - формулировка вопроса
* options_csv - опции, если "answer_type" == '2'
    - пример: 'Есть нарушение,Нет нарушения' 

### ext_systems
    Внешние системы
* ext_system_id - id внешней системы
  - если его не передают, то генерируем его на нашей стороне
* description - описание внешней системы
* post_results_url - url, по которому будет отправлять POST запрос с результатами

### sources
    Источники скриншотов (расписания, zip-архивы, url(?))
* source_id - id источника
* game_id - id игры, к которой относятся скриншоты из этого источника
* source_type - тип источника
    - 0 - zip-архив
    - 1 - расписание
* created_at - дата создания источника
    - для zip-архива - дата загрузки
    - для расписания - дата расписания (приходит от внешнего api)

### screenshots
    Скриншоты (Задачи)
* screenshot_id - id задания
* game_id - id игры
* source_id - id источника скриншотов
* filename - имя файла скриншота
* expert_answer - ответ эксперта
* users_answer - ответ пользователей

### users
    Пользователи
* inner_user_id - внутренний id пользователя
* user_id - id пользователя во внешней системе

### users__ext_systems
    Таблица-связка для отношения пользователь-внешняя система
* ext_system_id - id внешней системы
* user_id - id пользователя во внешней системе

### answers
    Ответы пользователей
* answer_id - id ответа
* game_id - id игры
* screenshot_id - id скриншота
* user_id - id пользователя во внешней системе
* value - ответ
