# Вычислитель отличий на Go

[![hexlet-check](https://github.com/knokkelmanns/go-from-scratch-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/knokkelmanns/go-from-scratch-project-244/actions)
[![ci](https://github.com/knokkelmanns/go-from-scratch-project-244/actions/workflows/ci.yml/badge.svg)](https://github.com/knokkelmanns/go-from-scratch-project-244/actions/workflows/ci.yml)

Консольная утилита для сравнения вложенных структур (JSON, YAML)

Учебный проект Хекслета: https://ru.hexlet.io/programs/go-from-scratch


## Стек

- Go

## Установка

```bash
git clone https://github.com/knokkelmanns/go-from-scratch-project-244.git
cd go-from-scratch-project-244
make build
```

## Разработка

```bash
make test           # тесты
make lint           # линтер (golangci-lint)
make test-coverage  # покрытие, падает если ниже COVERAGE_MIN (по умолчанию 80%)
```

## Использование

```bash
./bin/gendiff testdata/fixture/file1.json testdata/fixture/file2.json
```

```
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
```

### Сравнение плоских JSON-файлов

[![asciicast](https://asciinema.org/a/ld27wMIzbczC8etT.svg)](https://asciinema.org/a/ld27wMIzbczC8etT)

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
