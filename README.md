# go-word-counter

Программа читает из stdin строки, содержащие URL. На каждый URL нужно отправить HTTP-запрос методом GET и посчитать кол-во вхождений строки "Go" в теле ответа. В конце работы приложение выводит на экран общее количество найденных строк "Go" во всех переданных URL, например:

```bash
$ echo -e 'https://golang.org\nhttps://golang.org' | go run 1.go
Count for https://golang.org: 9
Count for https://golang.org: 9
Total: 18
```

Каждый URL должен начать обрабатываться сразу после считывания и параллельно со считыванием следующего. URL должны обрабатываться параллельно, но не более k=5 одновременно. Обработчики URL не должны порождать лишних горутин, т.е. если k=5, а обрабатываемых URL-ов всего 2, не должно создаваться 5 горутин.

Нужно обойтись без глобальных переменных и использовать только стандартную библиотеку.

## Использование
Скопируйте к себе репозиторий с помощью следующей команды (предполагается что у вас установлен **git**):

```bash
$ git clone https://github.com/igorzakhar/go-word-counter
```
и перейдите в каталог с проектом:
```bash
$ cd go-word-counter
```
В репозитории с проектом находится файл ```urls``` со списком url-адресов. Для проверки того, что в ходе работы программы ничего не ломается, при передаче невалидного url-адреса, в список url-адресов добавлены строки вида ```abcdefg```, ```1234565```, ```https://```.

Пример запуска программы:
```bash
$ cat urls | go run main.go
Count for https://golang.org/doc/: 75
Count for https://golang.org/doc/cmd: 16
Count for https://golang.org/doc/install: 44
Count for https://golang.org/: 20
Count for https://golang.org/project/: 42
Count for https://golang.org/pkg/: 36
Count for https://golang.org/pkg/bufio/: 29
Count for https://golang.org/ref/spec: 41
Total: 303
```
