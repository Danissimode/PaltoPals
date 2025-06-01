#!/bin/bash

# Корень проекта
ROOT_DIR="."

# 1. Заменяем Paltopals на Paltopals (чувствительно к регистру)
find "$ROOT_DIR" -type f -not -path '*/.git/*' -exec sed -i 's/Paltopals/Paltopals/g' {} +

# 2. Заменяем Paltopals на Paltopals (без учёта регистра, но не трогаем Paltopals)
find "$ROOT_DIR" -type f -not -path '*/.git/*' -exec sed -i 's/Paltopals/Paltopals/gI' {} +

# 3. Заменяем Palto на Palto (без учёта регистра, но не трогаем Paltopals/Paltopals)
find "$ROOT_DIR" -type f -not -path '*/.git/*' -exec sed -i 's/Palto/Palto/gI' {} +

# 4. Заменяем github.com/Danissimode/Palto на github.com/Danissimode/Palto
find "$ROOT_DIR" -type f -not -path '*/.git/*' -exec sed -i 's@github.com/Danissimode/Palto@github.com/Danissimode/Palto@g' {} +

# 5. Заменяем Danissimode на Danissimode
find "$ROOT_DIR" -type f -not -path '*/.git/*' -exec sed -i 's/Danissimode/Danissimode/g' {} +

echo "Автозамена завершена!" 