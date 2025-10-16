@echo off
setlocal enabledelayedexpansion

set "output_file=files_content.txt"
set "excluded_folders=.git xxx"

echo Удаляем старый файл отчета...
if exist "%output_file%" del "%output_file%"

for /r %%f in (*.go *.html *.json) do (
    set "file_path=%%f"
    set "file_dir=%%~dpf"
    set "skip_file=0"
    
    rem Пропускаем выходной файл
    if "%%~nxf"=="%output_file%" set "skip_file=1"
    
    rem Проверяем исключенные папки
    for %%e in (%excluded_folders%) do (
        if not "!skip_file!"=="1" (
            echo !file_dir! | findstr /i "\\%%e\\" >nul && set "skip_file=1"
        )
    )
    
    if "!skip_file!"=="0" (
        echo Обрабатывается: %%f
        echo FILE: %%f >> "%output_file%"
        type "%%f" >> "%output_file%" 2>nul
        if errorlevel 1 (
            echo [ОШИБКА ЧТЕНИЯ ФАЙЛА] >> "%output_file%"
        )
        echo. >> "%output_file%"
        echo --- >> "%output_file%"
        echo. >> "%output_file%"
    ) else (
        echo Пропущен: %%f
    )
)

echo.
echo Обработка завершена!
echo Результат в: %output_file%
pause