---
updated_at: 2025-12-17T15:08:16.730+10:00
tags:
  - bash
  - string
  - array
---
## Работа с массивами строк в Bash

### 1. **Объявление и инициализация массивов**

```bash
# Явное объявление (необязательно)
declare -a my_array

# Создание с элементами
fruits=("apple" "banana" "orange" "grape")

# Создание поэлементно
colors[0]="red"
colors[1]="green"
colors[2]="blue"

# Чтение строк из команды
files=(*.txt)  # все txt-файлы в директории
lines=($(cat file.txt))  # разбивка по пробелам
```

### 2. **Основные операции**

```bash
# Доступ к элементам
echo ${fruits[0]}    # "apple" (первый элемент)
echo ${fruits[-1]}   # последний элемент
echo ${fruits[@]}    # все элементы

# Длина массива
echo ${#fruits[@]}   # количество элементов
echo ${#fruits[0]}   # длина первого элемента

# Добавление элементов
fruits+=("kiwi")     # в конец
fruits=("mango" "${fruits[@]}")  # в начало

# Удаление
unset fruits[1]      # удалить элемент с индексом 1
unset fruits         # удалить весь массив
```

### 3. **Итерация по массиву**

```bash
# По элементам
for fruit in "${fruits[@]}"; do
    echo "$fruit"
done

# По индексам
for i in "${!fruits[@]}"; do
    echo "$i: ${fruits[$i]}"
done

# Счетчик
for ((i=0; i<${#fruits[@]}; i++)); do
    echo "${fruits[$i]}"
done
```

### 4. **Полезные операции с массивами**

```bash
# Срез массива
echo ${fruits[@]:1:2}  # элементы с 1 по 2

# Поиск индекса (Bash 4.0+)
# Нет встроенной функции, но можно через цикл

# Копирование массива
copy=("${fruits[@]}")

# Объединение массивов
combined=("${array1[@]}" "${array2[@]}")

# Проверка существования элемента
if [[ " ${fruits[@]} " =~ " banana " ]]; then
    echo "Банан есть!"
fi
```

### 5. **Ассоциативные массивы (Bash 4.0+)**

```bash
declare -A dict
dict["name"]="John"
dict["age"]="30"

# Итерация по ключам
for key in "${!dict[@]}"; do
    echo "$key: ${dict[$key]}"
done
```

### 6. **Преобразования**

```bash
# Массив → строку с разделителем
IFS=','; echo "${fruits[*]}"; unset IFS
# или
printf '%s\n' "${fruits[@]}" | paste -sd ','

# Строка → массив
string="apple,banana,orange"
IFS=',' read -ra arr <<< "$string"
```

## **Ограничения и особенности**

### 1. **Индексирование**
- Индексы начинаются с 0
- Можно использовать отрицательные индексы: `-1` - последний элемент
- Пропуски в индексах разрешены
```bash
arr[0]="a"
arr[5]="b"  # допустимо, индексы 1-4 не существуют
```

### 2. **Разделители**
```bash
# IFS влияет на создание массивов
data="a:b:c"
IFS=':' read -ra arr <<< "$data"
```

### 3. **Особенности цитирования**
```bash
# Без кавычек - проблемы с пробелами
arr=($(echo "a b c"))  # 3 элемента
arr=($(echo "a 'b c'"))  # всё равно 3 элемента!

# Правильно - использовать mapfile/readarray
readarray -t arr <<< "$(echo "a 'b c'")"
```

### 4. **Ограничения Bash**
- Одномерные массивы (кроме ассоциативных)
- Нет встроенных функций для сортировки, поиска
- Максимальный индекс: 2^63-1 (обычно достаточно)
- Размер ограничен памятью

### 5. **Проблемные случаи**
```bash
# Пустые элементы теряются
arr=("a" "" "c")  # ${#arr[@]} = 3, но
arr=($(echo "a  c"))  # ${#arr[@]} = 2!

# Элементы с пробелами
arr=("file with spaces.txt")  # 1 элемент
arr=(file with spaces.txt)    # 4 элемента!
```

## **Рекомендации**

1. **Всегда используйте кавычки:**
```bash
for item in "${array[@]}"; do
    # ...
done
```

2. **Для сложных операций** используйте внешние утилиты:
```bash
# Сортировка
sorted=($(printf '%s\n' "${array[@]}" | sort))

# Фильтрация
filtered=($(printf '%s\n' "${array[@]}" | grep "pattern"))
```

3. **Bash 4.0+** предоставляет больше возможностей:
```bash
# Проверка версии
if ((BASH_VERSINFO[0] >= 4)); then
    # используем ассоциативные массивы и пр.
fi
```

4. **Альтернативы для сложных структур:**
- Используйте Python/Perl для сложных структур данных
- Рассмотрите `jq` для работы с JSON
- Используйте временные файлы для больших данных

## **Полезные функции**

```bash
# Уникальные значения
unique=($(printf '%s\n' "${array[@]}" | sort -u))

# Пересечение массивов
comm -12 <(printf '%s\n' "${arr1[@]}" | sort) \
         <(printf '%s\n' "${arr2[@]}" | sort)

# Разность массивов
grep -vxFf <(printf '%s\n' "${arr2[@]}") \
           <(printf '%s\n' "${arr1[@]}")
```

Массивы в Bash достаточно мощны для большинства задач системного администрирования, но для сложной обработки данных лучше использовать другие языки.

### **Добавление строки в конец массива (append)**

```bash
#!/bin/bash

# Функция для добавления строки в конец массива
# Использование: append <имя_массива> <строка_для_добавления>
append() {
    # Проверка количества аргументов
    if [[ $# -ne 2 ]]; then
        echo "Ошибка: функция ожидает 2 аргумента - имя массива и строку" >&2
        return 1
    fi

    local -n arr_ref="$1"  # Создаем локальную ссылку на массив (nameref)
    local value="$2"       # Локальная переменная для значения
    
    # Добавляем значение в конец массива
    arr_ref+=("$value")
}


# Пример использования:
# Объявляем глобальный массив
declare -a fruits=("apple" "banana" "orange")
# Добавляем элемент с помощью функции
append fruits "grape and tomato"
append fruits "pineapple"

# Выводим окончательный массив
echo "После добавления 'pineapple':"
printf '%s\n' "${fruits[@]}"
echo "Длина массива: ${#fruits[@]}"

# РАБОТАЕТ
```
### **Разделение строки на массив по разделителю (split_string_to_array)**
```bash
# Функция для разделения строки на массив по разделителю
# Использование: split_string_to_array <имя_массива> "<исходная_строка>" "<разделитель>"
split_string_to_array() {
    if [[ $# -ne 3 ]]; then
        echo "Ошибка: ожидается имя массива, строка и разделитель" >&2
        return 1
    fi
    
    local -n arr_ref="$1"
    local string="$2"
    local delimiter="$3"
    
    arr_ref=()
    
    # Обработка пустой строки
    if [[ -z "$string" ]]; then
        return 0
    fi
    
    # Обработка пустого разделителя
    if [[ -z "$delimiter" ]]; then
        arr_ref+=("$string")
        return 0
    fi
    
    # Проверка наличия разделителя в строке
    if [[ "$string" != *"$delimiter"* ]]; then
        arr_ref+=("$string")
        return 0
    fi
    
    # Основной алгоритм
    local rest="$string"
    
    while [[ "$rest" == *"$delimiter"* ]]; do
        # Извлекаем часть до первого разделителя
        local part="${rest%%"$delimiter"*}"
        arr_ref+=("$part")
        
        # Удаляем обработанную часть
        rest="${rest#*"$delimiter"}"
        
        # Обработка случая, когда разделитель в конце
        if [[ -z "$rest" ]]; then
            arr_ref+=("")
            break
        fi
    done
    
    # Добавляем последнюю часть (если осталась)
    if [[ -n "$rest" ]]; then
        arr_ref+=("$rest")
    fi
}

# Примеры использования

echo "=== Пример 1: Базовое использование ==="
declare -a result1
split_string_to_array result1 "apple,banana,orange,grape" ","
echo "Массив: ${result1[@]}"
echo "Количество элементов: ${#result1[@]}"
for i in "${!result1[@]}"; do
    echo "  [$i] = '${result1[i]}'"
done


```

### **Соединение массива в строку с заданным делиметром (join_array)**

```bash

# Функция для соединения элементов массива с разделителем
# Использование: join_array <имя_массива> <разделитель>
# Возвращает: строку с элементами массива, разделенными разделителем JOINED=$(join_array fruits "--")
join_array() {
    if [[ $# -lt 2 ]]; then
        echo "Ошибка: ожидается имя массива и разделитель" >&2
        return 1
    fi
    
    local -n arr_ref="$1"
    local delimiter="$2"
    
    # Обработка пустого массива
    if [[ ${#arr_ref[@]} -eq 0 ]]; then
        echo ""
        return 0
    fi
    
    # Обработка массива из одного элемента
    if [[ ${#arr_ref[@]} -eq 1 ]]; then
        echo "${arr_ref[0]}"
        return 0
    fi
    
    # Используем printf для эффективного соединения
    local first_element="${arr_ref[0]}"
    printf "%s" "$first_element"
    
    # Соединяем остальные элементы
    for element in "${arr_ref[@]:1}"; do
        printf "%s%s" "$delimiter" "$element"
    done
    
    # Добавляем перевод строки в конце
    echo ""
}

# Примеры использования

echo "=== Пример 1: Базовое использование ==="
declare -a fruits=("apple" "banana" "fish and chips" "grape")
JOINED=$(join_array fruits "--")
echo "${JOINED}"
