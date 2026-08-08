
# Excursion Backend

REST API that manages tours and the administration of an excursion website
## API Reference

#### Get all excursions

```http
  GET /api/v1/excursions
```

#### Get excursion

```http
  GET /api/v1/excursion/id={id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

#### Update excursion

```http
  GET /api/v1/excursion/id={id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to update |
| `img`      | `string` |Image of item to update|
| `title`      | `string` |Title of item to update|
| `desc`      | `string` |Description of item to update|
| `buy`      | `string` |BuyLink of item to update|
| `findMore`      | `string` |findMoreLink of item to update|

#### Delete excursion

```http
  GET /api/v1/excursion/id={id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to delete |

