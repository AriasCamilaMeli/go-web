# Bootcamp Go

## Práctica 2 Go Web

### Objetivo

El objetivo de esta guía práctica es afianzar los conceptos sobre el **Método POST**, vistos en el módulo. Para eso, vamos a plantear una serie de ejercicios simples e incrementales (trabajaremos y agregaremos complejidad a lo que tenemos que construir) a lo largo del módulo.

### Problema

Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Para poder hacer esto, necesitan un **servidor** que ejecute una API que les permita manipular los productos cargados desde distintos clientes. Los campos que conforman un producto son:

| Nombre       | Tipo de dato JSON | Tipo de dato GO | Descripción                          | Ejemplo              |
|--------------|-------------------|-----------------|--------------------------------------|----------------------|
| id           | number            | int             | Identificador en conjunto de datos   | 15                   |
| name         | string            | string          | Nombre caracterizado                | Cheese - St. Andre   |
| quantity     | number            | int             | Cantidad almacenada                 | 60                   |
| code_value   | string            | string          | Código alfanumérico característico  | S73191A              |
| is_published | boolean           | bool            | El producto se encuentra publicado o no | True             |
| expiration   | string            | string          | Fecha de vencimiento                | 12/04/2022           |
| price        | number            | float64         | Precio del producto                 | 50.15                |

## Ejercicio 1: Añadir un producto

En esta ocasión vamos a añadir un producto al slice cargado en memoria. Dentro de la ruta **/products** añadimos el método **POST**, al cual vamos a enviar en el cuerpo de la request el nuevo producto. El mismo tiene ciertas restricciones, conozcámoslas:

1. No es necesario pasar el Id, al momento de añadirlo se debe inferir del estado de la lista de productos, verificando que no se repitan ya que debe ser un campo único.
2. Ningún dato puede estar vacío, exceptuando **is_published** (vacío indica un valor false).
3. El campo **code_value** debe ser único para cada producto.
4. Los tipos de datos deben coincidir con los definidos en el planteo del problema.
5. La fecha de vencimiento debe tener el formato: **XX/XX/XXXX**, además debemos verificar que día, mes y año sean valores válidos.

Recordá: si una consulta está mal formulada por parte del cliente, el status code cae en los **4XX**.

## Ejercicio 2: Traer el producto

Realiza una consulta a un método **GET** con el id del producto recién añadido, tené en cuenta que la lista de productos se encuentra cargada en la memoria, si terminás la ejecución del programa este producto no estará en la próxima ejecución.