---
updated_at: 2026-01-20T11:05:24.019+10:00
tags:
  - commands
  - database
---
Программа для работы с данными миров.
# Инфраструктура
## Базы данных
Программа использует 2 базы данных.
- second_survey_data - External_DB
- worlds_data - Derived_DB

# Операции
[[import (trv_worlds)|import]] - скачать данные с сайта https://travellermap.com
[[survey (trv_worlds)|survey]] - дополнить отсутствующие данные в оперативную базу (для всех файлов external_db)
[[inspect (trv_worlds)|inspect]] - просмотреть данные по миру содержащиеся в External_DB