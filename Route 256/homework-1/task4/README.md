# Дополнительное задание: Определение спам номеров 

Вам необходимо реализовать функцию проверки спам номеров и тесты на нее. 

На вход функции будут передаваться пути до двух файлов: 
- inputFile — входной файл, в котором будут лежать номера телефонов
- outputFile — выходной файл, в котором будет лежать результат обработки

Пример входного файла: 
```
+74962255222
+7 (612) 122-22-22
8 (812) 222-22-22
84962212222
849622222221
```

Пример выходного файла:
```
+74962255222 OK
+7 (612) 122-22-22 SPAM
8 (812) 222-22-22 SPAM
84962212222 OK
849622222221 ERROR
```

Паттерны телефонов: 
- 8хххххххххх
- +7хххххххххх
- 7хххххххххх
- 7 (ххх) ххх-хх-хх
- +7 (ххх) ххх-хх-хх
- 8 (ххх) ххх-хх-хх
- 7(ххх)ххх-хх-хх
- 8(ххх)ххх-хх-хх
- +7(ххх)ххх-хх-хх
 

Валидный номер телефона состоит только из 11 символов.

Номер телефона состоит из: 
**+A(BBB)CCC-CC-CC**

Где:
А — код страны
B — код города 
С — номер телефона 

Номер считается спамом, если: 
1. Код города начинается не с 8
2. Номер телефона состоит из одинаковых цифр
