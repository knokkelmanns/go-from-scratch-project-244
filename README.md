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

Поддерживаемые форматы входных файлов: JSON (`.json`) и YAML (`.yml`, `.yaml`). Форматы можно смешивать.

```bash
./bin/gendiff testdata/fixture/file1.json testdata/fixture/file2.json
./bin/gendiff testdata/fixture/file1.yml testdata/fixture/file2.yaml
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

### Сравнение плоских YAML-файлов

[![asciicast](https://asciinema.org/a/ASCIINEMA_ID.svg)](https://asciinema.org/a/ASCIINEMA_ID)

### Рекурсивное сравнение (формат stylish)

```bash
./bin/gendiff testdata/fixture/nested1.json testdata/fixture/nested2.json
```

[![asciicast](https://asciinema.org/a/ASCIINEMA_NESTED_ID.svg)](https://asciinema.org/a/ASCIINEMA_NESTED_ID)

### Формат plain

```bash
./bin/gendiff --format plain testdata/fixture/nested1.json testdata/fixture/nested2.json
```

```
Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]
```

[![asciicast](https://asciinema.org/a/ASCIINEMA_PLAIN_ID.svg)](https://asciinema.org/a/ASCIINEMA_PLAIN_ID)

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
