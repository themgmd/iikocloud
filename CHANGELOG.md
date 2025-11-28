# Changelog

Все значительные изменения в этом проекте будут задокументированы в этом файле.

## [v0.0.8]

### Added

#### Справочники и базовая информация

- **StopLists** - Получение стоп-листов организаций (`/api/1/stop_lists`)
  - Метод `StopLists()` для получения информации о стоп-листах продуктов

- **Streets** - Получение списка улиц (`/api/1/streets`)
  - Метод `Streets()` для работы с адресными справочниками

- **Cities** - Получение списка городов (`/api/1/cities`)
  - Метод `Cities()` для работы с географическими справочниками

- **Regions** - Получение списка регионов (`/api/1/regions`)
  - Метод `Regions()` для работы с региональными справочниками

- **DeliveryZones** - Получение зон доставки (`/api/1/delivery/zones`)
  - Метод `DeliveryZones()` для работы с зонами доставки

- **MarketingSources** - Получение источников маркетинга (`/api/1/marketing_sources`)
  - Метод `MarketingSources()` для работы с маркетинговыми источниками

- **MarketingCampaigns** - Получение маркетинговых кампаний (`/api/1/marketing_campaigns`)
  - Метод `MarketingCampaigns()` для работы с маркетинговыми кампаниями

- **Operators** - Получение операторов (`/api/1/operators`)
  - Метод `Operators()` для работы с операторами

- **Employees** - Получение сотрудников (`/api/1/employees`)
  - Метод `Employees()` для работы с сотрудниками

- **DiscountTypes** - Получение типов скидок (`/api/1/discounts`)
  - Метод `DiscountTypes()` для работы с типами скидок

- **TipsTypes** - Получение типов чаевых (`/api/1/tips`)
  - Метод `TipsTypes()` для работы с типами чаевых

- **OrderTypes** - Получение типов заказов (`/api/1/order_types`)
  - Метод `OrderTypes()` для работы с типами заказов

- **Tables** - Получение списка столов (`/api/1/tables`)
  - Метод `Tables()` для работы с таблицами ресторана

#### Работа с заказами

- **OrdersByID** - Получение заказов по ID (`/api/1/orders/by_id`)
  - Метод `OrdersByID()` для получения информации о заказах

- **OrderStatus** - Получение статуса заказов (`/api/1/orders/status`)
  - Метод `OrderStatus()` для отслеживания статуса заказов

- **CloseOrder** - Закрытие заказа (`/api/1/orders/close`)
  - Метод `CloseOrder()` для закрытия заказов

- **CancelOrder** - Отмена заказа (`/api/1/orders/cancel`)
  - Метод `CancelOrder()` для отмены заказов

#### Работа с доставками

- **DeliveriesByID** - Получение доставок по ID (`/api/1/deliveries/by_id`)
  - Метод `DeliveriesByID()` для получения информации о доставках

- **DeliveriesStatus** - Получение статуса доставок (`/api/1/deliveries/status`)
  - Метод `DeliveriesStatus()` для отслеживания статуса доставок

#### Работа с клиентами (Loyalty)

- **CustomerByID** - Получение клиента по ID (`/api/1/loyalty/iiko/customer/get_by_id`)
  - Метод `CustomerByID()` для получения информации о клиенте

- **CustomerSearch** - Поиск клиентов (`/api/1/loyalty/iiko/customer/search`)
  - Метод `CustomerSearch()` для поиска клиентов по различным критериям

- **CustomerCreateOrUpdate** - Создание или обновление клиента (`/api/1/loyalty/iiko/customer/create_or_update`)
  - Метод `CustomerCreateOrUpdate()` для создания или обновления информации о клиенте

#### Работа с кошельком клиента

- **CustomerWalletBalance** - Получение баланса кошелька (`/api/1/loyalty/iiko/customer/wallet/balance`)
  - Метод `CustomerWalletBalance()` для получения баланса кошелька клиента

- **CustomerWalletHistory** - История операций кошелька (`/api/1/loyalty/iiko/customer/wallet/history`)
  - Метод `CustomerWalletHistory()` для получения истории операций

- **CustomerWalletHold** - Блокировка средств (`/api/1/loyalty/iiko/customer/wallet/hold`)
  - Метод `CustomerWalletHold()` для блокировки средств на кошельке клиента

- **CustomerWalletRefund** - Возврат средств (`/api/1/loyalty/iiko/customer/wallet/refund`)
  - Метод `CustomerWalletRefund()` для возврата средств на кошелек клиента

- **CustomerWalletChargeOff** - Списание средств (`/api/1/loyalty/iiko/customer/wallet/chargeoff`)
  - Метод `CustomerWalletChargeOff()` для списания средств с кошелька клиента

#### Работа с купонами

- **Coupons** - Получение информации о купонах (`/api/1/loyalty/iiko/coupons`)
  - Метод `Coupons()` для работы с купонами

- **CouponsByID** - Получение купонов по ID (`/api/1/loyalty/iiko/coupons/by_id`)
  - Метод `CouponsByID()` для получения информации о купонах по их ID

### Changed

- **CreateOrderRequest** - Улучшена структура запроса для создания заказа
  - Добавлен тип `CreateOrderCustomer`, расширяющий `OrderCustomer` дополнительными полями
  - Все поля сделаны опциональными с использованием `omitempty`
  - Улучшена совместимость с API документацией

- **V1CreateOrder endpoint** - Исправлена ошибка в пути endpoint
  - Исправлен двойной слеш в пути: `/api/1//order/create` → `/api/1/order/create`

### Summary

В этой версии добавлено **28 новых методов API**, которые покрывают большую часть функциональности iiko Cloud API:

- 13 методов для работы со справочниками (стоп-листы, улицы, города, регионы, зоны доставки, источники маркетинга, кампании, операторы, сотрудники, типы скидок/чаевых/заказов, столы)
- 4 метода для работы с заказами (получение по ID, статус, закрытие, отмена)
- 2 метода для работы с доставками (получение по ID, статус)
- 7 методов для работы с клиентами и программами лояльности (поиск, получение, создание/обновление, баланс кошелька, история, блокировка, возврат средств)
- 2 метода для работы с купонами

Все методы следуют единой структуре проекта, используют общий механизм обработки запросов и ответов, поддерживают кастомные таймауты и авторизацию через Bearer токен.

