# Задание: Корзина товаров

В рамках задания вам необходимо реализовать интерфейсы Product и Basket.
А также написать тесты на корзину. 
В рамках задания тесты должны быть реализованы в рамках **тестовых функций** (сьюты не использовать).

> Цена на товар в рамках задания будет в условных единицах, которые принадлежат множеству натуральных чисел.

> Вес товара в рамках задания будет в условных единицах, которые принадлежат множеству натуральных чисел.

### Product
Product — это товар, который присутствует на маркетплейсе. 

У товара есть следующие атрибуты: 
- Название 
- Идентификатор 
- Цена 
- Вес

### Basket
Basket — это корзина. 

В корзину можно: 
- Добавлять товар в любом количестве 
- Удалять товар из корзины (целиком)
- Получать список товаров
- Получать итоговую цену на товары с учетом цены на доставку
- Получать цену доставки

Цена доставки зависит от итоговой цены всех товаров: 
1. Если итоговая цена меньше 500 условных единиц — цена доставки 250
2. Если цена больше либо равна 500 но меньше 1000 — цена 100
3. Если цена больше либо равна 1000 — 0

### Требования к системе: 
- Цена на товар не может быть меньше 1
- Максимальный вес товаров в корзине = 100
- Максимальное количество товаров в корзице = 30

Товары можете выбирать по своему усмотрению. Ограничений нет
