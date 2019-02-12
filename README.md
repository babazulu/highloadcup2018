# github.com/covrom/highloadcup2018

Мое решение для Highload Cup 2018.

* Техническое задание: https://highloadcup.ru/media/condition/accounts_rules.html

В финале https://highloadcup.ru/ru/rating/ мое решение занимает 28 место с ~410 сек. штрафа и потреблением памяти 1,6-1,7 Гб.

Ноу-хау решения стал разработанный универсальный движок columnstore (`db/column.go`) для любого содержимого колонок 
и быстрые intersect/merge-итераторы без аллокации памяти (`db/coliter.go`).

Все алгоритмы хэш или основанные на деревьях - не пригодились по причине того, что они сильно тратят память на указатели (именно на сами указатели, а не на данные), которые дополнительно еще обрабатывает и GC. Из-за этого было принято решение реализовывать алгоритмы на массивах, без указателей.

Для сети использовался `fasthttp` и `fasthttprouter` без переделок.

В папке `docs` содержатся материалы по inmemory column - базам данных.

Содержимое некоторых пакетов:
* `avltree` - содержатся хорошие примеры реализации avl-деревьев, но они не пригодились при решении задачи.
* `flysort` - содержится бенчмарк алгоритмов добавления и сортировки "на лету" - линейный поиск и вставка, бинарный поиски и вставка, мин-хип.
Победитель (мин-хип) был использован в пакете suggest.
* `gentest` - лежит утилита обработки исходных данных, для подсчета некоторых статистик.
* `intsearch` - содержится бенчмарк различных вариантов реализации бинарного поиска в сортированном массиве. Безусловный победитель - алгоритм на ассемблере.
* `jumphash` - содержится хэш-функция, которая не пригодилась.
* `treebidimap` - содержится алгоритм двунаправленной мапы, который тоже не пригодился.
