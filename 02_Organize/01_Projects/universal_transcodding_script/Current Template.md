---
updated_at: 2025-12-19T19:22:45.648+10:00
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

---

```bash
#!/bin/bash

# ============================================================================
# ГЛОБАЛЬНЫЕ ПЕРЕМЕННЫЕ И МАССИВЫ
# ============================================================================

#Пути

ROOT_EXTERNAL="/mnt/pemaltynov/ROOT"

DESTINATION_ROOT="${ROOT_EXTERNAL}/EDIT"
ARCHIVE_ROOT="${ROOT_EXTERNAL}/IN"

ROOT_LOCAL="/home/pemaltynov/IN"
DIR_SOURCE="${ROOT_LOCAL}"
DIR_IN_PROGRESS="${ROOT_LOCAL}/_IN_PROGRESS"
DIR_DONE="${ROOT_LOCAL}/_DONE"
DIR_NOTIFICATIONS="${ROOT_LOCAL}/notifications"



# Основные параметры кодирования
MEDIA_AGENT="amedia"
VIDEO_FRAMERATE=24
VIDEO_ENABLED=true
VIDEO_PROXY_ENABLED=true
AUDIO_ENABLED=true
AUDIO_PROXY_ENABLED=false

# Массивы фильтров filter_complex
declare -a VIDEO_FILTER_COMPLEX_MAIN=(
    "scale=1920:-2"
    "setsar=1/1"
    "unsharp=3:3:0.3:3:3:0"
    "pad=1920:1080:-1:-1"
)
declare -a VIDEO_FILTER_COMPLEX_PROXY=(
    "scale=iw/2:ih"
    "setsar=(1/1)*2"
)

declare -a AUDIO_FILTER_COMPLEX_MAIN=(
    "aresample=48000"
    "atempo=25/(__)"
)
declare -a AUDIO_FILTER_COMPLEX_PROXY=(
    "aresample=44100"
    "atempo=25/(__)"
)


# Массивы фильтров вывода потоков
declare -a VIDEO_OUTSTREAM_FILTERS_MAIN=(
	"-c:v" "libx264" "-preset" "medium" "-crf" "16" "-pix_fmt" "yuv420p" "-g" "0"
)
declare -a VIDEO_OUTSTREAM_FILTERS_PROXY=(
	"-c:v" "libx264" "-x264opts" "interlaced=1" "-preset" "superfast"
	"-pix_fmt" "yuv420p" "-b:v" "2000k" "-maxrate" "2000k"
)

declare -a AUDIO_OUTSTREAM_FILTERS_MAIN=(
    "-c:a" "alac" "-compression_level" "0"
)
declare -a AUDIO_OUTSTREAM_FILTERS_PROXY=(
    "-c:a" "ac3" "-b:a" "128k"
)

# Массивы кодов потоков из входных файлов
declare -a STREAM_CODES_VIDEO=(
    "[0:v:0]"
)
declare -a STREAM_CODES_AUDIO=(
    "[0:a:1]"
    "[0:a:3]"
)

declare -a OUTPUT_CODES

#Словарь "Код выходного потока -> имя файла"
declare -A OUTPUT_CODE_TO_FILENAME

# Команда ffmpeg и входные файлы
declare -a FFMPEG_COMMAND=(
    "ffmpeg"
)
declare -a FFMPEG_INPUT_FILES=(
    "file1.mp4"
    "file2.mp4"
)

# ============================================================================
# ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ
# ============================================================================

# Функция добавления элементов в массив по ссылке
# Принимает имя массива и произвольное количество значений для добавления
# Проверяет наличие минимум 2 аргументов и добавляет только непустые строки
append() {
    if [[ $# -lt 2 ]]; then
        echo "Ошибка: функция ожидает минимум 2 аргумента - имя массива и строки" >&2
        return 1
    fi

    local -n arr_ref="$1"
    shift
    
    local value
    for value in "$@"; do
        if [[ -n "$value" ]]; then
            arr_ref+=("$value")
        fi
    done
}

# Функция объединения элементов массива в строку с заданным разделителем
# Принимает имя массива и строку-разделитель, возвращает объединенную строку
# Обрабатывает случаи пустого массива и массива с одним элементом
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

# Функция извлечения меток выходных потоков из строки filter_complex
# Принимает строку filter_complex и имя массива для результатов
# Находит все вхождения вида [out_*] и добавляет их в результирующий массив
extract_out_labels() {
    local filter_complex="$1"
    local -n array_ref="$2"
    
    array_ref=()
    
    # Используем while для чтения, обрабатывая пустые строки
    while IFS= read -r match || [[ -n "$match" ]]; do
        # Пропускаем пустые строки
        [[ -z "$match" ]] && continue
        array_ref+=("$match")
    done < <(printf '%s' "$filter_complex" | grep -oE '\[out_[^]]+\]')
}

# ============================================================================
# ФУНКЦИИ ПОСТРОЕНИЯ VIDEO ЧАСТИ FILTER_COMPLEX
# ============================================================================

# Функция создания цепочки фильтров для одного видеопотока
# Строит последовательность фильтров с опциональным разделением на основной и прокси потоки
# Принимает код входного потока и индекс видеопотока, возвращает строку фильтров
build_filter_complex_video_filter() {
    local stream_code="$1"
    local video_index="$2"
    
    local video_filter=""
    video_filter+="${stream_code}"
    
    # Добавление основных фильтров видео
    if [[ ${#VIDEO_FILTER_COMPLEX_MAIN[@]} -gt 0 ]]; then
        local filters_joined
        filters_joined=$(join_array VIDEO_FILTER_COMPLEX_MAIN ",")
        video_filter+="${filters_joined}"
    fi

    local main_out_code="[out_vm_${video_index}]"
	append OUTPUT_CODES "$main_out_code"
    
    # Проверка необходимости создания прокси видео
    if [[ "${VIDEO_PROXY_ENABLED}" == "true" ]] && [[ ${#VIDEO_FILTER_COMPLEX_PROXY[@]} -gt 0 ]]; then
        local proxy_out_code="[out_vp_${video_index}]"
        # Разделение потока на основной и прокси
        video_filter+=",split=2"
        video_filter+="${main_out_code}[in_proxy_${video_index}]"
        video_filter+=";"
        video_filter+="[in_proxy_${video_index}]"
        
        # Добавление фильтров прокси видео
        if [[ ${#VIDEO_FILTER_COMPLEX_PROXY[@]} -gt 0 ]]; then
            local filters_joined
            filters_joined=$(join_array VIDEO_FILTER_COMPLEX_PROXY ",")
            video_filter+="${filters_joined}"
        fi
        
        video_filter+="${proxy_out_code}"
		append OUTPUT_CODES "$proxy_out_code"
    else
        # Только основной поток без прокси
        video_filter+="${main_out_code}"
		append OUTPUT_CODES "$main_out_code"
    fi
    
    echo "$video_filter"
}

# Функция построения видео части filter_complex для всех видеопотоков
# Обрабатывает все видеопотоки, объединяя их фильтры через точку с запятой
# Работает с длиной массива STREAM_CODES_VIDEO
build_filter_complex_videos() {
    local -a video_filters_array=()
    local video_filter
    local result=""
    
    # Обработка всех видеопотоков из массива
    for ((i=0; i<${#STREAM_CODES_VIDEO[@]}; i++)); do
        local video_index=$i
        local stream_code="${STREAM_CODES_VIDEO[i]}"
        video_filter=$(build_filter_complex_video_filter "$stream_code" "$video_index")
        
        # Добавление фильтра в массив, если он не пустой
        if [[ -n "$video_filter" ]]; then
            append video_filters_array "$video_filter"
        fi
    done
    
    # Объединение всех видеофильтров через точку с запятой
    if [[ ${#video_filters_array[@]} -gt 0 ]]; then
        result=$(join_array video_filters_array ";")
    fi
    
    echo "$result"
}

# ============================================================================
# ФУНКЦИИ ПОСТРОЕНИЯ AUDIO ЧАСТИ FILTER_COMPLEX
# ============================================================================

# Функция создания цепочки фильтров для одного аудиопотока
# Строит последовательность фильтров с опциональным разделением на основной и прокси потоки
# Принимает код входного потока и индекс аудиопотока, возвращает строку фильтров
build_filter_complex_audio_filter() {
    local stream_code="$1"
    local audio_index="$2"
    
    local audio_filter=""
    audio_filter+="${stream_code}"
    
    # Добавление основных фильтров аудио
    if [[ ${#AUDIO_FILTER_COMPLEX_MAIN[@]} -gt 0 ]]; then
        local filters_joined
        filters_joined=$(join_array AUDIO_FILTER_COMPLEX_MAIN ",")
        audio_filter+="${filters_joined}"
    fi

    local main_out_code="[out_am_${audio_index}]"
	append OUTPUT_CODES "$main_out_code"
    
    # Проверка необходимости создания прокси аудио
    if [[ "${AUDIO_PROXY_ENABLED}" == "true" ]] && [[ ${#AUDIO_FILTER_COMPLEX_PROXY[@]} -gt 0 ]]; then
        local proxy_out_code="[out_ap_${audio_index}]"
        # Разделение аудиопотока на основной и прокси
        audio_filter+=",asplit=2"
        audio_filter+="${main_out_code}[pr_raw_a_${audio_index}]"
        audio_filter+=";"
        audio_filter+="[pr_raw_a_${audio_index}]"
        
        # Добавление фильтров прокси аудио
        if [[ ${#AUDIO_FILTER_COMPLEX_PROXY[@]} -gt 0 ]]; then
            local filters_joined
            filters_joined=$(join_array AUDIO_FILTER_COMPLEX_PROXY ",")
            audio_filter+="${filters_joined}"
        fi
        
        audio_filter+="${proxy_out_code}"
		append OUTPUT_CODES "$proxy_out_code"
    else
        # Только основной поток без прокси
        audio_filter+="${main_out_code}"
		append OUTPUT_CODES "$main_out_code"
    fi
    
    echo "$audio_filter"
}

# Функция построения аудио части filter_complex для всех аудиопотоков
# Обрабатывает все аудиопотоки, объединяя их фильтры через точку с запятой
# Работает с длиной массива STREAM_CODES_AUDIO
build_filter_complex_audios() {
    local -a audio_filters_array=()
    local audio_filter
    local result=""
    
    # Обработка всех аудиопотоков из массива
    for ((i=0; i<${#STREAM_CODES_AUDIO[@]}; i++)); do
        local audio_index=$i
        local stream_code="${STREAM_CODES_AUDIO[i]}"
        audio_filter=$(build_filter_complex_audio_filter "$stream_code" "$audio_index")
        
        # Добавление фильтра в массив, если он не пустой
        if [[ -n "$audio_filter" ]]; then
            append audio_filters_array "$audio_filter"
        fi
    done
    
    # Объединение всех аудиофильтров через точку с запятой
    if [[ ${#audio_filters_array[@]} -gt 0 ]]; then
        result=$(join_array audio_filters_array ";")
    fi
    
    echo "$result"
}

# ============================================================================
# ФУНКЦИЯ ПОСТРОЕНИЯ ПОЛНОГО FILTER_COMPLEX
# ============================================================================

# Основная функция сборки полной строки filter_complex для ffmpeg
# Объединяет видео и аудио части, обеспечивая правильный синтаксис разделения
# Работает с длинами массивов STREAM_CODES_VIDEO и STREAM_CODES_AUDIO
build_filter_complex() {
    local video_part=""
    local audio_part=""
    local result=""
    
    # Построение видео части
    if [[ ${#STREAM_CODES_VIDEO[@]} -gt 0 ]]; then
        video_part=$(build_filter_complex_videos)
        if [[ $? -ne 0 ]]; then
            echo "Ошибка при построении видео части filter_complex" >&2
            return 1
        fi
    fi
    
    # Построение аудио части
    if [[ ${#STREAM_CODES_AUDIO[@]} -gt 0 ]]; then
        audio_part=$(build_filter_complex_audios)
        if [[ $? -ne 0 ]]; then
            echo "Ошибка при построении аудио части filter_complex" >&2
            return 1
        fi
    fi
    
    # Объединение видео и аудио частей
    if [[ -n "$video_part" ]] && [[ -n "$audio_part" ]]; then
        result="${video_part};${audio_part}"
    elif [[ -n "$video_part" ]]; then
        result="$video_part"
    elif [[ -n "$audio_part" ]]; then
        result="$audio_part"
    fi
    
    echo "$result"
}


# Функция для получения фильтров по коду
get_stream_filters() {
    local code="$1"
    local -n target_array_ref="$2"  # -n создает ссылку на массив
    
    # Очищаем целевой массив
    target_array_ref=()
    
    # Определяем источник на основе кода
    case "$code" in
        *vm*) 
            target_array_ref=("${VIDEO_OUTSTREAM_FILTERS_MAIN[@]}")
            ;;
        *vp*) 
            target_array_ref=("${VIDEO_OUTSTREAM_FILTERS_PROXY[@]}")
            ;;
        *am*) 
            target_array_ref=("${AUDIO_OUTSTREAM_FILTERS_MAIN[@]}")
            ;;
        *ap*) 
            target_array_ref=("${AUDIO_OUTSTREAM_FILTERS_PROXY[@]}")
            ;;
        *) 
            # Оставляем массив пустым
            echo "Неизвестный код: $code" >&2
            return 1
            ;;
    esac
    
    return 0
}


# ============================================================================
# ФУНКЦИЯ ПОСТРОЕНИЯ КОМАНДЫ FFMPEG
# ============================================================================

# Функция формирования полной команды ffmpeg для выполнения
# Собирает массив аргументов из глобальных переменных и переданного filter_complex
# Выводит готовую команду построчно для удобства чтения и использования
build_ffmpeg_command() {
    local filter_complex_value="$1"
    
    local -a cmd=()
    
    # Базовые параметры команды
    cmd+=("${FFMPEG_COMMAND[@]}")
    cmd+=("-r" "${VIDEO_FRAMERATE}")
    
    # Входные файлы
    for input_file in "${FFMPEG_INPUT_FILES[@]}"; do
        cmd+=("-i" "${input_file}")
    done
    
    # Добавление filter_complex
    cmd+=("-filter_complex" "${filter_complex_value}")
	declare -a output_labels
    extract_out_labels "$filter_complex_value" output_labels

    for outstream_code in "${output_labels[@]}"; do
        cmd+=("${outstream_code}")
		
		declare -a outstream_filters
		if get_stream_filters "${outstream_code}" "outstream_filters"; then
			for i in "${!outstream_filters[@]}"; do
				cmd+=("${outstream_filters[$i]}")
			done
		cmd+=("-map_metadata" "-1" "-map_chapters" "-1")
		
		cmd+=("${outstream_code}_OUT_FILE_NAME")
		
	fi
		
		
		
		
		
    done
    
    printf '%s ' "${cmd[@]}"
	
}

# ============================================================================
# ОСНОВНАЯ ФУНКЦИЯ
# ============================================================================

# Главная функция скрипта, координирующая процесс построения и вывода результатов
# Последовательно выполняет построение filter_complex, извлечение меток и формирование команды
# Выводит диагностическую информацию и итоговую команду для выполнения
main() {
    ffmpeg_command=$(build_filter_complex)
    
    if [[ $? -ne 0 ]]; then
        echo "Ошибка при построении filter_complex" >&2
        return 1
    fi

    
    
    echo ""
    echo "Результат filter_complex:"
    echo "$ffmpeg_command"
    
    echo ""
    echo "Команда ffmpeg:"
    build_ffmpeg_command "$ffmpeg_command"
	
	
    
    return 0
}

# Запуск основной функции при прямом выполнении скрипта
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi

```