## Coffeeco

- reference: https://github.com/PacktPublishing/Domain-Driven-Design-with-GoLang/tree/main/chapter5

### Ubiquitous Languages

- **Coffee lovers**: What CoffeeCo calls its customers.
- **CoffeeBux**: This is the name of their loyalty program. Coffee lovers earn one CoffeeBux for each drink or accessory they purchase.
- **Tiny**, **medium**, and **massive**: The sizes of the drinks are in ascending order. Some drinks are only available in one size, others in all three. Everything on the menu fits into these categories.

### Domain Design

- **Store**
- **Products**
- **Loyalty**
- **Subscription**

### Business Logic / Scop

- Purchasing a drink or accessory using CoffeeBux
- Purchasing a drink or accessory with a debit/credit card
- Purchasing a drink or accessory with cash
- Earning CoffeeBux on purchases
- Store-specific (but not national) discounts
- We can assume all purchases are in USD for now; in the future, we need to support many currencies though
- Drinks only need to come in one size for now
