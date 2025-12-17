---
updated_at: 2025-12-17T17:56:05.373+10:00
---
```bash
#!/bin/bash

# Основные параметры обработки
PRIORITY=5                                         # Приоритет обработки
INPUT_FILE="_____________"                         # Имя исходного файла
OUTPUT_BASE="__________"                           # Базовое имя выходных файлов
REVISION_SUFFIX="_HD"                              # Суффикс для версии файла

# Базовые пути
HOME_ROOT="/home/pemaltynov"                       # Корневая домашняя директория
MOUNT_ROOT="/mnt/pemaltynov"                       # Корневая точка монтирования
ROOT_SUBPATH="ROOT"                                # Общий подпуть для архивных путей
EDIT_SUBPATH="EDIT/_______"                        # Подпуть для редактирования
ARCHIVE_SUBPATH="IN/_______"                       # Подпуть для архивации

# Параметры аудио
AUDIO_SAMPLE_RATE=48000                            # Частота дискретизации аудио (Гц)
AUDIO_TEMPO_DENOMINATOR="__"                       # Знаменатель для atempo (25/__)

# Параметры видео
VIDEO_FRAMERATE=25                                 # Частота кадров видео (fps)
VIDEO_SCALE_DIVISOR=2                              # Делитель масштаба для proxy видео
VIDEO_SAR_FACTOR=2                                 # Множитель для SAR proxy видео

# Параметры кодирования HD видео
VIDEO_HD_CODEC="libx264"                           # Кодек для HD видео
VIDEO_HD_PRESET="medium"                           # Пресет кодирования HD
VIDEO_HD_CRF=16                                    # CRF значение для HD
VIDEO_HD_PIX_FMT="yuv420p"                         # Пиксельный формат HD
VIDEO_HD_GOP_SIZE=0                                # Размер GOP (0 - только ключевые кадры)

# Параметры кодирования proxy видео
VIDEO_PROXY_CODEC="libx264"                        # Кодек для proxy видео
VIDEO_PROXY_PRESET="superfast"                     # Пресет кодирования proxy
VIDEO_PROXY_PIX_FMT="yuv420p"                      # Пиксельный формат proxy
VIDEO_PROXY_BITRATE="2000k"                        # Битрейт для proxy видео
VIDEO_PROXY_MAXRATE="2000k"                        # Максимальный битрейт proxy
VIDEO_PROXY_INTERLACED=1                           # Флаг чересстрочности proxy

# Параметры кодирования HD аудио
AUDIO_HD_CODEC="alac"                              # Кодек для HD аудио
AUDIO_HD_COMPRESSION_LEVEL=0                       # Уровень сжатия ALAC

# Параметры кодирования proxy аудио
AUDIO_PROXY_CODEC="ac3"                            # Кодек для proxy аудио
AUDIO_PROXY_BITRATE="128k"                         # Битрейт для proxy аудио

# Таймауты и задержки
ARCHIVE_DELAY="10 hours"                           # Задержка перед архивацией

# Составные пути
# Входные директории
INPUT_DIR="${HOME_ROOT}/IN"
PROGRESS_DIR="${INPUT_DIR}/_IN_PROGRESS"
NOTIFICATION_DIR="${INPUT_DIR}/notifications"
DONE_DIR="${INPUT_DIR}/_DONE"
SCRIPT_ARCHIVE_DIR="${DONE_DIR}/bash"

# Редактирование и архивация
EDIT_BASE_PATH="${MOUNT_ROOT}/${ROOT_SUBPATH}/${EDIT_SUBPATH}"
ARCHIVE_BASE_PATH="${MOUNT_ROOT}/${ROOT_SUBPATH}/${ARCHIVE_SUBPATH}"

# Создание путей
ARCHIVE_OUTPUT_PATH="${ARCHIVE_BASE_PATH}/_DONE/${OUTPUT_BASE}"
FULL_EDIT_PATH="${EDIT_BASE_PATH}"
FULL_NOTIFICATION_PATH="${NOTIFICATION_DIR}/${OUTPUT_BASE}.done"
FULL_DONE_PATH="${DONE_DIR}"
READY_FILE_PATH="${FULL_EDIT_PATH}/${OUTPUT_BASE}.ready"

# ============================================================================
# ФУНКЦИИ ПОСТРОЕНИЯ БЛОКОВ КОМАНДЫ FFMEPG
# ============================================================================

# Функция для построения глобальных параметров
build_global_parameters() {
    local global_params="-r ${VIDEO_FRAMERATE}"
    echo "$global_params"
}

# Функция для построения блока входных файлов
build_inputs() {
    local inputs="-i \"${PROGRESS_DIR}/${INPUT_FILE}\""
    echo "$inputs"
}

# Функция для построения видеофильтров (часть filter_complex)
build_video_filters() {
    local video_filters="[0:v:0]split=2[vidHD][inProxy]; "
    video_filters+="[inProxy]scale=iw/${VIDEO_SCALE_DIVISOR}:ih, setsar=(1/1)*${VIDEO_SAR_FACTOR}[vidHD_pr]"
    echo "$video_filters"
}

# Функция для построения аудиофильтра для одной дорожки
build_audio_filter_for_track() {
    local track_index=$1
    local audio_filter_base="aresample=${AUDIO_SAMPLE_RATE},atempo=${VIDEO_FRAMERATE}/(${AUDIO_TEMPO_DENOMINATOR})"
    
    local audio_filter="[0:a:${track_index}]${audio_filter_base}[audio_in_${track_index}]; "
    audio_filter+="[audio_in_${track_index}]asplit=2[audio_${track_index}][audio_${track_index}_pr]"
    
    echo "$audio_filter"
}

# Функция для построения всех аудиофильтров
build_all_audio_filters() {
    local audio_count=$1
    shift
    local audio_suffixes=("$@")
    
    if [ "$audio_count" -eq 0 ]; then
        echo ""
        return 0
    fi
    
    local all_audio_filters=""
    
    for i in $(seq 0 $((audio_count - 1))); do
        all_audio_filters+=$(build_audio_filter_for_track "$i")
        all_audio_filters+="; "
    done
    
    # Удаляем последний разделитель
    all_audio_filters="${all_audio_filters%; }"
    
    echo "$all_audio_filters"
}

# Функция для построения всего filter_complex
build_filter_complex() {
    local audio_count=$1
    shift
    local audio_suffixes=("$@")
    
    local filter_complex=$(build_video_filters)
    local audio_filters=$(build_all_audio_filters "$audio_count" "${audio_suffixes[@]}")
    
    if [ -n "$audio_filters" ]; then
        filter_complex+="; "
        filter_complex+="$audio_filters"
    fi
    
    echo "$filter_complex"
}

# Функция для построения карты HD видео
build_video_hd_map() {
    local map="-map \"[vidHD]\" "
    map+="-c:v ${VIDEO_HD_CODEC} "
    map+="-preset ${VIDEO_HD_PRESET} "
    map+="-crf ${VIDEO_HD_CRF} "
    map+="-pix_fmt ${VIDEO_HD_PIX_FMT} "
    map+="-g ${VIDEO_HD_GOP_SIZE} "
    map+="-map_metadata -1 "
    map+="-map_chapters -1 "
    map+="\"${FULL_EDIT_PATH}/${OUTPUT_BASE}${REVISION_SUFFIX}.mp4\""
    
    echo "$map"
}

# Функция для построения карты proxy видео
build_video_proxy_map() {
    local map="-map \"[vidHD_pr]\" "
    map+="-c:v ${VIDEO_PROXY_CODEC} "
    map+="-x264opts interlaced=${VIDEO_PROXY_INTERLACED} "
    map+="-preset ${VIDEO_PROXY_PRESET} "
    map+="-pix_fmt ${VIDEO_PROXY_PIX_FMT} "
    map+="-b:v ${VIDEO_PROXY_BITRATE} "
    map+="-maxrate ${VIDEO_PROXY_MAXRATE} "
    map+="-map_metadata -1 "
    map+="-map_chapters -1 "
    map+="\"${FULL_EDIT_PATH}/${OUTPUT_BASE}${REVISION_SUFFIX}_proxy.mp4\""
    
    echo "$map"
}

# Функция для построения карты HD аудио для одной дорожки
build_audio_hd_map_for_track() {
    local track_index=$1
    local audio_suffix=$2
    
    local map="-map \"[audio_${track_index}]\" "
    map+="-c:a ${AUDIO_HD_CODEC} "
    map+="-compression_level ${AUDIO_HD_COMPRESSION_LEVEL} "
    map+="-map_metadata -1 "
    map+="-map_chapters -1 "
    map+="\"${FULL_EDIT_PATH}/${OUTPUT_BASE}${REVISION_SUFFIX}_${audio_suffix}.m4a\""
    
    echo "$map"
}

# Функция для построения карты proxy аудио для одной дорожки
build_audio_proxy_map_for_track() {
    local track_index=$1
    local audio_suffix=$2
    
    local map="-map \"[audio_${track_index}_pr]\" "
    map+="-c:a ${AUDIO_PROXY_CODEC} "
    map+="-b:a ${AUDIO_PROXY_BITRATE} "
    map+="\"${FULL_EDIT_PATH}/${OUTPUT_BASE}${REVISION_SUFFIX}_${audio_suffix}_proxy.ac3\""
    
    echo "$map"
}

# Функция для построения всех аудио карт (HD + proxy для каждой дорожки)
build_all_audio_maps() {
    local audio_count=$1
    shift
    local audio_suffixes=("$@")
    
    if [ "$audio_count" -eq 0 ]; then
        echo ""
        return 0
    fi
    
    local all_audio_maps=""
    
    for i in $(seq 0 $((audio_count - 1))); do
        all_audio_maps+=$(build_audio_hd_map_for_track "$i" "${audio_suffixes[i]}")
        all_audio_maps+=" "
        all_audio_maps+=$(build_audio_proxy_map_for_track "$i" "${audio_suffixes[i]}")
        all_audio_maps+=" "
    done
    
    # Удаляем последний пробел
    all_audio_maps="${all_audio_maps% }"
    
    echo "$all_audio_maps"
}

# Функция для построения всех видео карт (HD + proxy)
build_all_video_maps() {
    local video_maps=$(build_video_hd_map)
    video_maps+=" "
    video_maps+=$(build_video_proxy_map)
    
    echo "$video_maps"
}

# Функция для построения всех карт вывода (видео + аудио)
build_all_maps() {
    local audio_count=$1
    shift
    local audio_suffixes=("$@")
    
    local all_maps=$(build_all_video_maps)
    local audio_maps=$(build_all_audio_maps "$audio_count" "${audio_suffixes[@]}")
    
    if [ -n "$audio_maps" ]; then
        all_maps+=" "
        all_maps+="$audio_maps"
    fi
    
    echo "$all_maps"
}

# Функция для построения полной команды ffmpeg
build_ffmpeg_command() {
    local audio_count=$1
    shift
    local audio_suffixes=("$@")
    
    local global_params=$(build_global_parameters)
    local inputs=$(build_inputs)
    local filter_complex=$(build_filter_complex "$audio_count" "${audio_suffixes[@]}")
    local maps=$(build_all_maps "$audio_count" "${audio_suffixes[@]}")
    
    local command="fflite ${global_params} ${inputs} "
    command+="-filter_complex \"${filter_complex}\" "
    command+="${maps}"
    
    echo "$command"
}

# ============================================================================
# ОСНОВНОЙ КОД СКРИПТА
# ============================================================================

# Создание необходимых директорий
mkdir -p "${ARCHIVE_OUTPUT_PATH}"
mkdir -p "${FULL_EDIT_PATH}"

# Очистка экрана и перемещение входного файла
clear
mv "${INPUT_DIR}/${INPUT_FILE}" "${PROGRESS_DIR}/"

# Определение аудиодорожек с помощью внешней утилиты
AUDIO_SUFFIXES=($(defineAudioStreams "${PROGRESS_DIR}/${INPUT_FILE}"))
AUDIO_TRACK_COUNT=${#AUDIO_SUFFIXES[@]}

echo "Найдено аудиодорожек: ${AUDIO_TRACK_COUNT}"
if [ ${AUDIO_TRACK_COUNT} -gt 0 ]; then
    echo "Суффиксы аудиодорожек: ${AUDIO_SUFFIXES[@]}"
fi

# Построение и выполнение команды ffmpeg
echo "Построение команды ffmpeg..."
FFMPEG_CMD=$(build_ffmpeg_command "$AUDIO_TRACK_COUNT" "${AUDIO_SUFFIXES[@]}")
echo "Выполнение команды:"
echo "${FFMPEG_CMD}"
echo ""

eval ${FFMPEG_CMD} && \
  
# Создание файла готовности
touch "${READY_FILE_PATH}" && \
  
# Создание файла уведомления
echo "${EDIT_SUBPATH}/${OUTPUT_BASE}.ready" > "${FULL_NOTIFICATION_PATH}" && \
  
# Перемещение обработанного файла во временную директорию
mv "${PROGRESS_DIR}/${INPUT_FILE}" "${FULL_DONE_PATH}/" && \
  
# Отложенное перемещение файла в архив
at now + ${ARCHIVE_DELAY} <<< "mv ${FULL_DONE_PATH}/${INPUT_FILE} ${ARCHIVE_OUTPUT_PATH}" && \
  
# Очистка и архивация скрипта
clear
mv "$0" "${SCRIPT_ARCHIVE_DIR}/"
```

---

```bash
#!/bin/bash

# Основные параметры обработки
PRIORITY=5                                         # Приоритет обработки
INPUT_FILE="_____________"                         # Имя исходного файла
OUTPUT_BASE="__________"                           # Базовое имя выходных файлов
REVISION_SUFFIX="_HD"                              # Суффикс для версии файла

# Базовые пути
HOME_ROOT="/home/pemaltynov"                       # Корневая домашняя директория
MOUNT_ROOT="/mnt/pemaltynov"                       # Корневая точка монтирования
ROOT_SUBPATH="ROOT"                                # Общий подпуть для архивных путей
EDIT_SUBPATH="EDIT/_______"                        # Подпуть для редактирования
ARCHIVE_SUBPATH="IN/_______"                       # Подпуть для архивации

# Параметры аудио
AUDIO_SAMPLE_RATE=48000                            # Частота дискретизации аудио (Гц)
AUDIO_TEMPO_DENOMINATOR="__"                       # Знаменатель для atempo (25/__)

# Параметры видео
VIDEO_FRAMERATE=25                                 # Частота кадров видео (fps)
VIDEO_SCALE_DIVISOR=2                              # Делитель масштаба для proxy видео
VIDEO_SAR_FACTOR=2                                 # Множитель для SAR proxy видео

# Параметры кодирования HD видео
VIDEO_HD_CODEC="libx264"                           # Кодек для HD видео
VIDEO_HD_PRESET="medium"                           # Пресет кодирования HD
VIDEO_HD_CRF=16                                    # CRF значение для HD
VIDEO_HD_PIX_FMT="yuv420p"                         # Пиксельный формат HD
VIDEO_HD_GOP_SIZE=0                                # Размер GOP (0 - только ключевые кадры)

# Параметры кодирования proxy видео
VIDEO_PROXY_CODEC="libx264"                        # Кодек для proxy видео
VIDEO_PROXY_PRESET="superfast"                     # Пресет кодирования proxy
VIDEO_PROXY_PIX_FMT="yuv420p"                      # Пиксельный формат proxy
VIDEO_PROXY_BITRATE="2000k"                        # Битрейт для proxy видео
VIDEO_PROXY_MAXRATE="2000k"                        # Максимальный битрейт proxy
VIDEO_PROXY_INTERLACED=1                           # Флаг чересстрочности proxy

# Параметры кодирования HD аудио
AUDIO_HD_CODEC="alac"                              # Кодек для HD аудио
AUDIO_HD_COMPRESSION_LEVEL=0                       # Уровень сжатия ALAC

# Параметры кодирования proxy аудио
AUDIO_PROXY_CODEC="ac3"                            # Кодек для proxy аудио
AUDIO_PROXY_BITRATE="128k"                         # Битрейт для proxy аудио

# Таймауты и задержки
ARCHIVE_DELAY="10 hours"                           # Задержка перед архивацией

# Составные пути
# Входные директории
INPUT_DIR="${HOME_ROOT}/IN"
PROGRESS_DIR="${INPUT_DIR}/_IN_PROGRESS"
NOTIFICATION_DIR="${INPUT_DIR}/notifications"
DONE_DIR="${INPUT_DIR}/_DONE"
SCRIPT_ARCHIVE_DIR="${DONE_DIR}/bash"

# Редактирование и архивация
EDIT_BASE_PATH="${MOUNT_ROOT}/${ROOT_SUBPATH}/${EDIT_SUBPATH}"
ARCHIVE_BASE_PATH="${MOUNT_ROOT}/${ROOT_SUBPATH}/${ARCHIVE_SUBPATH}"

# Создание путей
ARCHIVE_OUTPUT_PATH="${ARCHIVE_BASE_PATH}/_DONE/${OUTPUT_BASE}"
FULL_EDIT_PATH="${EDIT_BASE_PATH}"
FULL_NOTIFICATION_PATH="${NOTIFICATION_DIR}/${OUTPUT_BASE}.done"
FULL_DONE_PATH="${DONE_DIR}"
READY_FILE_PATH="${FULL_EDIT_PATH}/${OUTPUT_BASE}.ready"

# ============================================================================
# ФУНКЦИИ ДЛЯ РАБОТЫ С МАССИВАМИ И СТРОКАМИ
# ============================================================================

# Функция для добавления строки в конец массива
append() {
    if [[ $# -ne 2 ]]; then
        echo "Ошибка: функция ожидает 2 аргумента - имя массива и строку" >&2
        return 1
    fi

    local -n arr_ref="$1"
    local value="$2"
    
    arr_ref+=("$value")
}

# Функция для разделения строки на массив по разделителю
split_string_to_array() {
    if [[ $# -ne 3 ]]; then
        echo "Ошибка: ожидается имя массива, строка и разделитель" >&2
        return 1
    fi
    
    local -n arr_ref="$1"
    local string="$2"
    local delimiter="$3"
    
    arr_ref=()
    
    if [[ -z "$string" ]]; then
        return 0
    fi
    
    if [[ -z "$delimiter" ]]; then
        arr_ref+=("$string")
        return 0
    fi
    
    if [[ "$string" != *"$delimiter"* ]]; then
        arr_ref+=("$string")
        return 0
    fi
    
    local rest="$string"
    
    while [[ "$rest" == *"$delimiter"* ]]; do
        local part="${rest%%"$delimiter"*}"
        arr_ref+=("$part")
        rest="${rest#*"$delimiter"}"
        
        if [[ -z "$rest" ]]; then
            arr_ref+=("")
            break
        fi
    done
    
    if [[ -n "$rest" ]]; then
        arr_ref+=("$rest")
    fi
}

# Функция для соединения элементов массива с разделителем
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
# ФУНКЦИИ ПОСТРОЕНИЯ КОМАНДЫ FFMPEG В ВИДЕ МАССИВА
# ============================================================================

# Глобальный массив для команды ffmpeg
declare -a FFMPEG_CMD_ARGS

# Функция для добавления глобальных параметров
add_global_parameters() {
    append FFMPEG_CMD_ARGS "-r"
    append FFMPEG_CMD_ARGS "${VIDEO_FRAMERATE}"
}

# Функция для добавления входных файлов
add_inputs() {
    append FFMPEG_CMD_ARGS "-i"
    append FFMPEG_CMD_ARGS "${PROGRESS_DIR}/${INPUT_FILE}"
}

# Функция для построения видеофильтров
build_video_filters() {
    local video_filters="[0:v:0]split=2[vidHD][inProxy]; "
    video_filters+="[inProxy]scale=iw/${VIDEO_SCALE_DIVISOR}:ih, setsar=(1/1)*${VIDEO_SAR_FACTOR}[vidHD_pr]"
    echo "$video_filters"
}

# Функция для построения аудиофильтра для одной дорожки
build_audio_filter_for_track() {
    local track_index="$1"
    local audio_filter_base="aresample=${AUDIO_SAMPLE_RATE},atempo=${VIDEO_FRAMERATE}/(${AUDIO_TEMPO_DENOMINATOR})"
    
    local audio_filter="[0:a:${track_index}]${audio_filter_base}[audio_in_${track_index}]; "
    audio_filter+="[audio_in_${track_index}]asplit=2[audio_${track_index}][audio_${track_index}_pr]"
    
    echo "$audio_filter"
}

# Функция для построения всех аудиофильтров
build_all_audio_filters() {
    local audio_count="$1"
    shift
    local audio_suffixes=("$@")
    
    if [[ "$audio_count" -eq 0 ]]; then
        echo ""
        return 0
    fi
    
    local all_audio_filters=""
    
    for i in $(seq 0 $((audio_count - 1))); do
        all_audio_filters+=$(build_audio_filter_for_track "$i")
        all_audio_filters+="; "
    done
    
    all_audio_filters="${all_audio_filters%; }"
    
    echo "$all_audio_filters"
}

# Функция для построения всего filter_complex
build_filter_complex() {
    local audio_count="$1"
    shift
    local audio_suffixes=("$@")
    
    local filter_complex
    filter_complex=$(build_video_filters)
    local audio_filters
    audio_filters=$(build_all_audio_filters "$audio_count" "${audio_suffixes[@]}")
    
    if [[ -n "$audio_filters" ]]; then
        filter_complex+="; "
        filter_complex+="$audio_filters"
    fi
    
    echo "$filter_complex"
}

# Функция для добавления фильтров в команду
add_filter_complex() {
    local audio_count="$1"
    shift
    local audio_suffixes=("$@")
    
    local filter_complex
    filter_complex=$(build_filter_complex "$audio_count" "${audio_suffixes[@]}")
    
    append FFMPEG_CMD_ARGS "-filter_complex"
    append FFMPEG_CMD_ARGS "$filter_complex"
}

# Функция для добавления карты HD видео
add_video_hd_map() {
    append FFMPEG_CMD_ARGS "-map"
    append FFMPEG_CMD_ARGS "[vidHD]"
    append FFMPEG_CMD_ARGS "-c:v"
    append FFMPEG_CMD_ARGS "${VIDEO_HD_CODEC}"
    append FFMPEG_CMD_ARGS "-preset"
    append FFMPEG_CMD_ARGS "${VIDEO_HD_PRESET}"
    append FFMPEG_CMD_ARGS "-crf"
    append FFMPEG_CMD_ARGS "${VIDEO_HD_CRF}"
    append FFMPEG_CMD_ARGS "-pix_fmt"
    append FFMPEG_CMD_ARGS "${VIDEO_HD_PIX_FMT}"
    append FFMPEG_CMD_ARGS "-g"
    append FFMPEG_CMD_ARGS "${VIDEO_HD_GOP_SIZE}"
    append FFMPEG_CMD_ARGS "-map_metadata"
    append FFMPEG_CMD_ARGS "-1"
    append FFMPEG_CMD_ARGS "-map_chapters"
    append FFMPEG_CMD_ARGS "-1"
    append FFMPEG_CMD_ARGS "${FULL_EDIT_PATH}/${OUTPUT_BASE}${REVISION_SUFFIX}.mp4"
}

# Функция для добавления карты proxy видео
add_video_proxy_map() {
    append FFMPEG_CMD_ARGS "-map"
    append FFMPEG_CMD_ARGS "[vidHD_pr]"
    append FFMPEG_CMD_ARGS "-c:v"
    append FFMPEG_CMD_ARGS "${VIDEO_PROXY_CODEC}"
    append FFMPEG_CMD_ARGS "-x264opts"
    append FFMPEG_CMD_ARGS "interlaced=${VIDEO_PROXY_INTERLACED}"
    append FFMPEG_CMD_ARGS "-preset"
    append FFMPEG_CMD_ARGS "${VIDEO_PROXY_PRESET}"
    append FFMPEG_CMD_ARGS "-pix_fmt"
    append FFMPEG_CMD_ARGS "${VIDEO_PROXY_PIX_FMT}"
    append FFMPEG_CMD_ARGS "-b:v"
    append FFMPEG_CMD_ARGS "${VIDEO_PROXY_BITRATE}"
    append FFMPEG_CMD_ARGS "-maxrate"
    append FFMPEG_CMD_ARGS "${VIDEO_PROXY_MAXRATE}"
    append FFMPEG_CMD_ARGS "-map_metadata"
    append FFMPEG_CMD_ARGS "-1"
    append FFMPEG_CMD_ARGS "-map_chapters"
    append FFMPEG_CMD_ARGS "-1"
    append FFMPEG_CMD_ARGS "${FULL_EDIT_PATH}/${OUTPUT_BASE}${REVISION_SUFFIX}_proxy.mp4"
}

# Функция для добавления карты HD аудио для одной дорожки
add_audio_hd_map_for_track() {
    local track_index="$1"
    local audio_suffix="$2"
    
    append FFMPEG_CMD_ARGS "-map"
    append FFMPEG_CMD_ARGS "[audio_${track_index}]"
    append FFMPEG_CMD_ARGS "-c:a"
    append FFMPEG_CMD_ARGS "${AUDIO_HD_CODEC}"
    append FFMPEG_CMD_ARGS "-compression_level"
    append FFMPEG_CMD_ARGS "${AUDIO_HD_COMPRESSION_LEVEL}"
    append FFMPEG_CMD_ARGS "-map_metadata"
    append FFMPEG_CMD_ARGS "-1"
    append FFMPEG_CMD_ARGS "-map_chapters"
    append FFMPEG_CMD_ARGS "-1"
    append FFMPEG_CMD_ARGS "${FULL_EDIT_PATH}/${OUTPUT_BASE}${REVISION_SUFFIX}_${audio_suffix}.m4a"
}

# Функция для добавления карты proxy аудио для одной дорожки
add_audio_proxy_map_for_track() {
    local track_index="$1"
    local audio_suffix="$2"
    
    append FFMPEG_CMD_ARGS "-map"
    append FFMPEG_CMD_ARGS "[audio_${track_index}_pr]"
    append FFMPEG_CMD_ARGS "-c:a"
    append FFMPEG_CMD_ARGS "${AUDIO_PROXY_CODEC}"
    append FFMPEG_CMD_ARGS "-b:a"
    append FFMPEG_CMD_ARGS "${AUDIO_PROXY_BITRATE}"
    append FFMPEG_CMD_ARGS "${FULL_EDIT_PATH}/${OUTPUT_BASE}${REVISION_SUFFIX}_${audio_suffix}_proxy.ac3"
}

# Функция для построения всех аудио карт
add_all_audio_maps() {
    local audio_count="$1"
    shift
    local audio_suffixes=("$@")
    
    if [[ "$audio_count" -eq 0 ]]; then
        return 0
    fi
    
    for i in $(seq 0 $((audio_count - 1))); do
        add_audio_hd_map_for_track "$i" "${audio_suffixes[i]}"
        add_audio_proxy_map_for_track "$i" "${audio_suffixes[i]}"
    done
}

# Функция для построения всей команды ffmpeg
build_ffmpeg_command_array() {
    local audio_count="$1"
    shift
    local audio_suffixes=("$@")
    
    # Инициализация массива
    FFMPEG_CMD_ARGS=()
    
    # Начинаем с команды fflite
    append FFMPEG_CMD_ARGS "fflite"
    
    # Добавляем глобальные параметры
    add_global_parameters
    
    # Добавляем входные файлы
    add_inputs
    
    # Добавляем фильтры
    add_filter_complex "$audio_count" "${audio_suffixes[@]}"
    
    # Добавляем карты видео
    add_video_hd_map
    add_video_proxy_map
    
    # Добавляем карты аудио
    add_all_audio_maps "$audio_count" "${audio_suffixes[@]}"
}

# ============================================================================
# ОСНОВНОЙ КОД СКРИПТА
# ============================================================================

# Создание необходимых директорий
mkdir -p "${ARCHIVE_OUTPUT_PATH}"
mkdir -p "${FULL_EDIT_PATH}"

# Очистка экрана и перемещение входного файла
clear
mv "${INPUT_DIR}/${INPUT_FILE}" "${PROGRESS_DIR}/"

# Определение аудиодорожек с помощью внешней утилиты
declare -a AUDIO_SUFFIXES
# Предполагаем, что defineAudioStreams возвращает строку с разделителями (например, пробел или запятая)
temp_streams=$(defineAudioStreams "${PROGRESS_DIR}/${INPUT_FILE}")
split_string_to_array AUDIO_SUFFIXES "$temp_streams" " "
AUDIO_TRACK_COUNT=${#AUDIO_SUFFIXES[@]}

echo "Найдено аудиодорожек: ${AUDIO_TRACK_COUNT}"
if [[ ${AUDIO_TRACK_COUNT} -gt 0 ]]; then
    echo "Суффиксы аудиодорожек: ${AUDIO_SUFFIXES[*]}"
fi

# Построение и выполнение команды ffmpeg
echo "Построение команды ffmpeg..."
build_ffmpeg_command_array "$AUDIO_TRACK_COUNT" "${AUDIO_SUFFIXES[@]}"

echo "Выполнение команды:"
printf '%s\n' "${FFMPEG_CMD_ARGS[@]}"
echo ""

# Запуск команды
"${FFMPEG_CMD_ARGS[@]}" && \
  
# Создание файла готовности
touch "${READY_FILE_PATH}" && \
  
# Создание файла уведомления
echo "${EDIT_SUBPATH}/${OUTPUT_BASE}.ready" > "${FULL_NOTIFICATION_PATH}" && \
  
# Перемещение обработанного файла во временную директорию
mv "${PROGRESS_DIR}/${INPUT_FILE}" "${FULL_DONE_PATH}/" && \
  
# Отложенное перемещение файла в архив
at now + ${ARCHIVE_DELAY} <<< "mv ${FULL_DONE_PATH}/${INPUT_FILE} ${ARCHIVE_OUTPUT_PATH}" && \
  
# Очистка и архивация скрипта
clear
mv "$0" "${SCRIPT_ARCHIVE_DIR}/"
```