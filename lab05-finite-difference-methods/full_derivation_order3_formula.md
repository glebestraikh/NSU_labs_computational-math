Ы# Полный вывод стартовой формулы 3-го порядка методом неопределённых коэффициентов

## Постановка задачи

Для уравнения $y' = -y$ с начальным условием $y_0 = -1$ требуется найти формулу для вычисления $y_1$ с **3-м порядком точности** (локальная ошибка O(h⁴)).

## Общая схема

Будем искать одношаговую разностную схему в виде:

$$
y_1 + \alpha_0 y_0 = h(\beta_1 y'_1 + \beta_0 y'_0) + h^2(\gamma_1 y''_1 + \gamma_0 y''_0) \quad (*)
$$

где коэффициенты $\alpha_0, \beta_0, \beta_1, \gamma_0, \gamma_1$ нужно определить.

## Подстановка производных

Для уравнения $y' = -y$ имеем:
- $y' = -y$
- $y'' = -y' = y$

Подставляем в (*):

$$
y_1 + \alpha_0 y_0 = h(-\beta_1 y_1 - \beta_0 y_0) + h^2(\gamma_1 y_1 + \gamma_0 y_0)
$$

Раскрываем скобки:

$$
y_1 + \alpha_0 y_0 = -h\beta_1 y_1 - h\beta_0 y_0 + h^2\gamma_1 y_1 + h^2\gamma_0 y_0
$$

Собираем члены с $y_1$ в левой части, с $y_0$ - в правой:

$$
y_1 + h\beta_1 y_1 - h^2\gamma_1 y_1 = -\alpha_0 y_0 - h\beta_0 y_0 + h^2\gamma_0 y_0
$$

Факторизуем:

$$
y_1(1 + h\beta_1 - h^2\gamma_1) = y_0(-\alpha_0 - h\beta_0 + h^2\gamma_0)
$$

Получаем:

$$
y_1 = y_0 \cdot \frac{-\alpha_0 - h\beta_0 + h^2\gamma_0}{1 + h\beta_1 - h^2\gamma_1} \quad (**)
$$

## Разложение точного решения в ряд Тейлора

Точное решение $y(x_0 + h)$ разложим в ряд Тейлора в окрестности $x_0$:

$$
y(x_0 + h) = y_0 + hy'_0 + \frac{h^2}{2}y''_0 + \frac{h^3}{6}y'''_0 + \frac{h^4}{24}y^{(4)}_0 + O(h^5)
$$

Для уравнения $y' = -y$ производные равны:
- $y'_0 = -y_0$
- $y''_0 = y_0$
- $y'''_0 = -y_0$
- $y^{(4)}_0 = y_0$

Подставляем:

$$
\begin{aligned}
y(x_0 + h) &= y_0 + h(-y_0) + \frac{h^2}{2}(y_0) + \frac{h^3}{6}(-y_0) + \frac{h^4}{24}(y_0) + O(h^5) \\
&= y_0\left(1 - h + \frac{h^2}{2} - \frac{h^3}{6} + \frac{h^4}{24} + O(h^5)\right)
\end{aligned}
$$

## Разложение численной формулы

Обозначим:
- $P(h) = -\alpha_0 - h\beta_0 + h^2\gamma_0$
- $Q(h) = 1 + h\beta_1 - h^2\gamma_1$

Тогда из (**):

$$
\frac{y_1}{y_0} = \frac{P(h)}{Q(h)} = P(h) \cdot Q(h)^{-1}
$$

### Разложение $Q(h)^{-1}$

Используем разложение $(1 + u)^{-1} = 1 - u + u^2 - u^3 + u^4 + O(u^5)$:

$$
Q(h)^{-1} = (1 + h\beta_1 - h^2\gamma_1)^{-1}
$$

Обозначим $u = h\beta_1 - h^2\gamma_1$:

$$
\begin{aligned}
(1 + u)^{-1} &= 1 - u + u^2 - u^3 + u^4 + O(u^5) \\
&= 1 - (h\beta_1 - h^2\gamma_1) + (h\beta_1 - h^2\gamma_1)^2 - (h\beta_1 - h^2\gamma_1)^3 + (h\beta_1)^4 + O(h^5)
\end{aligned}
$$

Вычислим степени $u$:

$$
u = h\beta_1 - h^2\gamma_1
$$

$$
u^2 = (h\beta_1 - h^2\gamma_1)^2 = h^2\beta_1^2 - 2h^3\beta_1\gamma_1 + h^4\gamma_1^2 = h^2\beta_1^2 - 2h^3\beta_1\gamma_1 + O(h^4)
$$

$$
u^3 = (h\beta_1 - h^2\gamma_1)^3 = h^3\beta_1^3 + O(h^4)
$$

$$
u^4 = h^4\beta_1^4 + O(h^5)
$$

Подставляем:

$$
\begin{aligned}
Q(h)^{-1} &= 1 - (h\beta_1 - h^2\gamma_1) + (h^2\beta_1^2 - 2h^3\beta_1\gamma_1) - h^3\beta_1^3 + h^4\beta_1^4 + O(h^5) \\
&= 1 - h\beta_1 + h^2\gamma_1 + h^2\beta_1^2 - 2h^3\beta_1\gamma_1 - h^3\beta_1^3 + h^4\beta_1^4 + O(h^5) \\
&= 1 - h\beta_1 + h^2(\gamma_1 + \beta_1^2) + h^3(-2\beta_1\gamma_1 - \beta_1^3) + h^4\beta_1^4 + O(h^5)
\end{aligned}
$$

### Произведение $P(h) \cdot Q(h)^{-1}$

$$
P(h) = -\alpha_0 - h\beta_0 + h^2\gamma_0
$$

$$
Q(h)^{-1} = 1 - h\beta_1 + h^2(\gamma_1 + \beta_1^2) + h^3(-2\beta_1\gamma_1 - \beta_1^3) + h^4\beta_1^4
$$

Перемножаем (до степени $h^3$ включительно):

$$
\begin{aligned}
P \cdot Q^{-1} &= (-\alpha_0 - h\beta_0 + h^2\gamma_0) \times [1 - h\beta_1 + h^2(\gamma_1 + \beta_1^2) + h^3(-2\beta_1\gamma_1 - \beta_1^3) + h^4\beta_1^4]
\end{aligned}
$$

**Коэффициент при $h^0$:**
$$
-\alpha_0 \cdot 1 = -\alpha_0
$$

**Коэффициент при $h^1$:**
$$
-\alpha_0 \cdot (-\beta_1) + (-\beta_0) \cdot 1 = \alpha_0\beta_1 - \beta_0
$$

**Коэффициент при $h^2$:**
$$
-\alpha_0 \cdot (\gamma_1 + \beta_1^2) + (-\beta_0) \cdot (-\beta_1) + \gamma_0 \cdot 1
$$
$$
= -\alpha_0\gamma_1 - \alpha_0\beta_1^2 + \beta_0\beta_1 + \gamma_0
$$
$$
= \gamma_0 + \beta_0\beta_1 - \alpha_0\gamma_1 - \alpha_0\beta_1^2
$$

**Коэффициент при $h^3$:**
$$
-\alpha_0 \cdot (-2\beta_1\gamma_1 - \beta_1^3) + (-\beta_0) \cdot (\gamma_1 + \beta_1^2) + \gamma_0 \cdot (-\beta_1)
$$
$$
= 2\alpha_0\beta_1\gamma_1 + \alpha_0\beta_1^3 - \beta_0\gamma_1 - \beta_0\beta_1^2 - \gamma_0\beta_1
$$
$$
= -\gamma_0\beta_1 - \beta_0\beta_1^2 - \beta_0\gamma_1 + 2\alpha_0\beta_1\gamma_1 + \alpha_0\beta_1^3
$$

**Коэффициент при $h^4$:**
$$
-\alpha_0 \cdot \beta_1^4 + (-\beta_0) \cdot (-2\beta_1\gamma_1 - \beta_1^3) + \gamma_0 \cdot (\gamma_1 + \beta_1^2)
$$
$$
= -\alpha_0\beta_1^4 + 2\beta_0\beta_1\gamma_1 + \beta_0\beta_1^3 + \gamma_0\gamma_1 + \gamma_0\beta_1^2
$$
$$
= \gamma_0\gamma_1 + \gamma_0\beta_1^2 + \beta_0\beta_1^3 + 2\beta_0\beta_1\gamma_1 - \alpha_0\beta_1^4
$$

## Система уравнений

Приравниваем к точному решению $1 - h + \frac{h^2}{2} - \frac{h^3}{6} + \frac{h^4}{24}$:

**Уравнение 1** (коэффициент при $h^0$):
$$
-\alpha_0 = 1 \implies \boxed{\alpha_0 = -1}
$$

**Уравнение 2** (коэффициент при $h^1$):
$$
\alpha_0\beta_1 - \beta_0 = -1
$$
$$
-\beta_1 - \beta_0 = -1
$$
$$
\boxed{\beta_0 + \beta_1 = 1} \quad (E1)
$$

**Уравнение 3** (коэффициент при $h^2$):
$$
\gamma_0 + \beta_0\beta_1 - \alpha_0\gamma_1 - \alpha_0\beta_1^2 = \frac{1}{2}
$$
$$
\gamma_0 + \beta_0\beta_1 + \gamma_1 + \beta_1^2 = \frac{1}{2}
$$
$$
\boxed{\gamma_0 + \beta_0\beta_1 + \gamma_1 + \beta_1^2 = \frac{1}{2}} \quad (E2)
$$

**Уравнение 4** (коэффициент при $h^3$):
$$
-\gamma_0\beta_1 - \beta_0\beta_1^2 - \beta_0\gamma_1 + 2\alpha_0\beta_1\gamma_1 + \alpha_0\beta_1^3 = -\frac{1}{6}
$$
$$
-\gamma_0\beta_1 - \beta_0\beta_1^2 - \beta_0\gamma_1 - 2\beta_1\gamma_1 - \beta_1^3 = -\frac{1}{6}
$$
$$
\boxed{\gamma_0\beta_1 + \beta_0\beta_1^2 + \beta_0\gamma_1 + 2\beta_1\gamma_1 + \beta_1^3 = \frac{1}{6}} \quad (E3)
$$

## Решение системы уравнений

У нас 3 уравнения и 5 неизвестных ($\alpha_0, \beta_0, \beta_1, \gamma_0, \gamma_1$). Нужно выбрать 2 свободных параметра.

### Выбор: $\beta_1 = \frac{1}{2}$

Из (E1):
$$
\beta_0 = 1 - \beta_1 = 1 - \frac{1}{2} = \frac{1}{2}
$$

Из (E2):
$$
\gamma_0 + \frac{1}{2} \cdot \frac{1}{2} + \gamma_1 + \left(\frac{1}{2}\right)^2 = \frac{1}{2}
$$
$$
\gamma_0 + \frac{1}{4} + \gamma_1 + \frac{1}{4} = \frac{1}{2}
$$
$$
\gamma_0 + \gamma_1 + \frac{1}{2} = \frac{1}{2}
$$
$$
\boxed{\gamma_0 + \gamma_1 = 0} \quad (E2')
$$

Из (E3):
$$
\gamma_0 \cdot \frac{1}{2} + \frac{1}{2} \cdot \frac{1}{4} + \frac{1}{2} \cdot \gamma_1 + 2 \cdot \frac{1}{2} \cdot \gamma_1 + \left(\frac{1}{2}\right)^3 = \frac{1}{6}
$$
$$
\frac{\gamma_0}{2} + \frac{1}{8} + \frac{\gamma_1}{2} + \gamma_1 + \frac{1}{8} = \frac{1}{6}
$$
$$
\frac{\gamma_0}{2} + \frac{3\gamma_1}{2} + \frac{1}{4} = \frac{1}{6}
$$
$$
\frac{\gamma_0 + 3\gamma_1}{2} = \frac{1}{6} - \frac{1}{4} = \frac{2 - 3}{12} = -\frac{1}{12}
$$
$$
\gamma_0 + 3\gamma_1 = -\frac{1}{6}
$$
$$
\boxed{\gamma_0 + 3\gamma_1 = -\frac{1}{6}} \quad (E3')
$$

Решаем систему (E2') и (E3'):

Из (E2'): $\gamma_0 = -\gamma_1$

Подставляем в (E3'):
$$
-\gamma_1 + 3\gamma_1 = -\frac{1}{6}
$$
$$
2\gamma_1 = -\frac{1}{6}
$$
$$
\boxed{\gamma_1 = -\frac{1}{12}}
$$

$$
\boxed{\gamma_0 = -\gamma_1 = \frac{1}{12}}
$$

## Итоговые коэффициенты

$$
\begin{cases}
\alpha_0 = -1 \\
\beta_0 = \frac{1}{2} \\
\beta_1 = \frac{1}{2} \\
\gamma_0 = \frac{1}{12} \\
\gamma_1 = -\frac{1}{12}
\end{cases}
$$

## Подстановка в формулу

Из (**):

$$
y_1 = y_0 \cdot \frac{-\alpha_0 - h\beta_0 + h^2\gamma_0}{1 + h\beta_1 - h^2\gamma_1}
$$

Подставляем коэффициенты:

$$
y_1 = y_0 \cdot \frac{-(-1) - h \cdot \frac{1}{2} + h^2 \cdot \frac{1}{12}}{1 + h \cdot \frac{1}{2} - h^2 \cdot \left(-\frac{1}{12}\right)}
$$

$$
y_1 = y_0 \cdot \frac{1 - \frac{h}{2} + \frac{h^2}{12}}{1 + \frac{h}{2} + \frac{h^2}{12}}
$$

Умножим числитель и знаменатель на 12:

$$
\boxed{y_1 = y_0 \cdot \frac{12 - 6h + h^2}{12 + 6h + h^2}}
$$

## Проверка порядка точности

Разложим полученную формулу в ряд Тейлора. Обозначим:

$$
f(h) = \frac{1 - \frac{h}{2} + \frac{h^2}{12}}{1 + \frac{h}{2} + \frac{h^2}{12}}
$$

Используем разложение $\frac{1 + a}{1 + b} = (1 + a)(1 + b)^{-1}$:

$$
f(h) = \left(1 - \frac{h}{2} + \frac{h^2}{12}\right) \cdot \left(1 + \frac{h}{2} + \frac{h^2}{12}\right)^{-1}
$$

Разложим $\left(1 + \frac{h}{2} + \frac{h^2}{12}\right)^{-1}$:

Пусть $u = \frac{h}{2} + \frac{h^2}{12}$:

$$
(1 + u)^{-1} = 1 - u + u^2 - u^3 + O(u^4)
$$

$$
\begin{aligned}
u &= \frac{h}{2} + \frac{h^2}{12} \\
u^2 &= \left(\frac{h}{2}\right)^2 + 2 \cdot \frac{h}{2} \cdot \frac{h^2}{12} + O(h^4) = \frac{h^2}{4} + \frac{h^3}{12} + O(h^4) \\
u^3 &= \left(\frac{h}{2}\right)^3 + O(h^4) = \frac{h^3}{8} + O(h^4)
\end{aligned}
$$

$$
\begin{aligned}
(1 + u)^{-1} &= 1 - \left(\frac{h}{2} + \frac{h^2}{12}\right) + \left(\frac{h^2}{4} + \frac{h^3}{12}\right) - \frac{h^3}{8} + O(h^4) \\
&= 1 - \frac{h}{2} - \frac{h^2}{12} + \frac{h^2}{4} + \frac{h^3}{12} - \frac{h^3}{8} + O(h^4) \\
&= 1 - \frac{h}{2} + h^2\left(\frac{1}{4} - \frac{1}{12}\right) + h^3\left(\frac{1}{12} - \frac{1}{8}\right) + O(h^4) \\
&= 1 - \frac{h}{2} + h^2 \cdot \frac{3 - 1}{12} + h^3 \cdot \frac{2 - 3}{24} + O(h^4) \\
&= 1 - \frac{h}{2} + \frac{h^2}{6} - \frac{h^3}{24} + O(h^4)
\end{aligned}
$$

Теперь умножаем:

$$
f(h) = \left(1 - \frac{h}{2} + \frac{h^2}{12}\right) \cdot \left(1 - \frac{h}{2} + \frac{h^2}{6} - \frac{h^3}{24}\right)
$$

$$
\begin{aligned}
f(h) &= 1 - \frac{h}{2} + \frac{h^2}{6} - \frac{h^3}{24} \\
&\quad - \frac{h}{2} + \frac{h^2}{4} - \frac{h^3}{12} + O(h^4) \\
&\quad + \frac{h^2}{12} - \frac{h^3}{24} + O(h^4) \\
&= 1 - h + h^2\left(\frac{1}{6} + \frac{1}{4} + \frac{1}{12}\right) + h^3\left(-\frac{1}{24} - \frac{1}{12} - \frac{1}{24}\right) + O(h^4)
\end{aligned}
$$

Вычисляем коэффициент при $h^2$:
$$
\frac{1}{6} + \frac{1}{4} + \frac{1}{12} = \frac{2 + 3 + 1}{12} = \frac{6}{12} = \frac{1}{2}
$$

Вычисляем коэффициент при $h^3$:
$$
-\frac{1}{24} - \frac{1}{12} - \frac{1}{24} = -\frac{1 + 2 + 1}{24} = -\frac{4}{24} = -\frac{1}{6}
$$

Итого:

$$
\boxed{f(h) = 1 - h + \frac{h^2}{2} - \frac{h^3}{6} + O(h^4)}
$$

Это **точно совпадает** с разложением точного решения до члена $h^3$ включительно!

## Вывод

Формула 
$$
\boxed{y_1 = y_0 \cdot \frac{12 - 6h + h^2}{12 + 6h + h^2}}
$$

обеспечивает **3-й порядок точности** (локальная ошибка O(h⁴)), что достаточно для "разгона" основной двухшаговой схемы 4-го порядка.

