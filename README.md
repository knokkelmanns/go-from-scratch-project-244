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
./bin/gendiff [--format stylish|plain|json] <file1> <file2>
```

Поддерживаемые форматы входных файлов: JSON (`.json`) и YAML (`.yml`, `.yaml`), их можно смешивать. Формат вывода по умолчанию — `stylish`.

[![asciicast](https://asciinema.org/a/N84afidljTWFz2D0.svg)](https://asciinema.org/a/N84afidljTWFz2D0)

### Формат stylish

```bash
./bin/gendiff testdata/fixture/nested1.json testdata/fixture/nested2.yml
```

```
{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow: 
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}
```

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

### Формат json

```bash
./bin/gendiff --format json testdata/fixture/nested1.json testdata/fixture/nested2.json
```

Выводит дерево диффа массивом узлов. У каждого узла есть `key` и `type` (`added`, `removed`, `unchanged`, `changed`, `nested`); в зависимости от типа добавляется `value`, пара `oldValue`/`newValue` или `children`.

```json
[
  {
    "key": "common",
    "type": "nested",
    "children": [
      { "key": "follow", "type": "added", "value": false },
      { "key": "setting3", "type": "changed", "oldValue": true, "newValue": null }
    ]
  },
  { "key": "group2", "type": "removed", "value": { "abc": 12345, "deep": { "id": 45 } } }
]
```

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
