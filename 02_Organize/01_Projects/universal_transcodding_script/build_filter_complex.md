---
updated_at: 2025-12-17T20:25:53.571+10:00
tags:
  - universal_transcoding_script
  - filter_complex
---
```bash
#!/bin/bash

# Параметры видео
VIDEO_FRAMERATE=25                                 # Частота дискретизации видео (fps)

# Флаги включения видео
VIDEO_ENABLED=true                                 # Включить обработку видео
VIDEO_PROXY_ENABLED=true                           # Включить создание proxy видео

# Флаги включения аудио
AUDIO_ENABLED=true                                 # Включить обработку аудио
AUDIO_PROXY_ENABLED=true                           # Включить создание proxy аудио

# Абстрактные переменные фильтров видео
VIDEO_FILTER_1=""                                  # Дополнительный фильтр для видео
VIDEO_PROXY_FILTER_1=""                            # Дополнительный фильтр для proxy видео

# Абстрактные переменные фильтров аудио
AUDIO_FILTER_1=""                                  # Дополнительный фильтр для аудио
AUDIO_PROXY_FILTER_1=""                            # Дополнительный фильтр для proxy аудио

append() {
    if [[ $# -lt 2 ]]; then
        echo "Ошибка: функция ожидает минимум 2 аргумента - имя массива и строки" >&2
        return 1
    fi

    local -n arr_ref="$1"
    shift  # Убираем имя массива из списка аргументов
    
    local value
    for value in "$@"; do
        # Проверяем, что строка не пустая
        if [[ -n "$value" ]]; then
            arr_ref+=("$value")
        fi
    done
}

join_array() {
    if [[ $# -lt 2 ]]; then
        echo "Ошибка: ожидается имя массива и разделитель" >&2
        return 1
    fi
    
    local -n arr_ref="$1"
    local delimiter="$2"
    
    if [[ ${#arr_ref[@]} -eq 0 ]]; then
        echo ""
        return 0
    fi
    
    if [[ ${#arr_ref[@]} -eq 1 ]]; then
        echo "${arr_ref[0]}"
        return 0
    fi
    
    local first_element="${arr_ref[0]}"
    printf "%s" "$first_element"
    
    for element in "${arr_ref[@]:1}"; do
        printf "%s%s" "$delimiter" "$element"
    done
    
    echo ""
}

# ============================================================================
# ФУНКЦИИ ПОСТРОЕНИЯ VIDEO ЧАСТИ FILTER_COMPLEX
# ============================================================================

build_filter_complex_video_filter() {
    local index="$1"
    
    if [[ -z "$index" ]]; then
        echo "Ошибка: функция требует один аргумент (индекс видеопотока)" >&2
        return 1
    fi
    
    if ! [[ "$index" =~ ^[0-9]+$ ]]; then
        echo "Ошибка: индекс должен быть числом, получено: '$index'" >&2
        return 1
    fi
    
    if [[ "${VIDEO_ENABLED}" != "true" ]] || [[ -z "${VIDEO_FILTER_1}" ]]; then
        echo ""
        return 0
    fi
    
    local video_filter=""
    video_filter+="[0:v:${index}]"
    video_filter+="${VIDEO_FILTER_1}"
    
    if [[ "${VIDEO_PROXY_ENABLED}" == "true" ]] && [[ -n "${VIDEO_PROXY_FILTER_1}" ]]; then
        video_filter+=",split=2"
        video_filter+="[video_${index}][in_proxy_${index}]"
        video_filter+=";"
        video_filter+="[in_proxy_${index}]"
        video_filter+="${VIDEO_PROXY_FILTER_1}"
        video_filter+="[video_${index}_proxy]"
    else
        video_filter+="[video_${index}]"
    fi
    
    echo "$video_filter"
}

build_filter_complex_videos() {
    if [[ $# -ne 1 ]]; then
        echo "Ошибка: функция требует один аргумент (имя массива с индексами)" >&2
        return 1
    fi
    
    local -n index_array_ref="$1"
    local -a video_filters_array=()
    local video_filter
    local result=""
    
    for index in "${index_array_ref[@]}"; do
        video_filter=$(build_filter_complex_video_filter "$index")
        
        if [[ -n "$video_filter" ]]; then
            append video_filters_array "$video_filter"
        fi
    done
    
    if [[ ${#video_filters_array[@]} -gt 0 ]]; then
        result=$(join_array video_filters_array ";")
    fi
    
    echo "$result"
}

# ============================================================================
# ФУНКЦИИ ПОСТРОЕНИЯ AUDIO ЧАСТИ FILTER_COMPLEX
# ============================================================================

build_filter_complex_audio_filter() {
    local index="$1"
    
    if [[ -z "$index" ]]; then
        echo "Ошибка: функция требует один аргумент (индекс аудиопотока)" >&2
        return 1
    fi
    
    if ! [[ "$index" =~ ^[0-9]+$ ]]; then
        echo "Ошибка: индекс должен быть числом, получено: '$index'" >&2
        return 1
    fi
    
    if [[ "${AUDIO_ENABLED}" != "true" ]] || [[ -z "${AUDIO_FILTER_1}" ]]; then
        echo ""
        return 0
    fi
    
    local audio_filter=""
    audio_filter+="[0:a:${index}]"
    audio_filter+="${AUDIO_FILTER_1}"
    
    if [[ "${AUDIO_PROXY_ENABLED}" == "true" ]] && [[ -n "${AUDIO_PROXY_FILTER_1}" ]]; then
        audio_filter+=",asplit=2"
        audio_filter+="[audio_${index}][audio_${index}_pr_raw]"
        audio_filter+=";"
        audio_filter+="[audio_${index}_pr_raw]"
        audio_filter+="${AUDIO_PROXY_FILTER_1}"
        audio_filter+="[audio_${index}_proxy]"
    else
        audio_filter+="[audio_${index}]"
    fi
    
    echo "$audio_filter"
}

build_filter_complex_audios() {
    if [[ $# -ne 1 ]]; then
        echo "Ошибка: функция требует один аргумент (имя массива с индексами)" >&2
        return 1
    fi
    
    local -n index_array_ref="$1"
    local -a audio_filters_array=()
    local audio_filter
    local result=""
    
    for index in "${index_array_ref[@]}"; do
        audio_filter=$(build_filter_complex_audio_filter "$index")
        
        if [[ -n "$audio_filter" ]]; then
            append audio_filters_array "$audio_filter"
        fi
    done
    
    if [[ ${#audio_filters_array[@]} -gt 0 ]]; then
        result=$(join_array audio_filters_array ";")
    fi
    
    echo "$result"
}

# ============================================================================
# ФУНКЦИЯ ПОСТРОЕНИЯ ПОЛНОГО FILTER_COMPLEX
# ============================================================================

build_filter_complex() {
    if [[ $# -ne 2 ]]; then
        echo "Ошибка: функция требует два аргумента (имена массивов с видео и аудио индексами)" >&2
        return 1
    fi
    
    local -n video_indices_ref="$1"
    local -n audio_indices_ref="$2"
    
    local video_part=""
    local audio_part=""
    local result=""
    
    if [[ ${#video_indices_ref[@]} -gt 0 ]]; then
        video_part=$(build_filter_complex_videos "$1")
    fi
    
    if [[ ${#audio_indices_ref[@]} -gt 0 ]]; then
        audio_part=$(build_filter_complex_audios "$2")
    fi
    
    if [[ -n "$video_part" ]] && [[ -n "$audio_part" ]]; then
        result="${video_part};${audio_part}"
    elif [[ -n "$video_part" ]]; then
        result="$video_part"
    elif [[ -n "$audio_part" ]]; then
        result="$audio_part"
    fi
    
    echo "$result"
}

# ============================================================================
# ТЕСТОВЫЕ НАЗНАЧЕНИЯ ПЕРЕМЕННЫХ И ВЫЗОВЫ
# ============================================================================

echo "=== Тест 1: Видео и аудио с proxy ==="
VIDEO_ENABLED="true"
VIDEO_PROXY_ENABLED="true"
VIDEO_FILTER_1="yadif=0:-1:0"
VIDEO_PROXY_FILTER_1="scale=iw/2:ih,setsar=(1/1)*2"

AUDIO_ENABLED="true"
AUDIO_PROXY_ENABLED="true"
AUDIO_FILTER_1="aresample=48000,atempo=25/24"
AUDIO_PROXY_FILTER_1="volume=0.8,highpass=f=80"

declare -a VIDEO_INDICES=(0)
declare -a AUDIO_INDICES=(0 1)

echo "Видео параметры:"
echo "  VIDEO_ENABLED=${VIDEO_ENABLED}"
echo "  VIDEO_PROXY_ENABLED=${VIDEO_PROXY_ENABLED}"
echo "  VIDEO_FILTER_1=${VIDEO_FILTER_1}"
echo "  VIDEO_PROXY_FILTER_1=${VIDEO_PROXY_FILTER_1}"
echo "Аудио параметры:"
echo "  AUDIO_ENABLED=${AUDIO_ENABLED}"
echo "  AUDIO_PROXY_ENABLED=${AUDIO_PROXY_ENABLED}"
echo "  AUDIO_FILTER_1=${AUDIO_FILTER_1}"
echo "  AUDIO_PROXY_FILTER_1=${AUDIO_PROXY_FILTER_1}"
echo "Индексы:"
echo "  VIDEO_INDICES: ${VIDEO_INDICES[@]}"
echo "  AUDIO_INDICES: ${AUDIO_INDICES[@]}"

result=$(build_filter_complex VIDEO_INDICES AUDIO_INDICES)
echo "Результат build_filter_complex:"
echo "$result"
echo ""

echo "=== Тест 2: Только видео (аудио отключено) ==="
VIDEO_ENABLED="true"
VIDEO_PROXY_ENABLED="true"
VIDEO_FILTER_1="yadif=0:-1:0,scale=1920:1080"
VIDEO_PROXY_FILTER_1="scale=iw/4:ih,setsar=(1/1)*4"

AUDIO_ENABLED="false"  # Аудио отключено
AUDIO_PROXY_ENABLED="true"
AUDIO_FILTER_1="aresample=48000"
AUDIO_PROXY_FILTER_1="volume=0.5"

declare -a VIDEO_INDICES_2=(0 1)
declare -a AUDIO_INDICES_2=(0)

echo "Видео параметры:"
echo "  VIDEO_ENABLED=${VIDEO_ENABLED}"
echo "  VIDEO_PROXY_ENABLED=${VIDEO_PROXY_ENABLED}"
echo "  VIDEO_FILTER_1=${VIDEO_FILTER_1}"
echo "  VIDEO_PROXY_FILTER_1=${VIDEO_PROXY_FILTER_1}"
echo "Аудио параметры:"
echo "  AUDIO_ENABLED=${AUDIO_ENABLED}"
echo "  AUDIO_PROXY_ENABLED=${AUDIO_PROXY_ENABLED}"
echo "  AUDIO_FILTER_1=${AUDIO_FILTER_1}"
echo "  AUDIO_PROXY_FILTER_1=${AUDIO_PROXY_FILTER_1}"
echo "Индексы:"
echo "  VIDEO_INDICES: ${VIDEO_INDICES_2[@]}"
echo "  AUDIO_INDICES: ${AUDIO_INDICES_2[@]}"

result2=$(build_filter_complex VIDEO_INDICES_2 AUDIO_INDICES_2)
echo "Результат build_filter_complex:"
echo "$result2"
echo ""

echo "=== Тест 3: Видео без proxy, аудио с proxy ==="
VIDEO_ENABLED="true"
VIDEO_PROXY_ENABLED="false"  # Proxy видео отключено
VIDEO_FILTER_1="hqdn3d=1.5,eq=brightness=0.1"
VIDEO_PROXY_FILTER_1="scale=iw/2:ih"

AUDIO_ENABLED="true"
AUDIO_PROXY_ENABLED="true"
AUDIO_FILTER_1="aresample=44100,atempo=1.04"
AUDIO_PROXY_FILTER_1="volume=0.7,lowpass=f=8000"

declare -a VIDEO_INDICES_3=(0)
declare -a AUDIO_INDICES_3=(0 1 2)

echo "Видео параметры:"
echo "  VIDEO_ENABLED=${VIDEO_ENABLED}"
echo "  VIDEO_PROXY_ENABLED=${VIDEO_PROXY_ENABLED}"
echo "  VIDEO_FILTER_1=${VIDEO_FILTER_1}"
echo "  VIDEO_PROXY_FILTER_1=${VIDEO_PROXY_FILTER_1}"
echo "Аудио параметры:"
echo "  AUDIO_ENABLED=${AUDIO_ENABLED}"
echo "  AUDIO_PROXY_ENABLED=${AUDIO_PROXY_ENABLED}"
echo "  AUDIO_FILTER_1=${AUDIO_FILTER_1}"
echo "  AUDIO_PROXY_FILTER_1=${AUDIO_PROXY_FILTER_1}"
echo "Индексы:"
echo "  VIDEO_INDICES: ${VIDEO_INDICES_3[@]}"
echo "  AUDIO_INDICES: ${AUDIO_INDICES_3[@]}"

result3=$(build_filter_complex VIDEO_INDICES_3 AUDIO_INDICES_3)
echo "Результат build_filter_complex:"
echo "$result3"
echo ""

echo "=== Тест 4: Отдельные функции видео и аудио ==="
VIDEO_ENABLED="true"
VIDEO_PROXY_ENABLED="true"
VIDEO_FILTER_1="yadif=1:-1:0"
VIDEO_PROXY_FILTER_1="scale=640:360"

declare -a VID_INDICES=(0 1)
echo "Тестируем build_filter_complex_videos:"
video_result=$(build_filter_complex_videos VID_INDICES)
echo "$video_result"
echo ""

AUDIO_ENABLED="true"
AUDIO_PROXY_ENABLED="false"
AUDIO_FILTER_1="aresample=48000"

declare -a AUD_INDICES=(0)
echo "Тестируем build_filter_complex_audios:"
audio_result=$(build_filter_complex_audios AUD_INDICES)
echo "$audio_result"
echo ""

echo "=== Тест 5: Пустые результаты ==="
VIDEO_ENABLED="false"
AUDIO_ENABLED="false"

declare -a EMPTY_VID_INDICES=(0)
declare -a EMPTY_AUD_INDICES=(0 1)

empty_result=$(build_filter_complex EMPTY_VID_INDICES EMPTY_AUD_INDICES)
echo "Результат при отключенном видео и аудио: '$empty_result'"
echo ""

echo "=== Тест 6: Ошибки аргументов ==="
echo "Тестируем build_filter_complex_video_filter без аргументов:"
build_filter_complex_video_filter
echo ""

echo "Тестируем build_filter_complex_videos без аргументов:"
build_filter_complex_videos
echo ""

echo "Тестируем build_filter_complex без аргументов:"
build_filter_complex
```