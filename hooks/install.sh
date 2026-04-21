#!/bin/sh

echo "🔧 Устанавливаем хуки..."
git config core.hooksPath hooks/
echo "✅ Хуки подключены!"
echo ""
echo "Активные хуки:"
echo "  pre-commit  — защита от .env файлов"
echo "  post-commit — отчёт с тестами и линтером"
echo "  pre-push    — ЧЕРЕМША + запуск тестов"
