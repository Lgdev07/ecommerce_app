<h1 align="center">
      <img alt="EcommerceApp" title="EcommerceApp" src=".github/logo.png" width="300px" />
</h1>

<h3 align="center">
  EcommerceApp Server
</h3>

<p align="center">Api to serve correios services in a easy way 📖</p>
<p align="center">Made with Golang 🚀</p>

<p align="center">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/Lgdev07/ecommerce_app?color=%2304D361">

  <img alt="Made by Lgdev07" src="https://img.shields.io/badge/made%20by-Lgdev07-%2304D361">

  <img alt="License" src="https://img.shields.io/badge/license-MIT-%2304D361">

  <a href="https://github.com/Lgdev07/ecommerce_app/stargazers">
    <img alt="Stargazers" src="https://img.shields.io/github/stars/Lgdev07/ecommerce_app?style=social">
  </a>
</p>

<p align="center">
  <a href="#-installation-and-execution">Installation and execution</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-available-routes">Available Routes</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>&nbsp;&nbsp;&nbsp;
</p>


## 🚀 Installation and execution

1. Clone this repository and go to the directory;
2. Rename sample .env;

<h4> 🔧 Development </h4>

1. Rename `.env_sample` to `.env`;
1. Run `docker-compose up`;
2. Make the Requests;

<h4> 🧪 Tests </h4>

1. Run `docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit`;

## 🛣️ Available Routes

  POST - Create Checkout:
  - **/checkout** <br>
  Expected Json Body Request:<br>
  ```
    {
        "products": [
            {
                "id": 1,
                "quantity": 1
            }
        ]
    }
  ```
  Expected Json Response:<br>
  ```
    {
        "total_amount": 20000,
        "total_amount_with_discount": 19500,
        "total_discount": 500,
        "products": [
            {
                "id": 1,
                "quantity": 2,
                "unit_amount": 10000,
                "total_amount": 20000,
                "discount": 500,
                "is_gift": false
            },
            {
                "id": 3,
                "quantity": 1,
                "unit_amount": 0,
                "total_amount": 0,
                "discount": 0,
                "is_gift": true
            }
        ]
    }
  ```

## 🤔 How to contribute

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m 'feat: My new feature'`;
- Push to your branch: `git push origin my-feature`.

After the merge of your pull request is done, you can delete your branch.

---