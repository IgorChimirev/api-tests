# 📋 Отчёт по проекту

> Сгенерировано автоматически при коммите

---

## 🔖 Информация о коммите

| Параметр | Значение |
|----------|----------|
| Хэш | `a47a1b3` |
| Сообщение | qwqwq |
| Автор | IgorChimirev |
| Дата | 2026-04-27 12:45:22 |
| Всего коммитов | 25 |

---

## 🧪 Результаты тестов

**Статус: ✅ Все тесты прошли**

| Метрика | Значение |
|---------|----------|
| Прошло | ✅ 40 |
| Упало | ❌ 0 |

<details>
<summary>Полный вывод тестов</summary>

```
?   	api-tests-template/internal/client/http/advertisement	[no test files]
?   	api-tests-template/internal/client/http/auth	[no test files]
?   	api-tests-template/internal/client/http/myAdvertisements	[no test files]
?   	api-tests-template/internal/client/http/photoCheck	[no test files]
?   	api-tests-template/internal/constants/path	[no test files]
?   	api-tests-template/internal/helpers/advertisement	[no test files]
?   	api-tests-template/internal/helpers/api-runner	[no test files]
?   	api-tests-template/internal/managers/advertisement	[no test files]
?   	api-tests-template/internal/managers/advertisement/models	[no test files]
?   	api-tests-template/internal/managers/auth	[no test files]
?   	api-tests-template/internal/managers/auth/models	[no test files]
?   	api-tests-template/internal/managers/myAdvertisements	[no test files]
?   	api-tests-template/internal/utils	[no test files]
?   	api-tests-template/tests	[no test files]
=== RUN   TestSuiteRun
2026/04/27 09:06:55 Init environment variables
2026/04/27 09:06:55 [Precondition]: Авторизация пользователя с кредами из переменных окружения
2026/04/27 09:06:56 [Precondition]: Чтение тестовых фотографий из папки testdata
=== RUN   TestSuiteRun/TestCreateAdvertisementAllFields
=== RUN   TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_1:_создаём_объявление_со_всеми_полями
    advertisement_test.go:80: Response body: {"id":"21c3bdf0-a401-4002-b62f-47ca1cbfc795","title":"Test eI9cXlAD","description":"Description BvTptlKbKCx1tQ1U","price":1000,"quantity":5,"user_id":"8958d32f-1de5-4bcc-9f2f-2c4633d64b53","created_at":"2026-04-27T06:06:56.787056Z","updated_at":"2026-04-27T06:06:56.787056Z","photos":[{"id":"01c367f8-4c80-4163-bee2-bc0c08795709","url":"https://storage.yandexcloud.net/testboard-ds-aaa/photos/1777270016_e848052f-80ec-4948-af2b-42aa3e1f852e.jpg","sort_order":0,"created_at":"2026-04-27T06:06:56.848404Z"},{"id":"56228e33-6347-43eb-8237-1ed0045be735","url":"https://storage.yandexcloud.net/testboard-ds-aaa/photos/1777270016_c1964ce1-5f4c-4ed2-aa7b-8f9f6f22bab6.jpg","sort_order":1,"created_at":"2026-04-27T06:06:56.89721Z"},{"id":"3e326e1b-c1d5-4441-a29e-6cb649e613bc","url":"https://storage.yandexcloud.net/testboard-ds-aaa/photos/1777270016_d2ba8a62-be08-40f0-82f1-42f1cd4a7dbc.jpg","sort_order":2,"created_at":"2026-04-27T06:06:56.930079Z"}]}
=== RUN   TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_2:_все_поля_ответа_соответствуют_переданным_значениям
=== RUN   TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_3:_GET_/advertisement?id={id}_возвращает_объект,_совпадающий_с_созданным
=== RUN   TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_4:_GET_/advertisements/{id}/photos_возвращает_все_фото_и_они_доступны_на_сервере
=== RUN   TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_5:_GET_/advertisements?search=..._—_объявление_находится_в_поиске_по_полному_названию
=== RUN   TestSuiteRun/TestCreateAdvertisementWithInvalidToken
=== RUN   TestSuiteRun/TestCreateAdvertisementWithInvalidToken/Создаём_объявление_с_невалидным_токеном
=== RUN   TestSuiteRun/TestCreateAdvertisementWithInvalidToken/Проверяем_сообщение_об_ошибке
=== RUN   TestSuiteRun/TestCreateAdvertisementWithNegativePrice
=== RUN   TestSuiteRun/TestCreateAdvertisementWithNegativePrice/Создаём_объявление_с_отрицательной_ценой
=== RUN   TestSuiteRun/TestCreateAdvertisementWithNegativePrice/Проверяем_наличие_ошибки_в_ответе
=== RUN   TestSuiteRun/TestCreateAdvertisementWithNegativeQuantity
=== RUN   TestSuiteRun/TestCreateAdvertisementWithNegativeQuantity/Создаём_объявление_с_отрицательным_quantity
=== RUN   TestSuiteRun/TestCreateAdvertisementWithNegativeQuantity/Проверяем_наличие_ошибки_в_ответе
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutDescription
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutDescription/Создаём_объявление_без_description
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutDescription/Проверяем_наличие_ошибки_в_ответе
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutPhotos
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutPhotos/Создаём_объявление_без_фото
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutPhotos/Проверяем_наличие_ошибки_в_ответе
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutTitle
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutTitle/Создаём_объявление_без_title
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutTitle/Проверяем_наличие_ошибки_в_ответе
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutToken
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutToken/Создаём_объявление_без_токена_авторизации
=== RUN   TestSuiteRun/TestCreateAdvertisementWithoutToken/Проверяем_сообщение_об_ошибке
=== RUN   TestSuiteRun/TestGetAdvertisementNotFound
=== RUN   TestSuiteRun/TestGetAdvertisementNotFound/Запрашиваем_объявление_по_несуществующему_id
=== RUN   TestSuiteRun/TestGetAdvertisementNotFound/Проверяем_сообщение_об_ошибке
=== RUN   TestSuiteRun/Удаляем_созданное_объявление_после_тестов
2026/04/27 09:06:59 Tear down suite
--- PASS: TestSuiteRun (4.50s)
    --- PASS: TestSuiteRun/TestCreateAdvertisementAllFields (1.48s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_1:_создаём_объявление_со_всеми_полями (0.51s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_2:_все_поля_ответа_соответствуют_переданным_значениям (0.00s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_3:_GET_/advertisement?id={id}_возвращает_объект,_совпадающий_с_созданным (0.08s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_4:_GET_/advertisements/{id}/photos_возвращает_все_фото_и_они_доступны_на_сервере (0.79s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementAllFields/Шаг_5:_GET_/advertisements?search=..._—_объявление_находится_в_поиске_по_полному_названию (0.08s)
    --- PASS: TestSuiteRun/TestCreateAdvertisementWithInvalidToken (0.25s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithInvalidToken/Создаём_объявление_с_невалидным_токеном (0.25s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithInvalidToken/Проверяем_сообщение_об_ошибке (0.00s)
    --- PASS: TestSuiteRun/TestCreateAdvertisementWithNegativePrice (0.38s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithNegativePrice/Создаём_объявление_с_отрицательной_ценой (0.38s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithNegativePrice/Проверяем_наличие_ошибки_в_ответе (0.00s)
    --- PASS: TestSuiteRun/TestCreateAdvertisementWithNegativeQuantity (0.19s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithNegativeQuantity/Создаём_объявление_с_отрицательным_quantity (0.19s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithNegativeQuantity/Проверяем_наличие_ошибки_в_ответе (0.00s)
    --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutDescription (0.42s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutDescription/Создаём_объявление_без_description (0.42s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutDescription/Проверяем_наличие_ошибки_в_ответе (0.00s)
    --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutPhotos (0.15s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutPhotos/Создаём_объявление_без_фото (0.15s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutPhotos/Проверяем_наличие_ошибки_в_ответе (0.00s)
    --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutTitle (0.49s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutTitle/Создаём_объявление_без_title (0.49s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutTitle/Проверяем_наличие_ошибки_в_ответе (0.00s)
    --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutToken (0.27s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutToken/Создаём_объявление_без_токена_авторизации (0.27s)
        --- PASS: TestSuiteRun/TestCreateAdvertisementWithoutToken/Проверяем_сообщение_об_ошибке (0.00s)
    --- PASS: TestSuiteRun/TestGetAdvertisementNotFound (0.17s)
        --- PASS: TestSuiteRun/TestGetAdvertisementNotFound/Запрашиваем_объявление_по_несуществующему_id (0.17s)
        --- PASS: TestSuiteRun/TestGetAdvertisementNotFound/Проверяем_сообщение_об_ошибке (0.00s)
    --- PASS: TestSuiteRun/Удаляем_созданное_объявление_после_тестов (0.08s)
PASS
ok  	api-tests-template/tests/scenarios/advertisement	(cached)
=== RUN   TestSuiteRun
2026/04/21 13:15:49 Init environment variables
2026/04/21 13:15:49 [Precondition]: Авторизация пользователя с кредами из переменных окружения и получение его параметров
=== RUN   TestSuiteRun/TestGetMyAdvertisementsIncorrectToken
=== RUN   TestSuiteRun/TestGetMyAdvertisementsIncorrectToken/Получаем_список_собственных_объявлений_с_неправильным_авторизационным_токеном
=== RUN   TestSuiteRun/TestGetMyAdvertisementsIncorrectToken/Проверяем,_что_показывается_правильная_ошибка
=== RUN   TestSuiteRun/TestGetMyAdvertisementsPositive
=== RUN   TestSuiteRun/TestGetMyAdvertisementsPositive/Получаем_список_собственных_объявлений
=== RUN   TestSuiteRun/TestGetMyAdvertisementsPositive/Проверяем,_что_у_нас_есть_несколько_добавленных_ранее_объявлений
=== RUN   TestSuiteRun/TestGetMyAdvertisementsPositive/Проверяем,_что_наши_объявления_имеют_принадлежность_к_нашему_пользователю_и_объявления_имеют_НЕ_пустые_параметры
2026/04/21 13:15:55 Tear down suite
--- PASS: TestSuiteRun (5.94s)
    --- PASS: TestSuiteRun/TestGetMyAdvertisementsIncorrectToken (0.17s)
        --- PASS: TestSuiteRun/TestGetMyAdvertisementsIncorrectToken/Получаем_список_собственных_объявлений_с_неправильным_авторизационным_токеном (0.16s)
        --- PASS: TestSuiteRun/TestGetMyAdvertisementsIncorrectToken/Проверяем,_что_показывается_правильная_ошибка (0.00s)
    --- PASS: TestSuiteRun/TestGetMyAdvertisementsPositive (0.12s)
        --- PASS: TestSuiteRun/TestGetMyAdvertisementsPositive/Получаем_список_собственных_объявлений (0.10s)
        --- PASS: TestSuiteRun/TestGetMyAdvertisementsPositive/Проверяем,_что_у_нас_есть_несколько_добавленных_ранее_объявлений (0.01s)
        --- PASS: TestSuiteRun/TestGetMyAdvertisementsPositive/Проверяем,_что_наши_объявления_имеют_принадлежность_к_нашему_пользователю_и_объявления_имеют_НЕ_пустые_параметры (0.00s)
PASS
ok  	api-tests-template/tests/scenarios/myAdvertisement	(cached)
```

</details>

---

## 🔍 Результаты линтера

**Статус: ✅ Ошибок не найдено**

<details>
<summary>Полный вывод линтера</summary>

```
Ошибок не найдено
```

</details>

---

## 📁 Статистика кода

| Метрика | Значение |
|---------|----------|
| Go файлов | 19 |
| Тест файлов | 2 |
| Строк кода | 945 |

---

## 🗂 Тест-сценарии

### Advertisement
- ✅ TestCreateAdvertisementAllFields {
- ✅ TestCreateAdvertisementWithoutToken {
- ✅ TestCreateAdvertisementWithInvalidToken {
- ✅ TestCreateAdvertisementWithoutTitle {
- ✅ TestCreateAdvertisementWithoutDescription {
- ✅ TestCreateAdvertisementWithNegativePrice {
- ✅ TestCreateAdvertisementWithNegativeQuantity {
- ✅ TestCreateAdvertisementWithoutPhotos {
- ✅ TestGetAdvertisementNotFound {

### My Advertisements
- ✅ TestGetMyAdvertisementsPositive {
- ✅ TestGetMyAdvertisementsIncorrectToken {

---
_Отчёт создан: 2026-04-27 12:45:25_
