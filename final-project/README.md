# Final Project
This is final project for Golang for Women class. Details on the project is explained in Google Classroom (not public).

## API
| METHOD | URL                        | Description                                 |
|--------|----------------------------|---------------------------------------------|
| GET    | `/login`                   | Login to the system with registered account |
| POST   | `/register`                | Register new login account (admin)          |
| GET    | `/products`                | Get all products                            |
| POST   | `/products`                | Create new product                          |
| GET    | `/products/:uuid`          | Get product details of specific UUID        |
| PUT    | `/products/:uuid`          | Update product of specific UUID             |
| DELETE | `/products/:uuid`          | Delete product of specific UUID             |
| GET    | `/products/variants`       | Get all variants                            |
| POST   | `/products/variants`       | Create new variant                          |
| GET    | `/products/variants/:uuid` | Get variant details of specific UUID        |
| PUT    | `/products/variants/:uuid` | Update variant of specific UUID             |
| DELETE | `/products/variants/:uuid` | Delete variant of specific UUID             |

### NOTES
- To access this project via postman, you can import [this](https://api.postman.com/collections/28968065-5c6aac55-3bac-49b6-a69a-f8575ab2168a?access_key=PMAT-01HAMKNEWHEPQ071RWMYP6N3K2) json.
- To try the live demo of this project, you may set the base URL to [https://golang-for-women-production.up.railway.app](https://golang-for-women-production.up.railway.app) (set it to the `API_URL` if you are accessing from postman)

## Technical Issues
- Unhandled case on updating a variant which `product_id` null (which means the product has been deleted) without updating the `product_id`. This is because by default, `product_id` should not be null, while updating a variant's `product_id` is optional. It also lead to another weird, unhandled case where the request should be sent twice  when updating the `product_id` before it is updated on the database.
