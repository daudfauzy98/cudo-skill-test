# cudo-skill-test

### SQL to create order-report
```sql
SELECT OrderDate, CategoryName, SUM(Quantity) AS Quantity FROM OrderDetails JOIN Orders ON OrderDetails.OrderID = Orders.OrderID JOIN Products ON OrderDetails.ProductID = Products.ProductID JOIN Categories ON Products.CategoryID = Categories.CategoryID GROUP BY OrderDate, CategoryName ORDER BY OrderDate ASC;
```
