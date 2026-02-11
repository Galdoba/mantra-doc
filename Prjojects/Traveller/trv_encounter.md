---
updated_at: 2026-01-31T17:01:02.447+10:00
---
программа для быстрой генерации энкаунтеров по таблицам.

# Комманды

`trv_encounter new` - создает новый файл энкаунтера (печать в тект, сохранить в базу)
`trv_encounter recall` - ищет уже имеющийся файл энкаунтера (печать в тект, сохранить в базу)

# Домены
Encounter: описание энкаунтера.
	CreatedAt time.Time
	LastEdited time.Time
	Tables []string
	Keys []string
	Description string
	